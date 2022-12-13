package funcs

import (
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
	"hash/fnv"
	"strconv"
)

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

func ValueLess[T constraints.Ordered](a, b T) bool {
	return a < b
}

func ValueGreater[T constraints.Ordered](a, b T) bool {
	return a > b
}

// ValueEqual bounded by native comparable
func ValueEqual[T comparable](a, b T) bool {
	return a == b
}

// NumHash simply returns the uint value, that means on a 32bit platform for a 64bit value, there would be considerable hash collisions in pattern.
func NumHash[T constraints.Integer | constraints.Float](n T) uint {
	return uint(n)
}

// NewStrHash returns a function closure of go native fnv.New32a or fnv.New64a depending on runtime platform
func NewStrHash() types.Hash[string] {
	if strconv.IntSize == 32 {
		fn := fnv.New32a()
		return func(s string) uint {
			_, _ = fn.Write([]byte(s))
			h := fn.Sum32()
			fn.Reset()
			return uint(h)
		}
	} else {
		fn := fnv.New64a()
		return func(s string) uint {
			_, _ = fn.Write([]byte(s))
			h := fn.Sum64()
			fn.Reset()
			return uint(h)
		}
	}
}

// ========== HOF =========

// CompToEq is a high-order function that converts a types.Compare to an types.Equal
func CompToEq[T any](comp types.Compare[T]) types.Equal[T] {
	return func(a, b T) bool { return comp(a, b) == 0 }
}

// CompToLess is a high-order function that converts a types.Compare to a types.Less
func CompToLess[T any](comp types.Compare[T]) types.Less[T] {
	return func(a, b T) bool { return comp(a, b) < 0 }
}

// InverseComp is a high-order function that changes a greater than types.Compare to a smaller than, and vice versa
func InverseComp[T any](comp types.Compare[T]) types.Compare[T] {
	return func(a, b T) int8 { return comp(b, a) }
}

// InverseLess is a high-order function that inverse a types.Less function
func InverseLess[T any](less types.Less[T]) types.Less[T] {
	return func(a, b T) bool { return less(b, a) }
}
