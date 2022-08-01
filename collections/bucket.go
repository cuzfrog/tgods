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
	Iterator() types.Iterator[T]
	External() node[T] // link to another node reference
}

// assert type
var _ bucket[int] = (*slNode[int])(nil)

func newLinkedListBucketOf[T any](v T) *slNode[T] {
	return &slNode[T]{v, nil}
}

func (n *slNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], T, bool) {
	if n == nil {
		return &slNode[T]{elem, nil}, utils.Nil[T](), false
	}
	h := n
	var np node[T]
	var cur node[T] = n
	for cur != nil {
		if eq(elem, cur.Value()) {
			old := cur.Value()
			cur.SetValue(elem)
			return h, old, true
		}
		np = cur
		cur = cur.Next()
	}
	np.SetNext(&slNode[T]{elem, nil})
	return h, utils.Nil[T](), false
}

func (n *slNode[T]) Get(elem T, eq types.Equal[T]) (T, bool) {
	var next node[T] = n
	for next != nil {
		if eq(elem, next.Value()) {
			return next.Value(), true
		}
		next = next.Next()
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
	for n.next != nil {
		v := n.next.Value()
		if eq(elem, v) {
			n.next = n.next.Next()
			return h, v, true
		}
	}
	return h, utils.Nil[T](), false
}

func (n *slNode[T]) Contains(elem T, eq types.Equal[T]) bool {
	var next node[T] = n
	for next != nil {
		if eq(elem, next.Value()) {
			return true
		}
		next = next.Next()
	}
	return false
}
