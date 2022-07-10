package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

// NewCircularArrayListOf creates an auto expandable circular array based list, auto shrinkable, but will not shrink if the length is <= defaultInitSize,
// the underlying array will be lazily created unless init values are provided, the init arr size is the same as init values'
func NewCircularArrayListOf[T comparable](values ...T) types.ArrayList[T] {
	return newCircularArrayOf[T](values...).withRole(list)
}

// NewCircularArrayList creates underlying array eagerly with the init size
func NewCircularArrayList[T comparable](initSize int) types.ArrayList[T] {
	return newCircularArray[T](initSize).withRole(list)
}

// NewCircularArrayListOfEq creates underlying array eagerly with the init size
func NewCircularArrayListOfEq[T any](initSize int, comp funcs.Equal[T]) types.ArrayList[T] {
	return newCircularArrayOfEq(initSize, comp).withRole(list)
}

func NewLinkedListOf[T comparable](values ...T) types.List[T] {
	return newLinkedListOf[T](values...).withRole(list)
}

func NewLinkedListOfEq[T any](equal funcs.Equal[T]) types.List[T] {
	return newLinkedListOfEq[T](equal).withRole(list)
}
