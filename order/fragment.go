package order

import (
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"time"

	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/shamir"
)

// An FragmentID is the Keccak256 hash of a Fragment.
type FragmentID [32]byte

// Equal returns an equality check between two orderFragment ID.
func (id FragmentID) Equal(other FragmentID) bool {
	return bytes.Equal(id[:], other[:])
}

// A Fragment is a secret share of an Order, created using Shamir's secret
// sharing on the secure fields in an Order.
type Fragment struct {
	OrderID         ID         `json:"orderID"`
	OrderType       Type       `json:"orderType"`
	OrderParity     Parity     `json:"orderParity"`
	OrderSettlement Settlement `json:"orderSettlement"`
	OrderExpiry     time.Time  `json:"orderExpiry"`
	ID              FragmentID `json:"id"`

	Tokens        shamir.Share `json:"tokens"`
	Price         CoExpShare   `json:"price"`
	Volume        CoExpShare   `json:"volume"`
	MinimumVolume CoExpShare   `json:"minimumVolume"`
	Nonce         shamir.Share `json:"nonce"`
}

// NewFragment returns a new Fragment and computes the FragmentID.
func NewFragment(orderID ID, orderType Type, orderParity Parity, orderSettlement Settlement, orderExpiry time.Time, tokens shamir.Share, price, volume, minimumVolume CoExpShare, nonce shamir.Share) Fragment {
	fragment := Fragment{
		OrderID:         orderID,
		OrderType:       orderType,
		OrderParity:     orderParity,
		OrderSettlement: orderSettlement,
		OrderExpiry:     orderExpiry,

		Tokens:        tokens,
		Price:         price,
		Volume:        volume,
		MinimumVolume: minimumVolume,
		Nonce:         nonce,
	}
	fragment.ID = FragmentID(fragment.Hash())
	return fragment
}

// Hash returns the Keccak256 hash of a Fragment. This hash is used to create
// the FragmentID and signature for a Fragment.
func (fragment *Fragment) Hash() [32]byte {
	hash := crypto.Keccak256(fragment.Bytes())
	hash32 := [32]byte{}
	copy(hash32[:], hash)
	return hash32
}

// Bytes returns a Fragment serialized into a bytes.
// TODO: This function should return an error.
func (fragment *Fragment) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, fragment.OrderType)
	binary.Write(buf, binary.BigEndian, fragment.OrderID)
	binary.Write(buf, binary.BigEndian, fragment.OrderParity)
	binary.Write(buf, binary.BigEndian, fragment.OrderSettlement)
	binary.Write(buf, binary.BigEndian, fragment.OrderExpiry.Unix())
	binary.Write(buf, binary.BigEndian, fragment.Tokens)
	binary.Write(buf, binary.BigEndian, fragment.Price)
	binary.Write(buf, binary.BigEndian, fragment.Volume)
	binary.Write(buf, binary.BigEndian, fragment.MinimumVolume)
	binary.Write(buf, binary.BigEndian, fragment.Nonce)
	return buf.Bytes()
}

// Equal returns an equality check between two Orders.
func (fragment *Fragment) Equal(other *Fragment) bool {
	return bytes.Equal(fragment.OrderID[:], other.OrderID[:]) &&
		fragment.OrderType == other.OrderType &&
		fragment.OrderParity == other.OrderParity &&
		fragment.OrderExpiry.Equal(other.OrderExpiry) &&
		bytes.Equal(fragment.ID[:], other.ID[:]) &&
		fragment.Tokens.Equal(&other.Tokens) &&
		fragment.Price.Equal(&other.Price) &&
		fragment.Volume.Equal(&other.Volume) &&
		fragment.MinimumVolume.Equal(&other.MinimumVolume) &&
		fragment.Nonce.Equal(&other.Nonce)
}

// Encrypt a Fragment using an rsa.PublicKey.
func (fragment *Fragment) Encrypt(pubKey rsa.PublicKey) (EncryptedFragment, error) {
	var err error
	encryptedFragment := EncryptedFragment{
		OrderID:         fragment.OrderID,
		OrderType:       fragment.OrderType,
		OrderParity:     fragment.OrderParity,
		OrderSettlement: fragment.OrderSettlement,
		OrderExpiry:     fragment.OrderExpiry,
		ID:              fragment.ID,
	}
	encryptedFragment.Tokens, err = fragment.Tokens.Encrypt(pubKey)
	if err != nil {
		return encryptedFragment, err
	}
	encryptedFragment.Price, err = fragment.Price.Encrypt(pubKey)
	if err != nil {
		return encryptedFragment, err
	}
	encryptedFragment.Volume, err = fragment.Volume.Encrypt(pubKey)
	if err != nil {
		return encryptedFragment, err
	}
	encryptedFragment.MinimumVolume, err = fragment.MinimumVolume.Encrypt(pubKey)
	if err != nil {
		return encryptedFragment, err
	}
	encryptedFragment.Nonce, err = fragment.Nonce.Encrypt(pubKey)
	if err != nil {
		return encryptedFragment, err
	}
	return encryptedFragment, nil
}

// An EncryptedFragment is a Fragment that has been encrypted by an RSA public
// key.
type EncryptedFragment struct {
	OrderID         ID                  `json:"orderId"`
	OrderType       Type                `json:"orderType"`
	OrderParity     Parity              `json:"orderParity"`
	OrderSettlement Settlement          `json:"orderSettlement"`
	OrderExpiry     time.Time           `json:"orderExpiry"`
	ID              FragmentID          `json:"id"`
	Tokens          []byte              `json:"tokens"`
	Price           EncryptedCoExpShare `json:"price"`
	Volume          EncryptedCoExpShare `json:"volume"`
	MinimumVolume   EncryptedCoExpShare `json:"minimumVolume"`
	Nonce           []byte              `json:"nonce"`
}

// Decrypt an EncryptedFragment using an rsa.PrivateKey.
func (fragment *EncryptedFragment) Decrypt(privKey rsa.PrivateKey) (Fragment, error) {
	var err error
	decryptedFragment := Fragment{
		OrderID:         fragment.OrderID,
		OrderType:       fragment.OrderType,
		OrderParity:     fragment.OrderParity,
		OrderSettlement: fragment.OrderSettlement,
		OrderExpiry:     fragment.OrderExpiry,
		ID:              fragment.ID,
	}
	if err := decryptedFragment.Tokens.Decrypt(privKey, fragment.Tokens); err != nil {
		return decryptedFragment, err
	}
	decryptedFragment.Price, err = fragment.Price.Decrypt(privKey)
	if err != nil {
		return decryptedFragment, err
	}
	decryptedFragment.Volume, err = fragment.Volume.Decrypt(privKey)
	if err != nil {
		return decryptedFragment, err
	}
	decryptedFragment.MinimumVolume, err = fragment.MinimumVolume.Decrypt(privKey)
	if err != nil {
		return decryptedFragment, err
	}
	if err := decryptedFragment.Nonce.Decrypt(privKey, fragment.Nonce); err != nil {
		return decryptedFragment, err
	}
	return decryptedFragment, nil
}
