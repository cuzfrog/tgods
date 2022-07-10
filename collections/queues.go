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
	l := newLinkedListOf[T]()
	return l
}

func NewLinkedListQueueOfEq[T any](eq funcs.Equal[T]) types.Queue[T] {
	return newLinkedListOfEq[T](eq)
}

func NewArrayListQueue[T comparable]() types.Queue[T] {
	l := newCircularArrayOf[T]()
	l.cl = queue
	return l
}

func NewArrayListQueueOfSize[T comparable](initSize int) types.Queue[T] {
	l := newCircularArray[T](initSize)
	l.cl = queue
	return l
}

func NewArrayListQueueOfEq[T any](initSize int, eq funcs.Equal[T]) types.Queue[T] {
	l := newCircularArrayOfEq[T](initSize, eq)
	l.cl = queue
	return l
}
