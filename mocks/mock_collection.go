//go:build test

package mocks

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

type MockCollection[T comparable] interface {
	types.Collection[T]
	SetElems(values ...T)
}

type mockCollection[T comparable] struct {
	arr  []T
	size int
	eq   funcs.Equal[T]
}

func NewMockCollectionOf[T comparable](values ...T) MockCollection[T] {
	return &mockCollection[T]{values, len(values), funcs.ValueEqual[T]}
}

func NewMockCollection[T comparable](size int) MockCollection[T] {
	arr := make([]T, size)
	return &mockCollection[T]{arr, 0, funcs.ValueEqual[T]}
}

func (mc *mockCollection[T]) SetElems(values ...T) {
	mc.arr = values
	mc.size = len(values)
}

func (mc *mockCollection[T]) Size() int {
	return mc.size
}

func (mc *mockCollection[T]) Contains(elem T) bool {
	for _, v := range mc.arr {
		if mc.eq(v, elem) {
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

func (mc *mockCollection[T]) Iterator() types.Iterator[T] {
	return &mockIterator[T]{mc.arr, -1}
}

func (mc *mockCollection[T]) Clear() {
	mc.arr = make([]T, 0)
	mc.size = 0
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
