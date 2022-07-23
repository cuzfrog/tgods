package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type bucket[T any] interface {
	//Save puts the elem into the bucket, returns:
	//  bucket - the either this or a changed bucket for performance based on size of the elem
	//  T - the old elem
	//  bool - if found an existing elem by the eq func
	Save(elem T, eq types.Equal[T]) (bucket[T], T, bool)
	Get(elem T, eq types.Equal[T]) (T, bool)               // finds and returns an elem by given eq func and input elem
	Delete(elem T, eq types.Equal[T]) (bucket[T], T, bool) // removes the elem from the bucket, return the elem and true if found
	Contains(elem T, eq types.Equal[T]) bool               // checks if the elem is in the bucket
}

// assert type
var _ bucket[int] = (*slNode[int])(nil)

type slNode[T any] struct {
	v T
	n *slNode[T]
}

func newLinkedListBucket[T any]() *slNode[T] {
	return &slNode[T]{utils.Nil[T](), nil}
}

func (n *slNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], T, bool) {
	if n == nil {
		return &slNode[T]{elem, nil}, utils.Nil[T](), false
	}
	h := n
	var np *slNode[T]
	for n != nil {
		if eq(elem, n.v) {
			return h, n.v, true
		}
		np = n
		n = n.n
	}
	np.n = &slNode[T]{elem, nil}
	return h, utils.Nil[T](), false
}

func (n *slNode[T]) Get(elem T, eq types.Equal[T]) (T, bool) {
	for n != nil {
		if eq(elem, n.v) {
			return n.v, true
		}
		n = n.n
	}
	return utils.Nil[T](), false
}

func (n *slNode[T]) Delete(elem T, eq types.Equal[T]) (bucket[T], T, bool) {
	if n == nil {
		return nil, utils.Nil[T](), false
	}
	if eq(elem, n.v) {
		v := n.v
		return nil, v, true
	}
	h := n
	for n.n != nil {
		v := n.n.v
		if eq(elem, v) {
			n.n = n.n.n
			return h, v, true
		}
	}
	return h, utils.Nil[T](), false
}

func (n *slNode[T]) Contains(elem T, eq types.Equal[T]) bool {
	for n != nil {
		if eq(elem, n.v) {
			return true
		}
		n = n.n
	}
	return false
}
