package collections

import (
	"github.com/cuzfrog/tgods/types"
)

// NewArrayListOf creates an auto expandable/shrinkable circular array based list with init values,
// the underlying array will be lazily created if init values are not provided.
func NewArrayListOf[T comparable](values ...T) types.ArrayList[T] {
	return newCircularArrayOf[T](values...).withRole(list)
}

// NewArrayListOfSize creates an auto expandable/shrinkable circular array based list with init cap
func NewArrayListOfSize[T comparable](initCap int) types.ArrayList[T] {
	return newCircularArray[T](initCap, AutoExpand+AutoShrink).withRole(list)
}

// NewArrayListOfEq creates an auto expandable/shrinkable circular array based list with init cap and an Equal func
func NewArrayListOfEq[T any](initCap int, comp types.Equal[T]) types.ArrayList[T] {
	return newCircularArrayOfEq(initCap, comp).withRole(list)
}

// NewLinkedListOf creates a linked list with init values
func NewLinkedListOf[T comparable](values ...T) types.LinkedList[T] {
	return newLinkedListOf[T](values...).withRole(list)
}

// NewLinkedListOfEq creates a linked list with an Equal func and init values
func NewLinkedListOfEq[T any](equal types.Equal[T], values ...T) types.LinkedList[T] {
	return newLinkedListOfEq[T](equal, values...).withRole(list)
}
