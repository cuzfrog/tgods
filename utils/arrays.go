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
