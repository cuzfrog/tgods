package collections

import "github.com/cuzfrog/tgods/types"

func NewArrayMultiMap[K any, V any](hs types.Hash[K], eq types.Equal[K]) types.ListMultiMap[K, V] {
	panic("not implemented")
	//return newArrayListMultiMap[K, V](hs, eq)
}
