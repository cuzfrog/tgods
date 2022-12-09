package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

func NewHeapPriorityQueue[T any](comp types.Compare[T]) types.Queue[T] {
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

func NewLinkedListQueueOfEq[T any](eq types.Equal[T]) types.Queue[T] {
	return newLinkedListOfEq[T](eq).withRole(queue)
}

func NewArrayListQueue[T comparable]() types.Queue[T] {
	return newCircularArrayOf[T]().withRole(queue)
}

func NewArrayListQueueOfSize[T comparable](initCap int) types.Queue[T] {
	return newCircularArray[T](initCap, AutoExpand+AutoShrink).withRole(queue)
}

func NewArrayListQueueOfSizeP[T comparable](initCap int, flag AutoSizingFlag) types.Queue[T] {
	return newCircularArray[T](initCap, flag).withRole(queue)
}

func NewArrayListQueueOfEq[T any](initCap int, eq types.Equal[T]) types.Queue[T] {
	return newCircularArrayOfEq[T](initCap, eq).withRole(queue)
}

func NewArrayListQueueOfEqP[T any](initCap int, eq types.Equal[T], flag AutoSizingFlag) types.Queue[T] {
	return newCircularArrayOfEqP[T](initCap, eq, flag).withRole(queue)
}
