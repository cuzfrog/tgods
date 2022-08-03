package collections

import (
	"github.com/cuzfrog/tgods/types"
)

type linkedHashTable[T any] struct {
	hashTable[T]
	head        node[T]
	tail        node[T] // new element is added to
	accessOrder AccessOrder
}

func newLinkedHashTable[T any](hs types.Hash[T], eq types.Equal[T], accessOrder AccessOrder) *linkedHashTable[T] {
	h := newHashTableOfSlxNode(hs, eq)
	return &linkedHashTable[T]{*h, nil, nil, accessOrder}
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
		n, old, found = h.hashTable.add(elem)
		n.SetExternal(x)
	} else {
		n, old, found = h.hashTable.add(elem)
		if found {
			if h.accessOrder&PutOrder > 0 {
				x := n.External()
				h.removeNode(x)
				h.appendToTail(x)
			}
		} else {
			x := newDlNode(elem, h.tail, nil)
			n.SetExternal(x)
			h.appendToTail(x)
		}
	}
	return
}
func (h *linkedHashTable[T]) appendToTail(x node[T]) {
	if h.tail == nil {
		h.tail = x
		h.head = x
	} else {
		x.SetPrev(h.tail)
		h.tail.SetNext(x)
		h.tail = x
	}
}

func (h *linkedHashTable[T]) Remove(elem T) bool {
	n := h.remove(elem)
	return n != nil
}
func (h *linkedHashTable[T]) remove(elem T) node[T] {
	if h.size == 0 {
		return nil
	}
	n := h.hashTable.remove(elem)
	if n != nil {
		h.removeNode(n.External())
	}
	return n
}

func (h *linkedHashTable[T]) removeNode(x node[T]) {
	if h.head == x {
		h.head = x.Next()
	}
	if h.tail == x {
		h.tail = x.Prev()
	}
	removeNodeFromList(x)
}

func (h *linkedHashTable[T]) Clear() {
	h.hashTable.Clear()
	h.head = nil
	h.tail = nil
}
