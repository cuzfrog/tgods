package queues

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/lists"
	"github.com/cuzfrog/tgods/utils"
	"golang.org/x/exp/constraints"
)

// assert HeapPriorityQueue implementation
var _ core.Queue[int] = (*HeapPriorityQueue[int])(nil)
var _ core.Bag[int] = (*HeapPriorityQueue[int])(nil)
var _ core.Iterable[int] = (*HeapPriorityQueue[int])(nil)

type HeapPriorityQueue[T comparable] struct {
	arr  core.ArrayList[T]
	comp utils.Compare[T]
}

func NewHeapPriorityQueue[T comparable](comparator utils.Compare[T]) *HeapPriorityQueue[T] {
	return &HeapPriorityQueue[T]{lists.NewCircularArrayList[T](), comparator}
}

func NewHeapPriorityQueueForMaxValue[T constraints.Ordered]() *HeapPriorityQueue[T] {
	return &HeapPriorityQueue[T]{lists.NewCircularArrayList[T](), utils.CompareOrdered[T]}
}
func NewHeapPriorityQueueForMinValue[T constraints.Ordered]() *HeapPriorityQueue[T] {
	fn := func(a, b T) int8 { return utils.CompareOrdered(b, a) }
	return &HeapPriorityQueue[T]{lists.NewCircularArrayList[T](), fn}
}

func (h *HeapPriorityQueue[T]) Size() int {
	return h.arr.Size()
}

func (h *HeapPriorityQueue[T]) Clear() {
	h.arr.Clear()
}

func (h *HeapPriorityQueue[T]) Add(elem T) bool {
	ret := h.arr.Add(elem)
	h.swim()
	return ret
}

func (h *HeapPriorityQueue[T]) Pop() (T, bool) {
	h.sink()
	return h.arr.Pop()
}

func (h *HeapPriorityQueue[T]) Peek() (T, bool) {
	return h.arr.Head()
}

// Contains delegates to underlying array, O(n)
func (h *HeapPriorityQueue[T]) Contains(elem T) bool {
	return h.arr.Contains(elem)
}

type hpqIterator[T comparable] struct {
	h     *HeapPriorityQueue[T]
	index int
	cur   T
}

func (it *hpqIterator[T]) Next() bool {
	if it.h.Size() <= 0 {
		return false
	}
	it.index++
	it.cur, _ = it.h.Pop()
	return true
}

// Index returns current index from iteration order, starting from 0.
// If Next return false, the number returned is meaningless, it must be guarded by Next
func (it *hpqIterator[T]) Index() int {
	return it.index
}

// Value returns current value.
// If Next return false, the value returned is meaningless, it must be guarded by Next
func (it *hpqIterator[T]) Value() T {
	return it.cur
}

func (h *HeapPriorityQueue[T]) Iterator() core.Iterator[T] {
	return &hpqIterator[T]{h, -1, utils.Nil[T]()}
}

// swim reheapifies by checking and moving up the last element of the arr, this should be called after adding
// i - child node index
// pi - parent node index
func (h *HeapPriorityQueue[T]) swim() {
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
func (h *HeapPriorityQueue[T]) sink() {
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
		h.arr.Swap(i-1, pi-1)
		if i >= nextSize {
			break
		}
		pi = i
	}
}
