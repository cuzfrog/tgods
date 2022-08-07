package collections

import (
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

type enumSet[T constraints.Integer] struct {
	enumMap[T, interface{}]
}

func newEnumSet[T constraints.Integer](max T, values ...T) *enumSet[T] {
	arr := make([]types.Entry[T, interface{}], len(values))
	for i, v := range values {
		arr[i] = keyEntry[T, interface{}]{v}
	}
	return &enumSet[T]{*newEnumMap[T, interface{}](max, arr...)}
}
