package utils

import (
	"github.com/cuzfrog/tgods/core"
)

// SliceFrom creates a new slice from the Iterator,
// len(slice) = min(size, elemCntFromIterator)
// cap(slice) = size
func SliceFrom[T any](it core.Iterator[T], size int) []T {
	arr := make([]T, size)
	var i int
	for it.Next() {
		i = it.Index()
		if i >= size {
			break
		}
		arr[i] = it.Value()
	}
	return arr[:Min(i+1, size)]
}
