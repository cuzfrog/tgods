package transform

import (
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

// MapSliceTo transforms elem into another Collection, returns elem count collected
func MapSliceTo[T any, R any](src []T, tgt types.Collection[R], mapFn func(elem T) R) int {
	cnt := 0
	for _, t := range src {
		tgt.Add(mapFn(t))
		cnt++
	}
	return cnt
}

// FilterMapSliceTo filters and transforms elem into another Collection, returns elem count collected
func FilterMapSliceTo[T any, R any](src []T, tgt types.Collection[R], filterFn func(elem T) bool, mapFn func(elem T) R) int {
	cnt := 0
	for _, t := range src {
		if filterFn(t) {
			tgt.Add(mapFn(t))
			cnt++
		}
	}
	return cnt
}

// FilterSliceTo filters elem into another Collection, returns elem count collected
func FilterSliceTo[T any](src []T, tgt types.Collection[T], filterFn func(elem T) bool) int {
	cnt := 0
	for _, t := range src {
		if filterFn(t) {
			tgt.Add(t)
			cnt++
		}
	}
	return cnt
}

// FlatMapSliceTo transforms elem into another Collection, returns transformed elem count collected
func FlatMapSliceTo[T any, R any](src []T, tgt types.Collection[R], mapFn func(elem T) []R) int {
	cnt := 0
	for _, t := range src {
		for _, s := range mapFn(t) {
			tgt.Add(s)
			cnt++
		}
	}
	return cnt
}

// FilterFlatMapSliceTo filters and transforms elem into another Collection, returns transformed elem count collected
func FilterFlatMapSliceTo[T any, R any](src []T, tgt types.Collection[R], filterFn func(elem T) bool, mapFn func(elem T) []R) int {
	cnt := 0
	for _, t := range src {
		if filterFn(t) {
			for _, r := range mapFn(t) {
				tgt.Add(r)
				cnt++
			}
		}
	}
	return cnt
}

func FlattenSliceTo[T any](src [][]T, tgt types.Collection[T]) int {
	cnt := 0
	for _, ts := range src {
		for _, t := range ts {
			tgt.Add(t)
			cnt++
		}
	}
	return cnt
}

func ReduceSlice[T any, R any](arr []T, identity R, reduceFn func(acc R, next T) R) R {
	acc := identity
	for _, t := range arr {
		acc = reduceFn(acc, t)
	}
	return acc
}

func CountSlice[T any](arr []T, conditionFunc func(elem T) bool) int {
	cnt := 0
	for _, t := range arr {
		if conditionFunc(t) {
			cnt++
		}
	}
	return cnt
}

func SumSlice[T constraints.Ordered](arr []T) T {
	var s T
	for _, t := range arr {
		s += t
	}
	return s
}
