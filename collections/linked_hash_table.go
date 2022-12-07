package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
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

func (h *linkedHashTable[T]) Replace(elem T) (T, bool) {
	_, old, found := h.add(elem)
	return old, found
}

// add returns the newly added or existing node
func (h *linkedHashTable[T]) add(elem T) (n node[T], old T, found bool) {
	if h.tail == nil {
		return h.addFirstNode(elem)
	}
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
	return
}

// should only be called when size == 0
func (h *linkedHashTable[T]) addFirstNode(elem T) (n node[T], old T, found bool) {
	h.tail = newDlNode(elem, nil, nil)
	h.head = h.tail
	x := h.tail
	n, old, found = h.hashTable.add(elem)
	n.SetExternal(x)
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

func (h *linkedHashTable[T]) prependToHead(x node[T]) {
	if h.head == nil {
		h.tail = x
		h.head = x
	} else {
		x.SetNext(h.head)
		h.head.SetPrev(x)
		h.head = x
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

func (h *linkedHashTable[T]) AddHead(elem T) (T, bool) {
	n, old, found := h.hashTable.add(elem)
	if found {
		x := n.External()
		h.removeNode(x)
		h.prependToHead(x)
	} else {
		x := newDlNode(elem, nil, h.head)
		n.SetExternal(x)
		h.prependToHead(x)
	}
	return old, found
}

func (h *linkedHashTable[T]) RemoveHead() (T, bool) {
	if h.head == nil {
		return utils.Nil[T](), false
	}
	elem := h.head.Value()
	h.remove(elem)
	return elem, true
}

func (h *linkedHashTable[T]) Head() (T, bool) {
	if h.head == nil {
		return utils.Nil[T](), false
	}
	return h.head.Value(), true
}

func (h *linkedHashTable[T]) RemoveTail() (T, bool) {
	if h.tail == nil {
		return utils.Nil[T](), false
	}
	elem := h.tail.Value()
	h.remove(elem)
	return elem, true
}

// AddTail equivalent to Add with different return types
func (h *linkedHashTable[T]) AddTail(elem T) (T, bool) {
	_, old, found := h.add(elem)
	return old, found
}

func (h *linkedHashTable[T]) Tail() (T, bool) {
	if h.tail == nil {
		return utils.Nil[T](), false
	}
	return h.tail.Value(), true
}
