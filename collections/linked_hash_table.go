package collections

import (
	"github.com/cuzfrog/tgods/types"
)

type linkedHashTable[T any] struct {
	*hashTable[T]
	h           *hashTable[T]
	head        node[T]
	tail        node[T] // new element is added to
	accessOrder byte    // 1 - put order, 2 - get order, 3 - both get and put order
}

func newLinkedHashTable[T any](hs types.Hash[T], eq types.Equal[T], accessOrder byte) *linkedHashTable[T] {
	h := newHashTableOfSlxNode(hs, eq)
	return &linkedHashTable[T]{h, h, nil, nil, accessOrder}
}
func (h *linkedHashTable[T]) Add(elem T) bool {
	h.add(elem)
	return true
}

// add returns the newly added or existing node
func (h *linkedHashTable[T]) add(elem T) (n node[T], old T, found bool) {
	if h.tail == nil {
		h.tail = newDlNode(elem, nil, nil)
		h.head = h.tail
		x := h.tail
		n, old, found = h.h.add(elem)
		n.SetExternal(x)
	} else {
		n, old, found = h.h.add(elem)
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
	return
}
func (h *linkedHashTable[T]) Remove(elem T) bool {
	n := h.remove(elem)
	return n != nil
}
func (h *linkedHashTable[T]) remove(elem T) node[T] {
	if h.size == 0 {
		return nil
	}
	n := h.h.remove(elem)
	if n != nil {
		x := n.External()
		if h.head == x {
			h.head = x.Next()
		}
		if h.tail == x {
			h.tail = x.Prev()
		}
		removeNodeFromList(x)
	}
	return n
}

func (h *linkedHashTable[T]) Clear() {
	h.h.Clear()
	h.head = nil
	h.tail = nil
}
