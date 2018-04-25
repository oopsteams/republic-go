package rpc

import (
	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/orderbook"
	"github.com/republicprotocol/republic-go/rpc/client"
	"github.com/republicprotocol/republic-go/rpc/dht"
	"github.com/republicprotocol/republic-go/rpc/relayer"
	"github.com/republicprotocol/republic-go/rpc/smpcer"
	"github.com/republicprotocol/republic-go/rpc/swarmer"
)

type RPC struct {
	crypter  crypto.Crypter
	dht      dht.DHT
	connPool client.ConnPool

	relayerClient relayer.Client
	relayer       relayer.Relayer

	smpcerClient smpcer.Client
	smpcer       smpcer.Smpcer

	swarmerClient swarmer.Client
	swarmer       swarmer.Swarmer

	onOpenOrder   func([]byte, order.Fragment) error
	onCancelOrder func([]byte, order.ID) error
}

func NewRPC(crypter crypto.Crypter, multiAddress identity.MultiAddress, orderbook *orderbook.Orderbook) *RPC {
	rpc := new(RPC)

	rpc.crypter = crypter
	rpc.dht = dht.NewDHT(multiAddress.Address(), 128)
	rpc.connPool = client.NewConnPool(256)

	rpc.relayerClient = relayer.NewClient(rpc.crypter, &rpc.dht, &rpc.connPool)
	rpc.relayer = relayer.NewRelayer(&rpc.relayerClient, orderbook)

	rpc.smpcerClient = smpcer.NewClient(rpc.crypter, multiAddress, &rpc.connPool)
	rpc.smpcer = smpcer.NewSmpcer(&rpc.smpcerClient, rpc)

	rpc.swarmerClient = swarmer.NewClient(rpc.crypter, multiAddress, &rpc.dht, &rpc.connPool)
	rpc.swarmer = swarmer.NewSwarmer(&rpc.swarmerClient)

	return rpc
}

// OpenOrder implements the smpcer.Delegate interface.
func (rpc *RPC) OpenOrder(signature []byte, orderFragment order.Fragment) error {
	println("RPC OpenOrder")
	if rpc.onOpenOrder != nil {
		return rpc.onOpenOrder(signature, orderFragment)
	}
	return nil
}

// CancelOrder implements the smpcer.Delegate interface.
func (rpc *RPC) CancelOrder(signature []byte, orderID order.ID) error {
	if rpc.onCancelOrder != nil {
		return rpc.onCancelOrder(signature, orderID)
	}
	return nil
}

// OnOpenOrder call the delegate method.
func (rpc *RPC) OnOpenOrder(delegate func([]byte, order.Fragment) error) {
	rpc.onOpenOrder = delegate
}

// OnCancelOrder call the delegate method.
func (rpc *RPC) OnCancelOrder(delegate func([]byte, order.ID) error) {
	rpc.onCancelOrder = delegate
}

// RelayerClient used by the RPC.
func (rpc *RPC) RelayerClient() *relayer.Client {
	return &rpc.relayerClient
}

// Relayer used by the RPC.
func (rpc *RPC) Relayer() *relayer.Relayer {
	return &rpc.relayer
}

// SmpcerClient used by the RPC.
func (rpc *RPC) SmpcerClient() *smpcer.Client {
	return &rpc.smpcerClient
}

// Smpcer used by the RPC.
func (rpc *RPC) Smpcer() *smpcer.Smpcer {
	return &rpc.smpcer
}

// SwarmerClient used by the RPC.
func (rpc *RPC) SwarmerClient() *swarmer.Client {
	return &rpc.swarmerClient
}

// Swarmer used by the RPC.
func (rpc *RPC) Swarmer() *swarmer.Swarmer {
	return &rpc.swarmer
}
