package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"golang.org/x/exp/constraints"
)

// max length is int despite the type constraint
type enumMap[K constraints.Integer, V any] struct {
	arr  []types.Entry[K, V]
	size int
}

func newEnumMap[K constraints.Integer, V any](max K, entries ...types.Entry[K, V]) *enumMap[K, V] {
	arr := make([]types.Entry[K, V], max+1)
	m := &enumMap[K, V]{arr, 0}
	for _, e := range entries {
		if arr[e.Key()] == nil {
			m.size++
		}
		arr[e.Key()] = EntryOf(e.Key(), e.Value())
	}
	return m
}

func (m *enumMap[K, V]) Add(entry types.Entry[K, V]) bool {
	m.Put(entry.Key(), entry.Value())
	return true
}

func (m *enumMap[K, V]) Contains(_ types.Entry[K, V]) bool {
	panic("Not supported, please check with ContainsKey.")
}

func (m *enumMap[K, V]) ContainsKey(k K) bool {
	return m.arr[k] != nil
}

func (m *enumMap[K, V]) Size() int {
	return m.size
}

func (m *enumMap[K, V]) Clear() {
	m.size = 0
	m.arr = make([]types.Entry[K, V], cap(m.arr))
}

func (m *enumMap[K, V]) Get(k K) (V, bool) {
	v := m.arr[k]
	if v != nil {
		return v.Value(), true
	}
	return utils.Nil[V](), false
}

func (m *enumMap[K, V]) Put(k K, v V) (old V, found bool) {
	e := m.arr[k]
	if e == nil {
		m.size++
	} else {
		old = e.Value()
		found = true
	}
	m.arr[k] = EntryOf(k, v)
	return
}

func (m *enumMap[K, V]) Remove(k K) (old V, found bool) {
	e := m.arr[k]
	if e != nil {
		m.size--
		old = e.Value()
		found = true
		m.arr[k] = nil
	}
	return
}

func (m *enumMap[K, V]) First() types.Entry[K, V] {
	for _, v := range m.arr {
		if v != nil {
			return v
		}
	}
	return nil
}

func (m *enumMap[K, V]) Last() types.Entry[K, V] {
	for i := len(m.arr) - 1; i >= 0; i-- {
		v := m.arr[i]
		if v != nil {
			return v
		}
	}
	return nil
}

func (m *enumMap[K, V]) RemoveFirst() types.Entry[K, V] {
	for i := 0; i < len(m.arr); i++ {
		v := m.arr[i]
		if v != nil {
			m.arr[i] = nil
			m.size--
			return v
		}
	}
	return nil
}

func (m *enumMap[K, V]) RemoveLast() types.Entry[K, V] {
	for i := len(m.arr) - 1; i >= 0; i-- {
		v := m.arr[i]
		if v != nil {
			m.arr[i] = nil
			m.size--
			return v
		}
	}
	return nil
}
