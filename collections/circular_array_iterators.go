package collections

import (
	"fmt"
	"github.com/cuzfrog/tgods/types"
)

func (l *circularArray[T]) Iterator() types.Iterator[T] {
	var getArrIndex func(i int) (int, bool)
	if l.cl == list {
		getArrIndex = func(i int) (int, bool) { return l.toArrIndex(i) }
	} else if l.cl == stack || l.cl == queue || l.cl == deque {
		getArrIndex = func(i int) (int, bool) { return l.toArrIndex(l.size - i - 1) }
	} else {
		panic(fmt.Sprintf("circularArray only implement classes [list(%d), stack(%d), queue(%d), deque(%d)], but the class is '%d'", list, stack, queue, deque, l.cl))
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
