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
