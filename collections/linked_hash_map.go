package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type linkedHashMap[K any, V any] struct {
	linkedHashTable[types.Entry[K, V]]
	limit int // the maximum size limit, 0 means unlimited
}

func newLinkedHashMap[K any, V any](hs types.Hash[K], eq types.Equal[K], sizeLimit int, accessOrder AccessOrder) *linkedHashMap[K, V] {
	hhs := func(a types.Entry[K, V]) uint { return hs(a.Key()) }
	heq := func(a, b types.Entry[K, V]) bool { return eq(a.Key(), b.Key()) }
	h := newLinkedHashTable[types.Entry[K, V]](hhs, heq, accessOrder)
	return &linkedHashMap[K, V]{*h, sizeLimit}
}

func (h *linkedHashMap[K, V]) Get(k K) (V, bool) {
	n := h.linkedHashTable.getNode(keyEntry[K, V]{k})
	if n != nil {
		if h.accessOrder&GetOrder > 0 {
			x := n.External()
			h.linkedHashTable.removeNode(x)
			h.linkedHashTable.appendToTail(x)
		}
		return n.Value().Value(), true
	}
	return utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) Put(k K, v V) (V, bool) {
	_, old, found := h.linkedHashTable.add(EntryOf[K, V](k, v))
	if h.limit > 0 && h.linkedHashTable.size > h.limit {
		head := h.head
		h.head = head.Next()
		removeNodeFromList(head)
		h.linkedHashTable.remove(keyEntry[K, V]{head.Value().Key()})
	}
	if found {
		return old.Value(), true
	}
	return utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) Remove(k K) (V, bool) {
	n := h.linkedHashTable.remove(keyEntry[K, V]{k})
	if n != nil {
		return n.Value().Value(), true
	}
	return utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) ContainsKey(k K) bool {
	return h.linkedHashTable.Contains(keyEntry[K, V]{k})
}

func (h *linkedHashMap[K, V]) PutHead(k K, v V) (V, bool) {
	e, found := h.linkedHashTable.AddHead(EntryOf[K, V](k, v))
	if found {
		return e.Value(), true
	}
	return utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) RemoveHead() (K, V, bool) {
	e, found := h.linkedHashTable.RemoveHead()
	if found {
		return e.Key(), e.Value(), true
	}
	return utils.Nil[K](), utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) Head() (K, V, bool) {
	e, found := h.linkedHashTable.Head()
	if found {
		return e.Key(), e.Value(), true
	}
	return utils.Nil[K](), utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) PutTail(k K, v V) (V, bool) {
	e, found := h.linkedHashTable.AddTail(EntryOf(k, v))
	if found {
		return e.Value(), true
	}
	return utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) RemoveTail() (K, V, bool) {
	e, found := h.linkedHashTable.RemoveTail()
	if found {
		return e.Key(), e.Value(), true
	}
	return utils.Nil[K](), utils.Nil[V](), false
}

func (h *linkedHashMap[K, V]) Tail() (K, V, bool) {
	e, found := h.linkedHashTable.Tail()
	if found {
		return e.Key(), e.Value(), true
	}
	return utils.Nil[K](), utils.Nil[V](), false
}
