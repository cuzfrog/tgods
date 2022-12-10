package collections

import (
	"github.com/cuzfrog/tgods/types"
)

// NewArrayListOf creates an auto expandable/shrinkable circular array based list with init values,
// the underlying array will be lazily created if init values are not provided.
func NewArrayListOf[T comparable](values ...T) types.ArrayList[T] {
	return newCircularArrayOf[T](values...).withRole(list)
}

// NewArrayListOfC creates an auto expandable/shrinkable circular array based list with init values of a constrained type,
// the underlying array will be lazily created if init values are not provided.
// 'C' stands for Client Customized Constrained type.
func NewArrayListOfC[T types.WithEqual[T]](values ...T) types.ArrayList[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newCircularArrayOfP[T](eq, AutoExpand+AutoShrink, values...).withRole(list)
}

// NewArrayListOfSize creates an auto expandable/shrinkable circular array based list with init cap
func NewArrayListOfSize[T comparable](initCap int) types.ArrayList[T] {
	return newCircularArray[T](initCap, AutoExpand+AutoShrink).withRole(list)
}

// NewArrayListOfSizeC creates an auto expandable/shrinkable circular array based list with init cap
func NewArrayListOfSizeC[T types.WithEqual[T]](initCap int) types.ArrayList[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newCircularArrayOfEq[T](initCap, eq).withRole(list)
}

// NewArrayListOfSizeP creates a circular array based list with init cap
func NewArrayListOfSizeP[T comparable](initCap int, flag AutoSizingFlag) types.ArrayList[T] {
	return newCircularArray[T](initCap, flag).withRole(list)
}

// NewArrayListOfSizePC creates a circular array based list with init cap
func NewArrayListOfSizePC[T types.WithEqual[T]](initCap int, flag AutoSizingFlag) types.ArrayList[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newCircularArrayOfEqP[T](initCap, eq, flag).withRole(list)
}

// NewArrayListOfEq creates an auto expandable/shrinkable circular array based list with init cap and an Equal func
func NewArrayListOfEq[T any](initCap int, eq types.Equal[T]) types.ArrayList[T] {
	return newCircularArrayOfEq(initCap, eq).withRole(list)
}

// NewArrayListOfEqP creates a circular array based list with init cap and an Equal func
func NewArrayListOfEqP[T any](initCap int, eq types.Equal[T], flag AutoSizingFlag) types.ArrayList[T] {
	return newCircularArrayOfEqP(initCap, eq, flag).withRole(list)
}

// NewLinkedListOf creates a linked list with init values
func NewLinkedListOf[T comparable](values ...T) types.LinkedList[T] {
	return newLinkedListOf[T](values...).withRole(list)
}

// NewLinkedListOfC creates a linked list with init values
// 'C' stands for Client Customized Constrained type.
func NewLinkedListOfC[T types.WithEqual[T]](values ...T) types.LinkedList[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newLinkedListOfEq[T](eq, values...).withRole(list)
}

// NewLinkedListOfEq creates a linked list with an Equal func and init values
func NewLinkedListOfEq[T any](eq types.Equal[T], values ...T) types.LinkedList[T] {
	return newLinkedListOfEq[T](eq, values...).withRole(list)
}
