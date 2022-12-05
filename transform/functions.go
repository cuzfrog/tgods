package transform

import "github.com/cuzfrog/tgods/types"

// MapTo transforms elem into another Collection, returns elem count collected
func MapTo[T any, R any](src types.Collection[T], tgt types.Collection[R], mapFn func(elem T) R) int {
	cnt := 0
	it := src.Iterator()
	for it.Next() {
		tgt.Add(mapFn(it.Value()))
		cnt++
	}
	return cnt
}

// FilterMapTo filters and transforms elem into another Collection, returns elem count collected
func FilterMapTo[T any, R any](src types.Collection[T], tgt types.Collection[R], filterFn func(elem T) bool, mapFn func(elem T) R) int {
	cnt := 0
	it := src.Iterator()
	for it.Next() {
		if filterFn(it.Value()) {
			tgt.Add(mapFn(it.Value()))
			cnt++
		}
	}
	return cnt
}

// FilterTo filters elem into another Collection, returns elem count collected
func FilterTo[T any](src types.Collection[T], tgt types.Collection[T], filterFn func(elem T) bool) int {
	cnt := 0
	it := src.Iterator()
	for it.Next() {
		if filterFn(it.Value()) {
			tgt.Add(it.Value())
			cnt++
		}
	}
	return cnt
}

// FlatMapTo transforms elem into another Collection, returns transformed elem count collected
func FlatMapTo[T any, R any](src types.Collection[T], tgt types.Collection[R], mapFn func(elem T) []R) int {
	cnt := 0
	it := src.Iterator()
	for it.Next() {
		for _, v := range mapFn(it.Value()) {
			tgt.Add(v)
			cnt++
		}
	}
	return cnt
}

// FilterFlatMapTo filters and transforms elem into another Collection, returns transformed elem count collected
func FilterFlatMapTo[T any, R any](src types.Collection[T], tgt types.Collection[R], filterFn func(elem T) bool, mapFn func(elem T) []R) int {
	cnt := 0
	it := src.Iterator()
	for it.Next() {
		if filterFn(it.Value()) {
			for _, v := range mapFn(it.Value()) {
				tgt.Add(v)
				cnt++
			}
		}
	}
	return cnt
}

func Reduce[T any, R any](col types.Collection[T], identity R, reduceFn func(acc R, next T) R) R {
	it := col.Iterator()
	acc := identity
	for it.Next() {
		acc = reduceFn(acc, it.Value())
	}
	return acc
}
