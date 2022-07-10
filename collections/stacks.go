package collections

import (
	"github.com/cuzfrog/tgods/types"
)

// NewArrayStack creates a stack with fixed cap, allocating the underlying array eagerly
func NewArrayStack[T comparable](size int) types.Stack[T] {
	return newArrayStack[T](size)
}

// NewLinkedStack create a stack with linkedList as the implementation
func NewLinkedStack[T comparable]() types.Stack[T] {
	return newLinkedListOf[T]().withRole(stack)
}

// NewCircularArrayStack create a stack with circularArray as the implementation
func NewCircularArrayStack[T comparable]() types.Stack[T] {
	return newCircularArrayOf[T]().withRole(stack)
}
