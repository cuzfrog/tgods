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

// NewHashSetMultiMapOf creates a hash multimap with custom Hash and Equal functions, and init values.
func NewHashSetMultiMapOf[K any, V any](khs types.Hash[K], keq types.Equal[K], vhs types.Hash[V], veq types.Equal[V],
	entries ...types.Entry[K, types.Collection[V]]) types.MultiMap[K, V] {
	mm := newHashSetMultiMap[K, V](khs, keq, vhs, veq)
	utils.AddSliceTo[types.Entry[K, types.Collection[V]]](entries, mm)
	return mm
}

// NewHashSetMultiMapOfStrKey creates a hash multimap with key type as string, and init values.
func NewHashSetMultiMapOfStrKey[V any](vhs types.Hash[V], veq types.Equal[V], entries ...types.Entry[string, types.Collection[V]]) types.MultiMap[string, V] {
	mm := newHashSetMultiMap[string, V](funcs.NewStrHash(), funcs.ValueEqual[string], vhs, veq)
	utils.AddSliceTo[types.Entry[string, types.Collection[V]]](entries, mm)
	return mm
}

// NewHashSetMultiMapOfNumKey creates a hash multimap with key type as constraints.Integer | constraints.Float, and init values.
func NewHashSetMultiMapOfNumKey[K constraints.Integer | constraints.Float, V any](vhs types.Hash[V], veq types.Equal[V], entries ...types.Entry[K, types.Collection[V]]) types.MultiMap[K, V] {
	mm := newHashSetMultiMap[K, V](funcs.NumHash[K], funcs.ValueEqual[K], vhs, veq)
	utils.AddSliceTo[types.Entry[K, types.Collection[V]]](entries, mm)
	return mm
}
