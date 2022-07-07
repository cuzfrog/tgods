package stacks

import (
	"fmt"
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/utils"
)

// assert ArrayStack implementation
var _ core.Bag[int] = (*ArrayStack[int])(nil)
var _ core.Stack[int] = (*ArrayStack[int])(nil)

// ArrayStack limited size array based stack
type ArrayStack[T comparable] struct {
	arr []T
	cur int
}

// NewArrayStack creates a stack with static size limit, allocating the underlying array eagerly
func NewArrayStack[T comparable](size int) *ArrayStack[T] {
	arr := make([]T, size)
	return &ArrayStack[T]{arr, -1}
}

func (s *ArrayStack[T]) Size() int {
	return s.cur + 1
}

func (s *ArrayStack[T]) Clear() {
	for i := 0; i <= s.cur; i++ {
		s.arr[i] = utils.Nil[T]()
	}
	s.cur = -1
}

func (s *ArrayStack[T]) Add(elem T) bool {
	s.cur++
	if s.cur >= len(s.arr) {
		return false
	}
	s.arr[s.cur] = elem
	return true
}

func (s *ArrayStack[T]) Pop() (elem T, found bool) {
	if s.cur < 0 {
		return elem, false
	}
	elem = s.arr[s.cur]
	s.arr[s.cur] = utils.Nil[T]()
	s.cur--
	return elem, true
}

func (s *ArrayStack[T]) Peek() (elem T, found bool) {
	if s.cur < 0 {
		return elem, false
	}
	return s.arr[s.cur], true
}

func (s *ArrayStack[T]) Contains(elem T) bool {
	for i := 0; i <= s.cur; i++ {
		if s.arr[i] == elem {
			return true
		}
	}
	return false
}

func (s *ArrayStack[T]) Iterator() core.Iterator[T] {
	return &iterator[T]{s, -1}
}

type iterator[T comparable] struct {
	s   *ArrayStack[T]
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
