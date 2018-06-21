package ome

import (
	"errors"

	"github.com/republicprotocol/republic-go/order"
)

// ErrComputationNotFound is returned when the Storer cannot find a Computation
// associated with a ComputationID.
var ErrComputationNotFound = errors.New("computation not found")

// ComputationStorer for the Computations that are synchronised.
type ComputationStorer interface {
	PutComputation(computation Computation) error
	DeleteComputation(id ComputationID) error
	Computation(id ComputationID) (Computation, error)
	Computations() (ComputationIterator, error)
}

// ComputationIterator is used to iterate over a Computation collection.
type ComputationIterator interface {

	// Next progresses the cursor. Returns true if the new cursor is still in
	// the range of the Computation collection, otherwise false.
	Next() bool

	// Cursor returns the Computation at the current cursor location. Returns
	// an error if the cursor is out of range.
	Cursor() (Computation, error)

	// Collect all Computations in the iterator into a slice.
	Collect() ([]Computation, error)
}

// OrderFragmentStorer for the order.Fragments that are received.
type OrderFragmentStorer interface {
	OrderFragment(id order.ID) (order.Fragment, error)
}

// Storer combines the ComputationStorer interface and the
// OrderFragmentStorer interface into a unified set of storage functions.
type Storer interface {
	ComputationStorer
	OrderFragmentStorer
}
