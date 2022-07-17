package collections

import (
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

// NewTreeSetOf creates a red black tree backed SortedSet with init values
func NewTreeSetOf[T constraints.Ordered](values ...T) types.SortedSet[T] {
	return newRbTreeOf(values...)
}
