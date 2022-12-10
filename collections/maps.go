package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

// NewTreeMapOf creates a tree map with init values.
func NewTreeMapOf[K constraints.Ordered, V any](entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	return NewTreeMapOfComp(funcs.ValueCompare[K], entries...)
}

// NewTreeMapOfComp creates a tree map with custom Compare func and init values.
func NewTreeMapOfComp[K any, V any](comp types.Compare[K], entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	m := newTreeMapOfComp[K, V](comp)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewHashMapC creates a hash map with a constrained key type that implements custom Hash and Equal
func NewHashMapC[K types.WithHashAndEqual[K], V any]() types.Map[K, V] {
	hs := func(key K) uint { return key.Hash() }
	eq := func(a, b K) bool { return a.Equal(b) }
	return newHashMap[K, V](hs, eq)
}

// NewHashMapOf creates a hash map with custom Hash and Equal functions, and init values.
func NewHashMapOf[K any, V any](hs types.Hash[K], eq types.Equal[K], entries ...types.Entry[K, V]) types.Map[K, V] {
	m := newHashMap[K, V](hs, eq)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewHashMapOfStrKey creates a hash map with key type as string and init values.
func NewHashMapOfStrKey[V any](entries ...types.Entry[string, V]) types.Map[string, V] {
	m := newHashMap[string, V](funcs.NewStrHash(), funcs.ValueEqual[string])
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewHashMapOfNumKey creates a hash map with key type as constraints.Integer | constraints.Float, and init values.
func NewHashMapOfNumKey[K constraints.Integer | constraints.Float, V any](entries ...types.Entry[K, V]) types.Map[K, V] {
	m := newHashMap[K, V](funcs.NumHash[K], funcs.ValueEqual[K])
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewLinkedHashMap creates a linked hash map with custom Hash and Equal functions, and init values.
//
//	hs - the key hash function.
//	eq - the key equal function.
//	sizeLimit - limit the maximum size of elements, extra elements will be removed upon Put based on element order defined by AccessOrder. 0 means unlimited.
//	accessOrder - defines how elements are ordered. For an LRU cache, you can provide PutOrder + GetOrder
func NewLinkedHashMap[K any, V any](hs types.Hash[K], eq types.Equal[K], sizeLimit int, accessOrder AccessOrder) types.LinkedMap[K, V] {
	return newLinkedHashMap[K, V](hs, eq, sizeLimit, accessOrder)
}

// NewLinkedHashMapC creates a linked hash map with a constrained key type that implements custom Hash and Equal, and init values.
// 'C' stands for Client Customized Constrained type.
//
//	sizeLimit - limit the maximum size of elements, extra elements will be removed upon Put based on element order defined by AccessOrder. 0 means unlimited.
//	accessOrder - defines how elements are ordered. For an LRU cache, you can provide PutOrder + GetOrder
func NewLinkedHashMapC[K types.WithHashAndEqual[K], V any](sizeLimit int, accessOrder AccessOrder) types.LinkedMap[K, V] {
	hs := func(key K) uint { return key.Hash() }
	eq := func(a, b K) bool { return a.Equal(b) }
	return newLinkedHashMap[K, V](hs, eq, sizeLimit, accessOrder)
}

// NewLinkedHashMapOf creates a linked hash map with custom Hash and Equal functions, and init values.
// No size limit. Iteration will be of the original put order.
func NewLinkedHashMapOf[K any, V any](hs types.Hash[K], eq types.Equal[K], entries ...types.Entry[K, V]) types.LinkedMap[K, V] {
	m := newLinkedHashMap[K, V](hs, eq, 0, OriginalOrder)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewLinkedHashMapOfC creates a linked hash map with a constrained key type that implements custom Hash and Equal, and init values.
// 'C' stands for Client Customized Constrained type.
// No size limit. Iteration will be of the original put order.
func NewLinkedHashMapOfC[K types.WithHashAndEqual[K], V any](entries ...types.Entry[K, V]) types.LinkedMap[K, V] {
	hs := func(key K) uint { return key.Hash() }
	eq := func(a, b K) bool { return a.Equal(b) }
	m := newLinkedHashMap[K, V](hs, eq, 0, OriginalOrder)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewLinkedHashMapOfStrKey creates a linked hash map with key type as string and init values.
// No size limit. Iteration will be of the original put order.
func NewLinkedHashMapOfStrKey[V any](entries ...types.Entry[string, V]) types.LinkedMap[string, V] {
	m := newLinkedHashMap[string, V](funcs.NewStrHash(), funcs.ValueEqual[string], 0, OriginalOrder)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewLinkedHashMapOfNumKey creates a linked hash map with key type as constraints.Integer | constraints.Float, and init values.
// No size limit. Iteration will be of the original put order.
func NewLinkedHashMapOfNumKey[K constraints.Integer | constraints.Float, V any](entries ...types.Entry[K, V]) types.LinkedMap[K, V] {
	m := newLinkedHashMap[K, V](funcs.NumHash[K], funcs.ValueEqual[K], 0, OriginalOrder)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}

// NewLRUCache alias for NewLinkedHashMap
func NewLRUCache[K any, V any](hs types.Hash[K], eq types.Equal[K], sizeLimit int, accessOrder AccessOrder) types.Map[K, V] {
	return NewLinkedHashMap[K, V](hs, eq, sizeLimit, accessOrder)
}

// NewLRUCacheC alias for NewLinkedHashMapC
func NewLRUCacheC[K types.WithHashAndEqual[K], V any](sizeLimit int, accessOrder AccessOrder) types.Map[K, V] {
	return NewLinkedHashMapC[K, V](sizeLimit, accessOrder)
}

// NewLRUCacheOfStrKey creates a linkedHashMap as an LRU cache, eviction and iteration will follow provided AccessOrder
func NewLRUCacheOfStrKey[V any](sizeLimit int, accessOrder AccessOrder) types.Map[string, V] {
	return NewLinkedHashMap[string, V](funcs.NewStrHash(), funcs.ValueEqual[string], sizeLimit, accessOrder)
}

// NewLRUCacheOfNumKey creates a linkedHashMap as an LRU cache, eviction and iteration will follow provided AccessOrder
func NewLRUCacheOfNumKey[K constraints.Integer | constraints.Float, V any](sizeLimit int, accessOrder AccessOrder) types.Map[K, V] {
	return NewLinkedHashMap[K, V](funcs.NumHash[K], funcs.ValueEqual[K], sizeLimit, accessOrder)
}

// NewEnumMap creates an array based enum map
//
//	max - the max value, which defines the cap of the array
func NewEnumMap[K constraints.Integer, V any](max K, entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	return newEnumMap(max, entries...)
}
