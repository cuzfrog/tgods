package funcs

import "golang.org/x/exp/constraints"

// Compare is a func that return 0 when a == b, 1 when a > b, -1 when a < b
type Compare[T any] func(a, b T) int8

type Equal[T any] func(a, b T) bool

// ValueCompare bounded by constraints.Ordered
func ValueCompare[T constraints.Ordered](a, b T) int8 {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

// ValueEqual bounded by native comparable
func ValueEqual[T comparable](a, b T) bool {
	return a == b
}

// ========== HOF =========

// CompToEq is a high-order function that converts a Compare to an Equal
func CompToEq[T any](comp Compare[T]) Equal[T] {
	return func(a, b T) bool { return comp(a, b) == 0 }
}

// InverseComp changes a greater than Compare to a smaller than, and vice versa
func InverseComp[T any](comp Compare[T]) Compare[T] {
	return func(a, b T) int8 { return comp(b, a) }
}
