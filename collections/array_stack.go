package collections

import (
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

func (s *arrayStack[T]) Push(elem T) bool {
	s.cur++
	if s.cur >= len(s.arr) {
		return false
	}
	s.arr[s.cur] = elem
	return true
}

func (s *arrayStack[T]) Add(elem T) bool {
	return s.Push(elem)
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
