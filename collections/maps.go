package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

func NewTreeMapOf[K constraints.Ordered, V any](entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	return NewTreeMapOfComp(funcs.ValueCompare[K], entries...)
}

func NewTreeMapOfComp[K any, V any](comp types.Compare[K], entries ...types.Entry[K, V]) types.SortedMap[K, V] {
	m := newTreeMapOfComp[K, V](comp)
	for _, e := range entries {
		m.Put(e.Key(), e.Value())
	}
	return m
}
