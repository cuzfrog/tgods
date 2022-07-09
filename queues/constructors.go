package queues

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/lists"
	"golang.org/x/exp/constraints"
)

func NewHeapPriorityQueue[T comparable](comparator core.Compare[T]) *binaryHeap[T] {
	return &binaryHeap[T]{lists.NewCircularArrayListOf[T](), comparator}
}

func NewHeapPriorityQueueForMaxValue[T constraints.Ordered]() *binaryHeap[T] {
	return &binaryHeap[T]{lists.NewCircularArrayListOf[T](), core.CompareOrdered[T]}
}
func NewHeapPriorityQueueForMinValue[T constraints.Ordered]() *binaryHeap[T] {
	fn := func(a, b T) int8 { return core.CompareOrdered(b, a) }
	return &binaryHeap[T]{lists.NewCircularArrayListOf[T](), fn}
}
