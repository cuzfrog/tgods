package utils

import "golang.org/x/exp/constraints"

// Comparator is a func that return 0 when a == b, 1 when a > b, -1 when a < b
type Comparator[T any] func(a, b T) int8

func ValueComparator[T constraints.Ordered](a, b T) int8 {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
