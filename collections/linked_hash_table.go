package collections

import (
	"github.com/cuzfrog/tgods/types"
)

type linkedHashTable[T any] struct {
	*hashTable[T]
	head  node[T]
	tail  node[T] // new element is added to
	limit int     // the maximum size limit, 0 means unlimited
}

func newLinkedHashTable[T any](hs types.Hash[T], eq types.Equal[T]) *linkedHashTable[T] {
	return &linkedHashTable[T]{newHashTableOfSlxNode(hs, eq), nil, nil, 0}
}

func (h *linkedHashTable[T]) Add(elem T) bool {
	if h.tail == nil {
		h.tail = newDlNode(elem, nil, nil)
		h.head = h.tail
		x := h.tail
		n, _, _ := h.add(elem)
		n.SetExternal(x)
	} else {
		n, _, found := h.add(elem)
		if found {
			x := n.External()
			removeNodeFromList(x)
			x.SetPrev(h.tail)
			h.tail.SetNext(x)
		} else {
			x := newDlNode(elem, h.tail, nil)
			h.tail.SetNext(x)
			n.SetExternal(x)
		}
		h.tail = h.tail.Next()
	}
	return true
}

func (h *linkedHashTable[T]) Remove(elem T) bool {
	if h.size == 0 {
		return false
	}
	n := h.remove(elem)
	if n != nil {
		x := n.External()
		removeNodeFromList(x)
	}
	return n != nil
}
