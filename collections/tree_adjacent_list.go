package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type treeAdjacencyList[V any, E any] struct {
	treeMap[V, types.Map[V, E]]
	inward treeMap[V, types.Map[V, E]] // keeps inward record for fast deletion and inward iteration
	comp   types.Compare[V]
	newMap func() types.Map[V, E]
}

func newTreeAdjacencyList[V any, E any](comp types.Compare[V]) *treeAdjacencyList[V, E] {
	newMap := func() types.Map[V, E] { return newTreeMapOfComp[V, E](comp) }
	return &treeAdjacencyList[V, E]{*newTreeMapOfComp[V, types.Map[V, E]](comp), *newTreeMapOfComp[V, types.Map[V, E]](comp), comp, newMap}
}

func (t *treeAdjacencyList[V, E]) Add(vertex V) bool {
	utils.ComputeIfAbsent[V, types.Map[V, E]](&t.treeMap, vertex, t.newMap)
	return true
}

func (t *treeAdjacencyList[V, E]) Contains(vertex V) bool {
	return t.ContainsKey(vertex)
}

func (t *treeAdjacencyList[V, E]) Connect(from, to V, edge E) (E, bool) {
	outwards, _ := utils.ComputeIfAbsent[V, types.Map[V, E]](&t.treeMap, from, t.newMap)
	t.Add(to)
	if t.comp(from, to) == 0 {
		return utils.Nil[E](), false
	}
	inwards, _ := utils.ComputeIfAbsent[V, types.Map[V, E]](&t.inward, to, t.newMap)
	inwards.Put(from, edge)
	return outwards.Put(to, edge)
}

func (t *treeAdjacencyList[V, E]) Disconnect(from, to V) (E, bool) {
	if t.comp(from, to) == 0 {
		return utils.Nil[E](), false
	}
	outwards, found := t.treeMap.Get(from)
	if !found {
		return utils.Nil[E](), false
	}
	oldE, hasTo := outwards.Remove(to)
	if !hasTo {
		return utils.Nil[E](), false
	}
	inwards, _ := t.inward.Get(to)
	inwards.Remove(from)
	return oldE, true
}

func (t *treeAdjacencyList[V, E]) InwardCount(vertex V) int {
	inwards, found := t.inward.Get(vertex)
	if !found {
		return 0
	}
	return inwards.Size()
}

func (t *treeAdjacencyList[V, E]) OutwardCount(vertex V) int {
	outwards, found := t.treeMap.Get(vertex)
	if !found {
		return 0
	}
	return outwards.Size()
}

func (t *treeAdjacencyList[V, E]) InwardEdges(vertex V) (types.Iterator[types.Entry[V, E]], int) {
	inwards, found := t.inward.Get(vertex)
	if !found {
		return newEmptyIterator[types.Entry[V, E]](), 0
	}
	return inwards.Iterator(), inwards.Size()
}

func (t *treeAdjacencyList[V, E]) OutwardEdges(vertex V) (types.Iterator[types.Entry[V, E]], int) {
	outwards, found := t.treeMap.Get(vertex)
	if !found {
		return newEmptyIterator[types.Entry[V, E]](), 0
	}
	return outwards.Iterator(), outwards.Size()
}

func (t *treeAdjacencyList[V, E]) Remove(vertex V) bool {
	outwards, exist := t.treeMap.Remove(vertex)
	if !exist {
		return false
	}
	oit := outwards.Iterator()
	for oit.Next() {
		otherInwards, _ := t.inward.Get(oit.Value().Key())
		otherInwards.Remove(vertex)
	}
	inwards, hasInwards := t.inward.Remove(vertex)
	if hasInwards && inwards.Size() > 0 {
		it := inwards.Iterator()
		for it.Next() {
			otherOutwards, _ := t.treeMap.Get(it.Value().Key())
			otherOutwards.Remove(vertex)
		}
	}
	return true
}
