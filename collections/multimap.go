package collections

import (
	"github.com/cuzfrog/tgods/types"
)

type baseMultiMap[K any, V any] struct {
	types.Map[K, types.Collection[V]]
	size   int
	newCol func() types.Collection[V]
}

func newHashSetMultiMap[K any, V any](khs types.Hash[K], keq types.Equal[K], vhs types.Hash[V], veq types.Equal[V]) *baseMultiMap[K, V] {
	m := NewHashMapOf[K, types.Collection[V]](khs, keq)
	newCol := func() types.Collection[V] { return newHashTableOfInitCap(defaultArrInitSize, vhs, veq) }
	return &baseMultiMap[K, V]{m, 0, newCol}
}

func newArrayListMultiMap[K any, V any](khs types.Hash[K], keq types.Equal[K]) *baseMultiMap[K, V] {
	m := NewHashMapOf[K, types.Collection[V]](khs, keq)
	newCol := func() types.Collection[V] { return newCircularArrayOfEq[V](defaultArrInitSize, nil) }
	return &baseMultiMap[K, V]{m, 0, newCol}
}

func (l *baseMultiMap[K, V]) PutSingle(k K, v V) {
	s, found := l.Get(k)
	if !found {
		s = l.newCol()
	}
	oldSize := s.Size()
	s.Add(v)
	newSize := s.Size()
	l.size += newSize - oldSize
	l.Map.Put(k, s)
}

func (l *baseMultiMap[K, V]) Put(k K, v types.Collection[V]) (types.Collection[V], bool) {
	old, found := l.Map.Put(k, v)
	if found {
		l.size += v.Size() - old.Size()
	} else {
		l.size += v.Size()
	}
	return old, found
}

func (l *baseMultiMap[K, V]) Add(entry types.Entry[K, types.Collection[V]]) bool {
	l.Put(entry.Key(), entry.Value())
	return true
}

func (l *baseMultiMap[K, V]) Size() int {
	return l.size
}

func (l *baseMultiMap[K, V]) KeySize() int {
	return l.Map.Size()
}

func (l *baseMultiMap[K, V]) Contains(entry types.Entry[K, types.Collection[V]]) bool {
	return l.Map.ContainsKey(entry.Key())
}
