package collections

import (
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
