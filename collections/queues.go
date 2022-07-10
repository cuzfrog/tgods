package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

func NewHeapPriorityQueue[T any](comp funcs.Compare[T]) types.Queue[T] {
	return &binaryHeap[T]{newCircularArrayOfEq[T](0, funcs.CompToEq(comp)), comp}
}

func NewHeapMaxPriorityQueue[T constraints.Ordered]() types.Queue[T] {
	return &binaryHeap[T]{newCircularArrayOfEq[T](0, funcs.CompToEq(funcs.ValueCompare[T])), funcs.ValueCompare[T]}
}
func NewHeapMinPriorityQueue[T constraints.Ordered]() types.Queue[T] {
	return &binaryHeap[T]{newCircularArrayOfEq[T](0, funcs.CompToEq(funcs.ValueCompare[T])), funcs.InverseComp(funcs.ValueCompare[T])}
}
