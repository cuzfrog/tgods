package utils

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

type mockCollection[T any] struct {
	arr  []T
	size int
	eq   funcs.Equal[T]
}

func newMockCollectionOf[T any](value ...T) *mockCollection[T] {
	return &mockCollection[T]{value, len(value), nil}
}

func newMockCollection[T any](eq funcs.Equal[T], value ...T) *mockCollection[T] {
	return &mockCollection[T]{value, len(value), eq}
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
