package collections

import "github.com/cuzfrog/tgods/types"

func NewLinkedListDeque[T comparable]() types.Deque[T] {
	return newLinkedListOf[T]().withRole(deque)
}

func NewLinkedListDequeOfEq[T any](eq types.Equal[T]) types.Deque[T] {
	return newLinkedListOfEq[T](eq).withRole(deque)
}

func NewArrayListDeque[T comparable]() types.Deque[T] {
	return newCircularArrayOf[T]().withRole(deque)
}

func NewArrayListDequeOfSize[T comparable](initSize int) types.Deque[T] {
	return newCircularArray[T](initSize).withRole(deque)
}

func NewArrayListDequeOfEq[T any](initSize int, eq types.Equal[T]) types.Deque[T] {
	return newCircularArrayOfEq[T](initSize, eq).withRole(deque)
}
