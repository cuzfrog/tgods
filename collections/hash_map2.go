package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type hashMap[K any, V any] struct {
	hashTable[types.Entry[K, V]]
}

func newHashMap[K any, V any](hs types.Hash[K], eq types.Equal[K]) *hashMap[K, V] {
	hhs := func(a types.Entry[K, V]) uint { return hs(a.Key()) }
	heq := func(a, b types.Entry[K, V]) bool { return eq(a.Key(), b.Key()) }
	h := newHashTable[types.Entry[K, V]](hhs, heq)
	return &hashMap[K, V]{*h}
}

func (h *hashMap[K, V]) Get(k K) (V, bool) {
	n := h.hashTable.getNode(keyEntry[K, V]{k})
	if n != nil {
		return n.Value().Value(), true
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) Put(k K, v V) (V, bool) {
	_, e, found := h.hashTable.add(EntryOf(k, v))
	if found {
		return e.Value(), found
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) Remove(k K) (V, bool) {
	n := h.hashTable.remove(keyEntry[K, V]{k})
	if n != nil {
		return n.Value().Value(), true
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) ContainsKey(k K) bool {
	return h.hashTable.Contains(keyEntry[K, V]{k})
}
