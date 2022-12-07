package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"golang.org/x/exp/constraints"
)

type enumSet[T constraints.Integer] struct {
	enumMap[T, interface{}]
}

func newEnumSet[T constraints.Integer](max T, values ...T) *enumSet[T] {
	m := newEnumMap[T, interface{}](max)
	s := &enumSet[T]{*m}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s *enumSet[T]) Add(elem T) bool {
	_, existing := s.Replace(elem)
	return !existing
}

func (s *enumSet[T]) Replace(elem T) (T, bool) {
	e := s.arr[elem]
	if e == nil {
		s.size++
		s.arr[elem] = keyEntry[T, interface{}]{elem}
		return utils.Nil[T](), false
	}
	return e.Key(), true
}

func (s *enumSet[T]) Contains(elem T) bool {
	e := s.arr[elem]
	return e != nil
}

func (s *enumSet[T]) Remove(elem T) bool {
	e := s.arr[elem]
	if e != nil {
		s.size--
		s.arr[elem] = nil
		return true
	}
	return false
}

func (s *enumSet[T]) First() (T, bool) {
	for i := 0; i < len(s.arr); i++ {
		e := s.arr[i]
		if e != nil {
			return e.Key(), true
		}
	}
	return utils.Nil[T](), false
}

func (s *enumSet[T]) Last() (T, bool) {
	for i := len(s.arr) - 1; i >= 0; i-- {
		e := s.arr[i]
		if e != nil {
			return e.Key(), true
		}
	}
	return utils.Nil[T](), false
}

func (s *enumSet[T]) RemoveFirst() (T, bool) {
	for i := 0; i < len(s.arr); i++ {
		e := s.arr[i]
		if e != nil {
			s.size--
			s.arr[i] = nil
			return e.Key(), true
		}
	}
	return utils.Nil[T](), false
}

func (s *enumSet[T]) RemoveLast() (T, bool) {
	for i := len(s.arr) - 1; i >= 0; i-- {
		e := s.arr[i]
		if e != nil {
			s.size--
			s.arr[i] = nil
			return e.Key(), true
		}
	}
	return utils.Nil[T](), false
}
