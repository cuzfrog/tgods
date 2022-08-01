package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type hashMap[K any, V any] struct {
	*hashTable[types.Entry[K, V]]
	h   *hashTable[types.Entry[K, V]] // point to the same hashTable, extra reference to workaround method type
	khs types.Hash[K]
}

func newHashMap[K any, V any](hs types.Hash[K], eq types.Equal[K]) *hashMap[K, V] {
	hhs := func(a types.Entry[K, V]) uint { return hs(a.Key()) }
	heq := func(a, b types.Entry[K, V]) bool { return eq(a.Key(), b.Key()) }
	h := newHashTable[types.Entry[K, V]](hhs, heq)
	return &hashMap[K, V]{h, h, hs}
}

func (h *hashMap[K, V]) Get(k K) (V, bool) {
	if h.size == 0 {
		return utils.Nil[V](), false
	}
	i := hashToIndex(h.khs(k), cap(h.arr))
	b := h.arr[i]
	if b == nil {
		return utils.Nil[V](), false
	}
	e, found := b.Get(keyEntry[K, V]{k}, h.eq)
	if found {
		return e.Value(), found
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) Put(k K, v V) (V, bool) {
	_, e, found := h.h.add(EntryOf(k, v))
	if found {
		return e.Value(), found
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) Remove(k K) (V, bool) {
	e, found := h.h.remove(keyEntry[K, V]{k})
	if found {
		return e.Value(), found
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) ContainsKey(k K) bool {
	return h.h.Contains(keyEntry[K, V]{k})
}
