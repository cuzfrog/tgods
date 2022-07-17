package collections

import "github.com/cuzfrog/tgods/types"

func NewTreeSetOf[T comparable](values ...T) types.SortedSet[T] {
	return newTreeSetOf(values...)
}
