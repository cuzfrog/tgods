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
	s := newLinkedListOf[T]()
	return s
}

// NewCircularArrayStack create a stack with circularArray as the implementation
func NewCircularArrayStack[T comparable]() types.Stack[T] {
	s := newCircularArrayOf[T]()
	s.cl = stack
	return s
}
