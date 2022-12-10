package collections

import (
	"github.com/cuzfrog/tgods/types"
)

// NewArrayStack creates an array stack with fixed cap, allocating the underlying array eagerly
func NewArrayStack[T comparable](size int) types.Stack[T] {
	return newArrayStack[T](size)
}

// NewArrayStackC creates an array stack with fixed cap, allocating the underlying array eagerly
// 'C' stands for Client Customized Constrained type.
func NewArrayStackC[T types.WithEqual[T]](size int) types.Stack[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newArrayStackOfEq[T](size, eq)
}

// NewLinkedListStack creates a stack from a linkedList
func NewLinkedListStack[T comparable]() types.Stack[T] {
	return newLinkedListOf[T]().withRole(stack)
}

// NewLinkedListStackC creates a stack from a linkedList
// 'C' stands for Client Customized Constrained type.
func NewLinkedListStackC[T types.WithEqual[T]]() types.Stack[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newLinkedListOfEq[T](eq).withRole(stack)
}

// NewLinkedListStackOfEq creates a stack from a linkedList with custom Equal
func NewLinkedListStackOfEq[T any](eq types.Equal[T]) types.Stack[T] {
	return newLinkedListOfEq[T](eq).withRole(stack)
}

// NewCircularArrayStack creates a stack from an auto-recap circularArray
func NewCircularArrayStack[T comparable]() types.Stack[T] {
	return newCircularArrayOf[T]().withRole(stack)
}

// NewCircularArrayStackC creates a stack from an auto expansion/shrinking circularArray
// 'C' stands for Client Customized Constrained type.
func NewCircularArrayStackC[T types.WithEqual[T]]() types.Stack[T] {
	eq := func(a, b T) bool { return a.Equal(b) }
	return newCircularArrayOfEq[T](0, eq).withRole(stack)
}

// NewCircularArrayStackOfEq creates a stack from an auto-recap circularArray with initCap and custom Equal
func NewCircularArrayStackOfEq[T any](initCap int, eq types.Equal[T]) types.Stack[T] {
	return newCircularArrayOfEq[T](initCap, eq).withRole(stack)
}
