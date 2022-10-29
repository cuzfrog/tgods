package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type arrayListMultiMap[K any, V any] struct {
	types.Map[K, types.List[V]]
	size    int
	dummyEq types.Equal[V]
}

func newArrayListMultiMap[K any, V any](hs types.Hash[K], eq types.Equal[K]) *arrayListMultiMap[K, V] {
	m := NewHashMapOf[K, types.List[V]](hs, eq)
	return &arrayListMultiMap[K, V]{m, 0, nil}
}

func (l *arrayListMultiMap[K, V]) PutSingle(k K, v V) {
	l.size++
	list, found := l.Get(k)
	if !found {
		list = NewArrayListOfEq(defaultArrInitSize, l.dummyEq)
	}
	list.Add(v)
	l.Map.Put(k, list)
}

func (l *arrayListMultiMap[K, V]) Put(k K, v types.List[V]) (types.List[V], bool) {
	old, found := l.Map.Put(k, v)
	if found {
		l.size += v.Size() - old.Size()
	} else {
		l.size += v.Size()
	}
	return old, found
}

func (l *arrayListMultiMap[K, V]) Add(entry types.Entry[K, types.List[V]]) bool {
	l.Put(entry.Key(), entry.Value())
	return true
}

func (l *arrayListMultiMap[K, V]) AllValues() types.List[V] {
	list := NewArrayListOfEq(l.Size(), l.dummyEq)
	itor := l.Map.Iterator()
	for itor.Next() {
		entry := itor.Value()
		utils.AddAll[V](entry.Value(), list)
	}
	return list
}

func (l *arrayListMultiMap[K, V]) Size() int {
	return l.size
}

func (l *arrayListMultiMap[K, V]) KeySize() int {
	return l.Map.Size()
}

func (l *arrayListMultiMap[K, V]) Contains(entry types.Entry[K, types.List[V]]) bool {
	return l.Map.ContainsKey(entry.Key())
}
