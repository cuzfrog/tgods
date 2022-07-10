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

func NewLinkedListQueue[T comparable]() types.Queue[T] {
	return newLinkedListOf[T]().withRole(queue)
}

func NewLinkedListQueueOfEq[T any](eq funcs.Equal[T]) types.Queue[T] {
	return newLinkedListOfEq[T](eq).withRole(queue)
}

func NewArrayListQueue[T comparable]() types.Queue[T] {
	return newCircularArrayOf[T]().withRole(queue)
}

func NewArrayListQueueOfSize[T comparable](initSize int) types.Queue[T] {
	return newCircularArray[T](initSize).withRole(queue)
}

func NewArrayListQueueOfEq[T any](initSize int, eq funcs.Equal[T]) types.Queue[T] {
	return newCircularArrayOfEq[T](initSize, eq).withRole(queue)
}
