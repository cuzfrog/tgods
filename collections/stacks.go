package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

// NewArrayStack creates an array stack with fixed cap, allocating the underlying array eagerly
func NewArrayStack[T comparable](size int) types.Stack[T] {
	return newArrayStack[T](size)
}

// NewLinkedListStack creates a stack from a linkedList
func NewLinkedListStack[T comparable]() types.Stack[T] {
	return newLinkedListOf[T]().withRole(stack)
}

// NewLinkedListStackOfEq creates a stack from a linkedList with custom Equal
func NewLinkedListStackOfEq[T any](eq funcs.Equal[T]) types.Stack[T] {
	return newLinkedListOfEq[T](eq).withRole(stack)
}

// NewCircularArrayStack creates a stack from an auto-recap circularArray
func NewCircularArrayStack[T comparable]() types.Stack[T] {
	return newCircularArrayOf[T]().withRole(stack)
}

// NewCircularArrayStackOfEq creates a stack from an auto-recap circularArray with initSize and custom Equal
func NewCircularArrayStackOfEq[T any](initSize int, eq funcs.Equal[T]) types.Stack[T] {
	return newCircularArrayOfEq[T](initSize, eq).withRole(stack)
}
