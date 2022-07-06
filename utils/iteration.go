package utils

import "github.com/cuzfrog/tgods/core"

func Slice[T any](it core.Iterator[T], size int) []T {
	arr := make([]T, size)
	for it.Next() {
		arr[it.Index()] = it.Value()
	}
	return arr
}
