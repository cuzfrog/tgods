package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

// NewTreeMapOf creates a tree map with init values
func NewTreeMapOf[K constraints.Ordered, V any](entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	return NewTreeMapOfComp(funcs.ValueCompare[K], entries...)
}

// NewTreeMapOfComp creates a tree map with custom Compare func and init values
func NewTreeMapOfComp[K any, V any](comp types.Compare[K], entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	m := newTreeMapOfComp[K, V](comp)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewHashMapOf creates a hash map with custom Hash and Equal functions, and init values
func NewHashMapOf[K any, V any](hs types.Hash[K], eq types.Equal[K], entries ...types.Entry[K, V]) types.Map[K, V] {
	m := newHashMap[K, V](hs, eq)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewHashMapOfStrKey creates a hash map with key type as string and init values
func NewHashMapOfStrKey[V any](entries ...types.Entry[string, V]) types.Map[string, V] {
	m := newHashMap[string, V](funcs.NewStrHash(), funcs.ValueEqual[string])
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewHashMapOfNumKey creates a hash map with key type as constraints.Integer | constraints.Float, and init values
func NewHashMapOfNumKey[K constraints.Integer | constraints.Float, V any](entries ...types.Entry[K, V]) types.Map[K, V] {
	m := newHashMap[K, V](funcs.NumHash[K], funcs.ValueEqual[K])
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}
