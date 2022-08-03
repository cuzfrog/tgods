package collections

import (
	"github.com/cuzfrog/tgods/types"
)

type internalMap[K any, V any] interface {
	Size() int
	Hash(k K) uint
	Equal(a, b types.Entry[K, V]) bool
	Buckets() []bucket[types.Entry[K, V]]
}

func getNodeFromMap[K any, V any](m internalMap[K, V], k K) node[types.Entry[K, V]] {
	if m.Size() == 0 {
		return nil
	}
	arr := m.Buckets()
	i := hashToIndex(m.Hash(k), cap(arr))
	b := arr[i]
	return findNodeFromBucket[types.Entry[K, V]](b, keyEntry[K, V]{k}, m.Equal)
}
