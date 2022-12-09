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

func NewArrayListDequeOfSize[T comparable](initCap int) types.Deque[T] {
	return newCircularArray[T](initCap, AutoExpand+AutoShrink).withRole(deque)
}

func NewArrayListDequeOfSizeP[T comparable](initCap int, autoSizingFlag AutoSizingFlag) types.Deque[T] {
	return newCircularArray[T](initCap, autoSizingFlag).withRole(deque)
}

func NewArrayListDequeOfEq[T any](initCap int, eq types.Equal[T]) types.Deque[T] {
	return newCircularArrayOfEq[T](initCap, eq).withRole(deque)
}

func NewArrayListDequeOfEqP[T any](initCap int, eq types.Equal[T], autoSizingFlag AutoSizingFlag) types.Deque[T] {
	return newCircularArrayOfEqP[T](initCap, eq, autoSizingFlag).withRole(deque)
}
