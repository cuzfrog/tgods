package collections

import (
	"github.com/cuzfrog/tgods/types"
)

// NewCircularArrayList creates an auto expandable circular array based list, auto shrinkable, but will not shrink if the length is <= defaultInitSize,
// the underlying array will be lazily created unless init values are provided, the init arr size is the same as init values'
func NewCircularArrayList[T comparable](values ...T) types.ArrayList[T] {
	return newCircularArrayOf[T](values...).withRole(list)
}

// NewCircularArrayListOfSize creates underlying array eagerly with the init size
func NewCircularArrayListOfSize[T comparable](initSize int) types.ArrayList[T] {
	return newCircularArray[T](initSize).withRole(list)
}

// NewCircularArrayListOfEq creates underlying array eagerly with the init size
func NewCircularArrayListOfEq[T any](initSize int, comp types.Equal[T]) types.ArrayList[T] {
	return newCircularArrayOfEq(initSize, comp).withRole(list)
}

func NewLinkedList[T comparable](values ...T) types.LinkedList[T] {
	return newLinkedListOf[T](values...).withRole(list)
}

func NewLinkedListOfEq[T any](equal types.Equal[T], values ...T) types.LinkedList[T] {
	return newLinkedListOfEq[T](equal, values...).withRole(list)
}
