// Package hash provides abstractions for digest algorithms.
//
// [Algo] is a set of functions that defines a digest algorithm.
// The digest is stored in a [State] buffer owned by the caller and passed to
// every call, so an Algo isn't tied to any specific digest. Each digest
// package provides one Algo for each variant.
package hash

import "solod.dev/so/c"

// Buffer sizes covering every digest in the library.
const (
	MaxState = 256 // largest digest state; sha512.Digest is 216 bytes
	MaxBlock = 128 // largest block size; sha512 is 128
	MaxSize  = 64  // largest digest size; sha512 is 64
)

// State is storage for a digest state. Its contents belong to the [Algo] that
// initialized it, and mean nothing without it.
//
// The alignment covers the widest field any digest state has.
//
//so:attr aligned(8)
type State [MaxState]byte

// InitFunc writes a fresh digest state into the caller's buffer.
// The state might hold a previous value and must be fully overwritten.
type InitFunc func(state *State)

// WriteFunc adds more data to the running digest.
type WriteFunc func(state *State, p []byte) (int, error)

// SumFunc appends the current digest to in and returns the resulting slice,
// leaving the state unchanged. Requires at least [Algo.Size] bytes of spare
// capacity in in.
type SumFunc func(state *State, in []byte) []byte

// Algo is a digest algorithm operating on a state the caller owns.
// Each variant of a digest is a separate Algo.
type Algo struct {
	Init  InitFunc
	Write WriteFunc
	Sum   SumFunc

	StateSize int // bytes of digest state, at most [MaxState]
	Size      int // digest size, at most [MaxSize]
	BlockSize int // block size, at most [MaxBlock]
}

// Check aborts if the algorithm does not fit this package's buffer sizes.
// Clients that keep digest state call it before using an Algo.
func (a *Algo) Check() {
	c.Assert(a.StateSize <= MaxState, "hash: state too large")
	c.Assert(a.Size <= MaxSize, "hash: digest too large")
	c.Assert(a.BlockSize <= MaxBlock, "hash: block too large")
}
