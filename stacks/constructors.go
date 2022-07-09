package stacks

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/lists"
)

func newArrayStack[T comparable](size int) *arrayStack[T] {
	arr := make([]T, size)
	return &arrayStack[T]{arr, -1}
}

// NewArrayStack creates a stack with fixed cap, allocating the underlying array eagerly
func NewArrayStack[T comparable](size int) core.Stack[T] {
	return newArrayStack[T](size)
}

// NewLinkedStack create a stack with linkedList as the implementation
func NewLinkedStack[T comparable]() core.Stack[T] {
	return lists.NewLinkedList[T]()
}

// NewCircularArrayStack create a stack with circularArrayList as the implementation
func NewCircularArrayStack[T comparable]() core.Stack[T] {
	return lists.NewCircularArrayListOf[T]()
}
