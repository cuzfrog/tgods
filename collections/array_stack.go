package collections

import (
	"fmt"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

// arrayStack limited size array based stack
type arrayStack[T comparable] struct {
	arr []T
	cur int
}

func newArrayStack[T comparable](size int) *arrayStack[T] {
	arr := make([]T, size)
	return &arrayStack[T]{arr, -1}
}

func (s *arrayStack[T]) Size() int {
	return s.cur + 1
}

func (s *arrayStack[T]) Clear() {
	for i := 0; i <= s.cur; i++ {
		s.arr[i] = utils.Nil[T]()
	}
	s.cur = -1
}

func (s *arrayStack[T]) Enstack(elem T) bool {
	s.cur++
	if s.cur >= len(s.arr) {
		return false
	}
	s.arr[s.cur] = elem
	return true
}

func (s *arrayStack[T]) Pop() (elem T, found bool) {
	if s.cur < 0 {
		return elem, false
	}
	elem = s.arr[s.cur]
	s.arr[s.cur] = utils.Nil[T]()
	s.cur--
	return elem, true
}

func (s *arrayStack[T]) Peek() (elem T, found bool) {
	if s.cur < 0 {
		return elem, false
	}
	return s.arr[s.cur], true
}

func (s *arrayStack[T]) Contains(elem T) bool {
	for i := 0; i <= s.cur; i++ {
		if s.arr[i] == elem {
			return true
		}
	}
	return false
}

func (s *arrayStack[T]) Iterator() types.Iterator[T] {
	return &iterator[T]{s, -1}
}

type iterator[T comparable] struct {
	s   *arrayStack[T]
	cur int
}

func (it *iterator[T]) Next() bool {
	it.cur++
	return it.cur <= it.s.cur
}

func (it *iterator[T]) Index() int {
	if it.cur > it.s.cur {
		panic(fmt.Sprintf("index(%d) out of range", it.cur))
	}
	return it.cur
}

func (it *iterator[T]) Value() T {
	if it.cur > it.s.cur {
		panic(fmt.Sprintf("index(%d) out of range", it.cur))
	}
	return it.s.arr[it.cur]
}
