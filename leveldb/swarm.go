package leveldb

import (
	"encoding/json"
	"errors"
	"math"
	"time"

	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/swarm"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// ErrNonceTooLow is returned if the nonce of the multiaddress is lower than the
// one present in the local store.
var ErrNonceTooLow = errors.New("nonce too low")

// SwarmMultiAddressValue is the storage format for multiaddresses being stored in
// LevelDB. It contains additional timestamping information so that LevelDB can
// provide pruning.
type SwarmMultiAddressValue struct {
	Nonce        uint64
	MultiAddress identity.MultiAddress
	Timestamp    time.Time
}

// SwarmMultiAddressesIterator implements the swarm.MultiAddressStorer using a
// LevelDB iterator.
type SwarmMultiAddressesIterator struct {
	inner iterator.Iterator
}

func newSwarmMultiAddressIterator(iter iterator.Iterator) *SwarmMultiAddressesIterator {
	return &SwarmMultiAddressesIterator{
		inner: iter,
	}
}

// Next implements the swarm.MultiAddressIterator interface.
func (iter *SwarmMultiAddressesIterator) Next() bool {
	return iter.inner.Next()
}

// Cursor implements the swarm.MultiAddressIterator interface.
func (iter *SwarmMultiAddressesIterator) Cursor() (identity.MultiAddress, uint64, error) {
	if !iter.inner.Valid() {
		return identity.MultiAddress{}, 0, swarm.ErrCursorOutOfRange
	}

	if isEqual(iter.inner.Key(), identity.Address(0).Hash()) {
		iter.Next()
		return iter.Cursor()
	}
	value := SwarmMultiAddressValue{}
	data := iter.inner.Value()
	if err := json.Unmarshal(data, &value); err != nil {
		return identity.MultiAddress{}, 0, swarm.ErrCursorOutOfRange
	}
	return value.MultiAddress, value.Nonce, iter.inner.Error()
}

// Collect implements the swarm.MultiAddressIterator interface.
func (iter *SwarmMultiAddressesIterator) Collect() ([]identity.MultiAddress, []uint64, error) {
	multiaddresses := []identity.MultiAddress{}
	nonces := []uint64{}
	for iter.Next() {
		multiaddress, nonce, err := iter.Cursor()
		if err != nil {
			return multiaddresses, nonces, err
		}

		multiaddresses = append(multiaddresses, multiaddress)
		nonces = append(nonces, nonce)
	}
	return multiaddresses, nonces, iter.inner.Error()
}

// Release implements the swarm.MultiAddressIterator interface.
func (iter *SwarmMultiAddressesIterator) Release() {
	iter.inner.Release()
}

// SwarmMultiAddressTable implements the swarm.MultiAddressStorer interface using
// LevelDB.
type SwarmMultiAddressTable struct {
	db     *leveldb.DB
	expiry time.Duration
}

// NewSwarmMultiAddressTable returns a new SwarmMultiAddressTable that uses the
// given LevelDB instance to store and load values from the disk.
func NewSwarmMultiAddressTable(db *leveldb.DB, expiry time.Duration) *SwarmMultiAddressTable {
	return &SwarmMultiAddressTable{db: db, expiry: expiry}
}

// PutMultiAddress implements the swarm.MultiAddressStorer interface.
func (table *SwarmMultiAddressTable) PutMultiAddress(multiaddress identity.MultiAddress, nonce uint64) (bool, error) {
	selfMultiAddr, _, err := table.Self()
	if err == nil && selfMultiAddr.String() == multiaddress.String() {
		return false, nil
	}
	isNew := false
	value := SwarmMultiAddressValue{
		Nonce:        nonce,
		MultiAddress: multiaddress,
		Timestamp:    time.Now(),
	}
	oldMultiAddr, oldNonce, err := table.MultiAddress(multiaddress.Address())
	if err != nil && err == swarm.ErrMultiAddressNotFound {
		isNew = true
	}
	// Return err if nonce is too low
	if oldNonce > nonce {
		return isNew, ErrNonceTooLow
	}
	// If there is a change in the multiaddress stored, then return true
	if err == nil && oldMultiAddr.String() != multiaddress.String() {
		isNew = true
	}

	data, err := json.Marshal(value)
	if err != nil {
		return isNew, err
	}
	return isNew, table.db.Put(table.key(multiaddress.Address().Hash()), data, nil)
}

// MultiAddress implements the swarm.MultiAddressStorer interface.
func (table *SwarmMultiAddressTable) MultiAddress(address identity.Address) (identity.MultiAddress, uint64, error) {
	selfMultiAddr, selfNonce, err := table.Self()
	if err == nil && selfMultiAddr.Address() == address {
		return selfMultiAddr, selfNonce, nil
	}

	return table.multiAddress(address)
}

// MultiAddresses implements the swarm.MultiAddressStorer interface.
func (table *SwarmMultiAddressTable) MultiAddresses() (swarm.MultiAddressIterator, error) {
	iter := table.db.NewIterator(&util.Range{Start: table.key(SwarmMultiAddressIterBegin), Limit: table.key(SwarmMultiAddressIterEnd)}, nil)
	return newSwarmMultiAddressIterator(iter), nil
}

// Prune iterates over all multiaddresses and deletes those that have expired.
func (table *SwarmMultiAddressTable) Prune() (err error) {
	iter := table.db.NewIterator(&util.Range{Start: table.key(SwarmMultiAddressIterBegin), Limit: table.key(SwarmMultiAddressIterEnd)}, nil)
	defer iter.Release()

	now := time.Now()
	for iter.Next() {
		key := iter.Key()
		value := SwarmMultiAddressValue{}
		if localErr := json.Unmarshal(iter.Value(), &value); localErr != nil {
			err = localErr
			continue
		}
		if value.Timestamp.Add(table.expiry).Before(now) {
			if localErr := table.db.Delete(key, nil); localErr != nil {
				err = localErr
			}
		}
	}
	return err
}

// Self implements the swarm.MultiAddressStorer interface.
func (table *SwarmMultiAddressTable) Self() (identity.MultiAddress, uint64, error) {
	return table.multiAddress(identity.Address(0))
}

// PutSelf implements the swarm.MultiAddressStorer interface.
func (table *SwarmMultiAddressTable) PutSelf(multiAddr identity.MultiAddress, nonce uint64) (uint64, error) {
	_, oldNonce, err := table.Self()
	if err != nil && err != swarm.ErrMultiAddressNotFound {
		return 0, err
	}
	value := SwarmMultiAddressValue{
		Nonce:        uint64(math.Max(float64(nonce), float64(oldNonce)) + 1),
		MultiAddress: multiAddr,
		Timestamp:    time.Now(),
	}
	data, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	return value.Nonce, table.db.Put(table.key(identity.Address(0).Hash()), data, nil)
}

func (table *SwarmMultiAddressTable) multiAddress(address identity.Address) (identity.MultiAddress, uint64, error) {
	data, err := table.db.Get(table.key(address.Hash()), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			err = swarm.ErrMultiAddressNotFound
		}
		return identity.MultiAddress{}, 0, err
	}
	value := SwarmMultiAddressValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return identity.MultiAddress{}, 0, err
	}
	return value.MultiAddress, value.Nonce, nil
}

func (table *SwarmMultiAddressTable) key(k []byte) []byte {
	return append(append(SwarmMultiAddressTableBegin, k...), SwarmMultiAddressTablePadding...)
}

func isEqual(key, address []byte) bool {
	if len(address) != 32 || len(key) != 66 {
		return false
	}
	for i := 0; i < 32; i++ {
		if address[i] != key[i+2] {
			return false
		}
	}
	return true
}
