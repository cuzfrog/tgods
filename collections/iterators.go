package collections

import (
	"fmt"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

// ======== arrayStack ========

func (s *arrayStack[T]) Iterator() types.Iterator[T] {
	return &arrayStackIterator[T]{s, s.cur + 1}
}

type arrayStackIterator[T comparable] struct {
	s   *arrayStack[T]
	cur int
}

func (it *arrayStackIterator[T]) Next() bool {
	it.cur--
	return it.cur >= 0
}

func (it *arrayStackIterator[T]) Index() int {
	return it.s.Size() - 1 - it.cur
}

func (it *arrayStackIterator[T]) Value() T {
	return it.s.arr[it.cur]
}

// ======== circularArray ========

func (l *circularArray[T]) Iterator() types.Iterator[T] {
	var getArrIndex func(i int) (int, bool)
	if l.r == list {
		getArrIndex = func(i int) (int, bool) { return l.toArrIndex(i) }
	} else if l.r == stack || l.r == queue || l.r == deque {
		getArrIndex = func(i int) (int, bool) { return l.toArrIndex(l.size - i - 1) }
	} else {
		panic(fmt.Sprintf("circularArray only implement classes [list(%d), stack(%d), queue(%d), deque(%d)], but the role is '%d'", list, stack, queue, deque, l.r))
	}
	return &circularArrayIterator[T]{l, -1, -1, getArrIndex}
}

type circularArrayIterator[T any] struct {
	l           *circularArray[T]
	index       int
	arrIndex    int
	getArrIndex func(i int) (int, bool)
}

func (it *circularArrayIterator[T]) Next() bool {
	it.index++
	arrIndex, ok := it.getArrIndex(it.index) // TODO: optimize
	it.arrIndex = arrIndex
	return ok
}

// Index returns current index, will not fail when invalid, should be guarded by Next()
func (it *circularArrayIterator[T]) Index() int {
	return it.index
}

func (it *circularArrayIterator[T]) Value() T {
	return it.l.arr[it.arrIndex]
}

// ======== linkedList ========

func (l *linkedList[T]) Iterator() types.Iterator[T] {
	var next func(n *node[T]) *node[T]
	var start *node[T]
	if l.r == list || l.r == stack {
		next = func(n *node[T]) *node[T] { return n.next }
		start = l.head
	} else if l.r == queue || l.r == deque {
		next = func(n *node[T]) *node[T] { return n.prev }
		start = l.tail
	} else {
		panic(fmt.Sprintf("linkedList only implement classes [list(%d), stack(%d), queue(%d), deque(%d)], but the role is '%d'", list, stack, queue, deque, l.r))
	}
	return &linkedListIterator[T]{-1, start, nil, next}
}

type linkedListIterator[T any] struct {
	index int
	start *node[T]
	cur   *node[T]
	next  func(n *node[T]) *node[T]
}

func (it *linkedListIterator[T]) Next() bool {
	if it.start == nil {
		return false
	}
	if it.index < 0 {
		it.cur = it.start
	} else if it.cur != nil {
		it.cur = it.next(it.cur)
	}
	it.index++
	return it.cur != nil
}

func (it *linkedListIterator[T]) Index() int {
	return it.index
}

func (it *linkedListIterator[T]) Value() T {
	if it.cur == nil {
		return utils.Nil[T]()
	}
	return it.cur.v
}

// ======== binaryHeap ========

func (h *binaryHeap[T]) Iterator() types.Iterator[T] {
	return &binaryHeapIterator[T]{h.Clone(), -1, utils.Nil[T]()}
}

type binaryHeapIterator[T any] struct {
	q     types.Queue[T]
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

func forEach[T any](c types.Collection[T], fn func(index int, v T)) {
	it := c.Iterator()
	for it.Next() {
		fn(it.Index(), it.Value())
	}
}

func (s *arrayStack[T]) Each(fn func(index int, elem T)) {
	forEach[T](s, fn)
}

func (h *binaryHeap[T]) Each(fn func(index int, elem T)) {
	forEach[T](h, fn)
}

func (h *circularArray[T]) Each(fn func(index int, elem T)) {
	forEach[T](h, fn)
}

func (h *linkedList[T]) Each(fn func(index int, elem T)) {
	forEach[T](h, fn)
}
