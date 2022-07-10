package collections

import (
	"fmt"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

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

type linkedListIterator[T any] struct {
	index int
	start *node[T]
	cur   *node[T]
	next  func(n *node[T]) *node[T]
}

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
