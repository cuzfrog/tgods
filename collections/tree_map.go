package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type treeMap[K any, V any] struct {
	*rbTree[types.Entry[K, V]]
	t *rbTree[types.Entry[K, V]] // point to the same tree, as a type workaround
}

func newTreeMapOfComp[K any, V any](comp types.Compare[K]) *treeMap[K, V] {
	entryComp := func(a, b types.Entry[K, V]) int8 { return comp(a.Key(), b.Key()) }
	t := &rbTree[types.Entry[K, V]]{nil, 0, entryComp}
	return &treeMap[K, V]{t, t}
}

func (m *treeMap[K, V]) Get(k K) (V, bool) {
	n := searchNode[types.Entry[K, V]](m.root, keyEntry[K, V]{k}, m.comp)
	if n == nil {
		return utils.Nil[V](), false
	}
	return n.v.Value(), true
}

func (m *treeMap[K, V]) Put(k K, v V) (V, bool) {
	e := EntryOf(k, v)
	old, found := rbTreeInsert[types.Entry[K, V]](m.t, e)
	if found {
		return old.Value(), true
	}
	return utils.Nil[V](), false
}

func (m *treeMap[K, V]) Remove(k K) (V, bool) {
	nd, found := m.t.delete(keyEntry[K, V]{k})
	if found {
		return nd.Value(), true
	}
	return utils.Nil[V](), false
}

func (m *treeMap[K, V]) ContainsKey(k K) bool {
	n := searchNode[types.Entry[K, V]](m.root, keyEntry[K, V]{k}, m.comp)
	return n != nil
}

func (m *treeMap[K, V]) Contains(_ types.Entry[K, V]) bool {
	panic("Not supported. Please use ContainsKey()")
}
