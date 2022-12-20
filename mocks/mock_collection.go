//go:build test

package mocks

import (
	"github.com/cuzfrog/tgods/types"
)

type MockCollection[T comparable] interface {
	types.Collection[T]
	types.IndexAccess[T]
	SetElems(values ...T)
	Elems() []T
	GetFlag(key string) interface{}
}

type MockList[T comparable] interface {
	types.List[T]
	SetElems(values ...T)
}

type mockCollection[T comparable] struct {
	arr   []T
	size  int
	flags map[string]interface{}
}

// ======== Constructors ========

func NewMockCollectionOf[T comparable](values ...T) MockCollection[T] {
	return &mockCollection[T]{values, len(values), make(map[string]interface{}, 10)}
}

func NewMockCollection[T comparable](size int) MockCollection[T] {
	arr := make([]T, size)
	return &mockCollection[T]{arr, 0, make(map[string]interface{}, 10)}
}

// ======== Mocks ========

func (mc *mockCollection[T]) SetElems(values ...T) {
	mc.arr = values
	mc.size = len(values)
}

func (mc *mockCollection[T]) Elems() []T {
	return mc.arr
}

func (mc *mockCollection[T]) GetFlag(key string) interface{} {
	return mc.flags[key]
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

func (mc *mockCollection[T]) Get(index int) (T, bool) {
	panic("implement me")
}
func (mc mockCollection[T]) MustGet(index int) T {
	panic("implement me")
}
func (mc *mockCollection[T]) Set(index int, elem T) (T, bool) {
	panic("implement me")
}
func (mc *mockCollection[T]) Swap(indexA, indexB int) bool {
	panic("implement me")
}
func (mc *mockCollection[T]) Sort(lessFn types.Less[T]) {
	mc.flags["SortLessFn"] = lessFn
}

// ======== Iterator ========

func (mc *mockCollection[T]) Iterator() types.Iterator[T] {
	return &mockIterator[T]{mc.arr, -1, mc.size}
}

type mockIterator[T any] struct {
	arr  []T
	cur  int
	size int
}

func (it *mockIterator[T]) Next() bool {
	it.cur++
	return it.cur < it.size
}

func (it *mockIterator[T]) Index() int {
	return it.cur
}

func (it *mockIterator[T]) Value() T {
	return it.arr[it.cur]
}
