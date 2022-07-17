package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

type treeSet[T any] struct {
	t *rbTree[T]
}

func newTreeSetOf[T comparable](values ...T) *treeSet[T] {
	t := newRbTree[T](funcs.ValueCompare[T])
	for _, v := range values {
		t.Insert(v)
	}
	return &treeSet[T]{t}
}

func (s *treeSet[T]) Add(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Contains(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Iterator() types.Iterator[T] {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Each(f func(index int, elem T)) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Remove(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) First() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Last() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) RemoveFirst() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) RemoveLast() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) HeadSet(toElem T) types.SortedSet[T] {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) TailSet(fromElem T) types.SortedSet[T] {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Higher(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Lower(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Ceiling(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) Floor(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (s *treeSet[T]) ReverseSet() types.SortedSet[T] {
	//TODO implement me
	panic("implement me")
}
