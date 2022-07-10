package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

type binaryHeap[T any] struct {
	arr  types.ArrayList[T]
	comp funcs.Compare[T]
}

func newBinaryHeap[T any](comp funcs.Compare[T]) *binaryHeap[T] {
	return &binaryHeap[T]{newCircularArrayOfEq(0, funcs.CompToEq(comp)), comp}
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

func (h *binaryHeap[T]) Clone() types.Queue[T] {
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
