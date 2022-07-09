package core

import "golang.org/x/exp/constraints"

// Compare is a func that return 0 when a == b, 1 when a > b, -1 when a < b
type Compare[T any] func(a, b T) int8

type Equal[T any] func(a, b T) bool

func CompareOrdered[T constraints.Ordered](a, b T) int8 {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func EqualComparable[T comparable](a, b T) bool {
	return a == b
}
