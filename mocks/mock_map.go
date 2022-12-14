//go:build test

package mocks

import (
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/maps"
)

type mockEntry[K comparable, V any] struct {
	k K
	v V
}

func (e mockEntry[K, V]) Key() K {
	return e.k
}

func (e mockEntry[K, V]) Value() V {
	return e.v
}

type mockMap[K comparable, V any] struct {
	m map[K]V
}

func newMockMap[K comparable, V any]() *mockMap[K, V] {
	m := make(map[K]V, 32)
	return &mockMap[K, V]{m}
}

func NewMockMap[K comparable, V any]() types.Map[K, V] {
	return newMockMap[K, V]()
}

func (m *mockMap[K, V]) Add(entry types.Entry[K, V]) bool {
	m.m[entry.Key()] = entry.Value()
	return true
}

func (m *mockMap[K, V]) Contains(entry types.Entry[K, V]) bool {
	_, ok := m.m[entry.Key()]
	return ok
}

func (m *mockMap[K, V]) Size() int {
	return len(m.m)
}

func (m *mockMap[K, V]) Clear() {
	maps.Clear(m.m)
}

func (m *mockMap[K, V]) Get(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}

func (m *mockMap[K, V]) Put(k K, v V) (V, bool) {
	old, ok := m.m[k]
	m.m[k] = v
	return old, ok
}

func (m *mockMap[K, V]) Remove(k K) (V, bool) {
	old, ok := m.m[k]
	delete(m.m, k)
	return old, ok
}

func (m *mockMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m *mockMap[K, V]) Iterator() types.Iterator[types.Entry[K, V]] {
	size := len(m.m)
	arr := make([]types.Entry[K, V], size)
	i := 0
	for k, v := range m.m {
		arr[i] = mockEntry[K, V]{k, v}
		i++
	}

	return &mockIterator[types.Entry[K, V]]{arr, -1, size}
}

func (m *mockMap[K, V]) Each(f func(index int, elem types.Entry[K, V])) {
	i := 0
	for k, v := range m.m {
		f(i, mockEntry[K, V]{k, v})
		i++
	}
}
