package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

// NewTreeSetOf creates a red black tree backed SortedSet with init values
func NewTreeSetOf[T constraints.Ordered](values ...T) types.SortedSet[T] {
	return newRbTreeOf(values...)
}

// NewTreeSetOfComp creates a red black tree backed SortedSet with a Compare func
func NewTreeSetOfComp[T any](comp types.Compare[T]) types.SortedSet[T] {
	return newRbTreeOfComp(comp)
}

// NewHashSet creates a hash table backed set with custom Hash and Equal functions
func NewHashSet[T any](hs types.Hash[T], eq types.Equal[T]) types.Set[T] {
	return newHashTable[T](hs, eq)
}

// NewHashSetC creates a hash table with a constrained type that implements custom Hash and Equal
func NewHashSetC[T types.HashAndEqual[T]]() types.Set[T] {
	hs := func(elem T) uint { return elem.Hash() }
	eq := func(a, b T) bool { return a.Equal(b) }
	return newHashTable[T](hs, eq)
}

// NewHashSetOfNum creates a hash table backed set with values, see funcs.NumHash
func NewHashSetOfNum[T constraints.Integer | constraints.Float](values ...T) types.Set[T] {
	h := newHashTable[T](funcs.NumHash[T], funcs.ValueEqual[T])
	for _, v := range values {
		h.Add(v)
	}
	return h
}

// NewHashSetOfStr creates a hash table backed set with values, see funcs.NewStrHash
func NewHashSetOfStr(values ...string) types.Set[string] {
	h := newHashTable[string](funcs.NewStrHash(), funcs.ValueEqual[string])
	for _, v := range values {
		h.Add(v)
	}
	return h
}

// NewLinkedHashSet creates a linked hash table backed set with custom Hash and Equal functions
func NewLinkedHashSet[T any](hs types.Hash[T], eq types.Equal[T]) types.LinkedSet[T] {
	return newLinkedHashTable[T](hs, eq, 0)
}

// NewLinkedHashSetOfNum creates a linked hash table backed set with values, see funcs.NumHash
func NewLinkedHashSetOfNum[T constraints.Integer | constraints.Float](values ...T) types.LinkedSet[T] {
	h := newLinkedHashTable[T](funcs.NumHash[T], funcs.ValueEqual[T], 0)
	for _, v := range values {
		h.Add(v)
	}
	return h
}

// NewLinkedHashSetOfStr creates a linked hash table backed set with values, see funcs.NewStrHash
func NewLinkedHashSetOfStr(values ...string) types.LinkedSet[string] {
	h := newLinkedHashTable[string](funcs.NewStrHash(), funcs.ValueEqual[string], 0)
	for _, v := range values {
		h.Add(v)
	}
	return h
}

// NewEnumSet creates an array based enum set
//
//	max - the max value, which defines the cap of the array
func NewEnumSet[T constraints.Integer](max T, values ...T) types.SortedSet[T] {
	return newEnumSet(max, values...)
}
