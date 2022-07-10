package utils

import (
	"github.com/cuzfrog/tgods/types"
)

// SliceFrom creates a new slice from the Iterator,
// len(slice) = min(size, elemCntFromIterator)
// cap(slice) = size
func SliceFrom[T any](col types.Collection[T]) []T {
	it, size := col.Iterator(), col.Size()
	arr := make([]T, size)
	for it.Next() {
		i := it.Index()
		arr[i] = it.Value()
	}
	return arr
}

// Shuffle redistributes elems in the slice a using Knuth algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func Shuffle[T any](a []T, randFunc func(n int) int) {
	for i := len(a) - 1; i >= 1; i-- {
		j := randFunc(i)
		a[i], a[j] = a[j], a[i]
	}
}
