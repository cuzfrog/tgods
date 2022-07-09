package lists

import "github.com/cuzfrog/tgods/core"

// NewCircularArrayListOf creates an auto expandable circular array based list, auto shrinkable, but will not shrink if the length is <= DefaultInitSize,
// the underlying array will be lazily created unless init values are provided, the init arr size is the same as init values'
func NewCircularArrayListOf[T comparable](values ...T) *circularArrayList[T] {
	var arr []T
	var size, start int
	length := len(values)
	if length == 0 {
		arr = nil
		size = 0
		start = -1
	} else {
		arr = values
		size = length
		start = 0
	}
	return &circularArrayList[T]{start, size, arr, size, core.EqualComparable[T]}
}

// NewCircularArrayList creates underlying array eagerly with the init size
func NewCircularArrayList[T comparable](size int) *circularArrayList[T] {
	return &circularArrayList[T]{-1, 0, make([]T, size), 0, core.EqualComparable[T]}
}

// NewCircularArrayListOfType creates underlying array eagerly with the init size
func NewCircularArrayListOfType[T any](size int, comp core.Equal[T]) *circularArrayList[T] {
	return &circularArrayList[T]{-1, 0, make([]T, size), 0, comp}
}
