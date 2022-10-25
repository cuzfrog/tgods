package collections

import "github.com/cuzfrog/tgods/types"

type arrayListMultiMap[K any, V any] struct {
	types.Map[K, types.List[V]]
	size int
	eq   types.Equal[V]
}

func newArrayListMultiMap[K any, V any](hs types.Hash[K], eq types.Equal[K]) *arrayListMultiMap[K, V] {
	m := NewHashMapOf[K, types.List[V]](hs, eq)
	return &arrayListMultiMap[K, V]{m, 0, nil}
}

func (l *arrayListMultiMap[K, V]) PutSingle(k K, v V) {
	l.size++
	list, found := l.Get(k)
	if !found {
		list = NewArrayListOfEq(defaultArrInitSize, l.eq)
	}
	list.Add(v)
}

func (l *arrayListMultiMap[K, V]) Put(k K, v types.List[V]) (types.List[V], bool) {
	panic("not implemented")
}

func (l *arrayListMultiMap[K, V]) Add(entry types.Entry[K, types.List[V]]) bool {
	panic("not implemented")
}

func (l *arrayListMultiMap[K, V]) AllValues() types.List[V] {
	list := NewArrayListOfEq(l.Size(), l.eq)
	itor := l.Map.Iterator()
	for itor.Next() {
		entry := itor.Value()
		subItor := entry.Value().Iterator() // TODO: replace with AddAll()
		for subItor.Next() {
			list.Add(subItor.Value())
		}
	}
	return list
}

func (l *arrayListMultiMap[K, V]) Size() int {
	return l.size
}

func (l *arrayListMultiMap[K, V]) KeySize() int {
	return l.Map.Size()
}
