package utils

import (
	"github.com/cuzfrog/tgods/types"
)

// SliceFrom creates a new slice from the Collection,
//   len(slice) = min(size, elemCntFromIterator)
//   cap(slice) = size
func SliceFrom[T any](col types.Collection[T]) []T {
	it, size := col.Iterator(), col.Size()
	return SliceFromIt(it, size)
}

// SliceFromIt creates a new slice from the Iterator,
func SliceFromIt[T any](it types.Iterator[T], size int) []T {
	arr := make([]T, size)
	for it.Next() {
		i := it.Index()
		arr[i] = it.Value()
	}
	return arr
}

// SliceProject creates a new slice from the Collection and projects the elem by given mapping function.
//   len(slice) = min(size, elemCntFromIterator)
//   cap(slice) = size
func SliceProject[T any, R any](col types.Collection[T], fn func(t T) R) []R {
	it, size := col.Iterator(), col.Size()
	return SliceProjectIt(it, size, fn)
}

// SliceProjectIt creates a new slice from the Iterator and projects the elem by given mapping function.
func SliceProjectIt[T any, R any](it types.Iterator[T], size int, fn func(t T) R) []R {
	arr := make([]R, size)
	for it.Next() {
		i := it.Index()
		arr[i] = fn(it.Value())
	}
	return arr
}
