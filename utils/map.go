package utils

import "github.com/cuzfrog/tgods/types"

func KeysFrom[K any, V any](m types.Map[K, V]) []K {
	s := make([]K, m.Size())
	it := m.Iterator()
	for it.Next() {
		s[it.Index()] = it.Value().Key()
	}
	return s
}

func ValuesFrom[K any, V any](m types.Map[K, V]) []V {
	s := make([]V, m.Size())
	it := m.Iterator()
	for it.Next() {
		s[it.Index()] = it.Value().Value()
	}
	return s
}

//Compute sets the computed value for the given key by a re-mapping function. Return the computed value.
//  fn - the re-mapping function: v - the old value or Nil, found - if there's an association of the key
func Compute[K any, V any](m types.Map[K, V], k K, fn func(v V, found bool) V) V {
	oldV, found := m.Get(k)
	newV := fn(oldV, found)
	m.Put(k, newV)
	return newV
}
