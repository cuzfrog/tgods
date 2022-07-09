package queues

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/lists"
	"github.com/cuzfrog/tgods/utils"
	"golang.org/x/exp/constraints"
)

// assert binaryHeap implementation
var _ core.Queue[int] = (*binaryHeap[int])(nil)

type binaryHeap[T comparable] struct {
	arr  core.ArrayList[T]
	comp core.Compare[T]
}

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

func (h *binaryHeap[T]) Size() int {
	return h.arr.Size()
}

func (h *binaryHeap[T]) Clear() {
	h.arr.Clear()
}

func (h *binaryHeap[T]) Enqueue(elem T) bool {
	ret := h.arr.Add(elem)
	h.swim()
	return ret
}

func (h *binaryHeap[T]) Dequeue() (T, bool) {
	h.sink()
	return h.arr.Remove()
}

func (h *binaryHeap[T]) Peek() (T, bool) {
	return h.arr.Head()
}

// Contains delegates to underlying array, O(n)
func (h *binaryHeap[T]) Contains(elem T) bool {
	return h.arr.Contains(elem)
}

type binaryHeapIterator[T comparable] struct {
	q     core.Queue[T]
	index int
	v     T
}

func (it *binaryHeapIterator[T]) Next() bool {
	if it.q.Size() <= 0 {
		return false
	}
	it.index++
	it.v, _ = it.q.Dequeue()
	return true
}

// Index returns current index from iteration order, starting from 0.
// If Next return false, the number returned is meaningless, it must be guarded by Next
func (it *binaryHeapIterator[T]) Index() int {
	return it.index
}

// Value returns current value.
// If Next return false, the value returned is meaningless, it must be guarded by Next
func (it *binaryHeapIterator[T]) Value() T {
	return it.v
}

func (h *binaryHeap[T]) Iterator() core.Iterator[T] {
	return &binaryHeapIterator[T]{h.Clone(), -1, utils.Nil[T]()}
}

func (h *binaryHeap[T]) Clone() core.Queue[T] {
	return &binaryHeap[T]{h.arr.Clone(), h.comp}
}

// swim reheapifies by checking and moving up the last element of the arr, this should be called after adding
// i - child node index
// pi - parent node index
func (h *binaryHeap[T]) swim() {
	if h.arr.Size() <= 1 {
		return
	}
	i := h.arr.Size()
	for {
		pi := i / 2
		vi, _ := h.arr.Get(i - 1)
		vpi, _ := h.arr.Get(pi - 1)
		if h.comp(vi, vpi) <= 0 {
			break
		}
		h.arr.Swap(i-1, pi-1)
		if pi <= 1 {
			break
		}
		i = pi
	}
}

// sink reheapifies by checking and moving down the root elem of the arr, this should be called before popping
// pi - parent node index
// i1 - left child node index
// i2 - right child node index
func (h *binaryHeap[T]) sink() {
	if h.arr.Size() <= 1 {
		return
	}
	nextSize := h.arr.Size() - 1
	h.arr.Swap(0, nextSize) // move max to the end
	pi := 1
	for {
		i1 := pi * 2
		if i1 > nextSize {
			break
		}
		i2 := i1 + 1
		var i int
		if i2 > nextSize {
			i = i1
		} else {
			vi1, _ := h.arr.Get(i1 - 1)
			vi2, _ := h.arr.Get(i2 - 1)
			if h.comp(vi1, vi2) >= 0 {
				i = i1
			} else {
				i = i2
			}
		}
		vi, _ := h.arr.Get(i - 1)
		vpi, _ := h.arr.Get(pi - 1)
		if h.comp(vi, vpi) > 0 {
			h.arr.Swap(i-1, pi-1)
		}
		if i >= nextSize {
			break
		}
		pi = i
	}
}
