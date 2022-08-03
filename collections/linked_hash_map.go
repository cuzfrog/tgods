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

func (h *linkedHashMap[K, V]) Hash(k K) uint {
	return h.hs(keyEntry[K, V]{k})
}
func (h *linkedHashMap[K, V]) Equal(a, b types.Entry[K, V]) bool {
	return h.eq(a, b)
}
func (h *linkedHashMap[K, V]) Buckets() []bucket[types.Entry[K, V]] {
	return h.arr
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
	if h.limit > 0 && h.size > h.limit {
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
