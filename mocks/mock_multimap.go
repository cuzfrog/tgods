package mocks

import "github.com/cuzfrog/tgods/types"

type mockMultiMap[K comparable, V comparable] struct {
	types.Map[K, types.Collection[V]]
}

func NewMockMultiMap[K comparable, V comparable]() types.MultiMap[K, V] {
	m := newMockMap[K, types.Collection[V]]()
	return &mockMultiMap[K, V]{m}
}

func (l *mockMultiMap[K, V]) Add(entry types.Entry[K, types.Collection[V]]) bool {
	l.Put(entry.Key(), entry.Value())
	return true
}

func (m *mockMultiMap[K, V]) Size() int {
	it := m.Map.Iterator()
	s := 0
	for it.Next() {
		s += it.Value().Value().Size()
	}
	return s
}

func (m *mockMultiMap[K, V]) PutSingle(k K, v V) {
	vs, ok := m.Get(k)
	if !ok {
		vs = NewMockCollection[V](32)
	}
	vs.Add(v)
	m.Put(k, vs)
}

func (m *mockMultiMap[K, V]) KeySize() int {
	return m.Map.Size()
}

func (l *mockMultiMap[K, V]) Contains(entry types.Entry[K, types.Collection[V]]) bool {
	return l.Map.ContainsKey(entry.Key())
}

func (m *mockMultiMap[K, V]) Iterator() types.Iterator[types.Entry[K, types.Collection[V]]] {
	return m.Map.Iterator()
}
