//go:build test

package mocks

import (
	"github.com/cuzfrog/tgods/types"
)

type MockCollection[T comparable] interface {
	types.Collection[T]
	SetElems(values ...T)
}

type MockList[T comparable] interface {
	types.List[T]
	SetElems(values ...T)
}

type mockCollection[T comparable] struct {
	arr  []T
	size int
}

// ======== Constructors ========

func NewMockCollectionOf[T comparable](values ...T) MockCollection[T] {
	return &mockCollection[T]{values, len(values)}
}

func NewMockCollection[T comparable](size int) MockCollection[T] {
	arr := make([]T, size)
	return &mockCollection[T]{arr, 0}
}

// ======== Mocks ========

func (mc *mockCollection[T]) SetElems(values ...T) {
	mc.arr = values
	mc.size = len(values)
}

// ======== Implementations ========

func (mc *mockCollection[T]) Size() int {
	return mc.size
}

func (mc *mockCollection[T]) Contains(elem T) bool {
	for _, v := range mc.arr {
		if v == elem {
			return true
		}
	}
	return false
}

func (mc *mockCollection[T]) Each(fn func(index int, elem T)) {
	for i, t := range mc.arr {
		fn(i, t)
	}
}

func (mc *mockCollection[T]) Clear() {
	mc.arr = make([]T, 0)
	mc.size = 0
}

func (mc *mockCollection[T]) AddHead(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (mc *mockCollection[T]) RemoveHead() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (mc *mockCollection[T]) Head() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (mc *mockCollection[T]) AddTail(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (mc *mockCollection[T]) RemoveTail() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (mc *mockCollection[T]) Add(elem T) bool {
	if len(mc.arr) == 0 {
		mc.arr = make([]T, 32)
	}
	mc.arr[mc.size] = elem
	mc.size++
	return true
}

func (mc *mockCollection[T]) Remove() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (mc *mockCollection[T]) Tail() (T, bool) {
	//TODO implement me
	panic("implement me")
}

// ======== Iterator ========

func (mc *mockCollection[T]) Iterator() types.Iterator[T] {
	return &mockIterator[T]{mc.arr, -1}
}

type mockIterator[T any] struct {
	arr []T
	cur int
}

func (it *mockIterator[T]) Next() bool {
	it.cur++
	return it.cur < len(it.arr)
}

func (it *mockIterator[T]) Index() int {
	return it.cur
}

func (it *mockIterator[T]) Value() T {
	return it.arr[it.cur]
}
