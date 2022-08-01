package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type internalMap[K any, V any] interface {
	Size() int
	Hash(k K) uint
	Equal(a, b types.Entry[K, V]) bool
	Buckets() []bucket[types.Entry[K, V]]
}

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

func (h *hashMap[K, V]) Hash(k K) uint {
	return h.khs(k)
}
func (h *hashMap[K, V]) Equal(a, b types.Entry[K, V]) bool {
	return h.eq(a, b)
}
func (h *hashMap[K, V]) Buckets() []bucket[types.Entry[K, V]] {
	return h.arr
}

func (h *hashMap[K, V]) Get(k K) (V, bool) {
	return getValueFromMap[K, V](h, k)
}
func getValueFromMap[K any, V any](m internalMap[K, V], k K) (V, bool) {
	if m.Size() == 0 {
		return utils.Nil[V](), false
	}
	arr := m.Buckets()
	i := hashToIndex(m.Hash(k), cap(arr))
	b := arr[i]
	n := findNodeFromBucket[types.Entry[K, V]](b, keyEntry[K, V]{k}, m.Equal)
	if n != nil {
		return n.Value().Value(), true
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
	n := h.h.remove(keyEntry[K, V]{k})
	if n != nil {
		return n.Value().Value(), true
	}
	return utils.Nil[V](), false
}

func (h *hashMap[K, V]) ContainsKey(k K) bool {
	return h.h.Contains(keyEntry[K, V]{k})
}
