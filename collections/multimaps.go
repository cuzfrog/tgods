package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"golang.org/x/exp/constraints"
)

// NewArrayListMultiMapOf creates a hash multimap with custom Hash and Equal functions, and init values.
func NewArrayListMultiMapOf[K any, V any](hs types.Hash[K], eq types.Equal[K], entries ...types.Entry[K, types.Collection[V]]) types.MultiMap[K, V] {
	mm := newArrayListMultiMap[K, V](hs, eq)
	utils.AddSliceTo[types.Entry[K, types.Collection[V]]](entries, mm)
	return mm
}

// NewArrayListMultiMapOfStrKey creates a hash multimap with key type as string, and init values.
func NewArrayListMultiMapOfStrKey[V any](entries ...types.Entry[string, types.Collection[V]]) types.MultiMap[string, V] {
	mm := newArrayListMultiMap[string, V](funcs.NewStrHash(), funcs.ValueEqual[string])
	utils.AddSliceTo[types.Entry[string, types.Collection[V]]](entries, mm)
	return mm
}

// NewArrayListMultiMapOfNumKey creates a hash multimap with key type as constraints.Integer | constraints.Float, and init values.
func NewArrayListMultiMapOfNumKey[K constraints.Integer | constraints.Float, V any](entries ...types.Entry[K, types.Collection[V]]) types.MultiMap[K, V] {
	mm := newArrayListMultiMap[K, V](funcs.NumHash[K], funcs.ValueEqual[K])
	utils.AddSliceTo[types.Entry[K, types.Collection[V]]](entries, mm)
	return mm
}
