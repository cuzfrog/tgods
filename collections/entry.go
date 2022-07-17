package collections

import "github.com/cuzfrog/tgods/types"

type mapEntry[K any, V any] struct {
	k K
	v V
}

func EntryOf[K any, V any](k K, v V) types.Entry[K, V] {
	return &mapEntry[K, V]{k, v}
}

func (e *mapEntry[K, V]) Key() K {
	return e.k
}

func (e *mapEntry[K, V]) Value() V {
	return e.v
}

type keyEntry[K any, V any] struct {
	k K
}

func (e keyEntry[K, V]) Key() K {
	return e.k
}

func (e keyEntry[K, V]) Value() V {
	panic("Unsupported operation.")
}
