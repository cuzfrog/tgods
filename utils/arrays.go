package utils

import (
	"github.com/cuzfrog/tgods/types"
)

// SliceFrom creates a new slice from the Collection,
//   len(slice) = min(size, elemCntFromIterator)
//   cap(slice) = size
func SliceFrom[T any](col types.Collection[T]) []T {
	it, size := col.Iterator(), col.Size()
	return SliceFromIterator(it, size)
}

// SliceFromIterator creates a new slice from the Iterator,
func SliceFromIterator[T any](it types.Iterator[T], size int) []T {
	arr := make([]T, size)
	for it.Next() {
		i := it.Index()
		arr[i] = it.Value()
	}
	return arr
}
