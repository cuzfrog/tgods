package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"golang.org/x/exp/constraints"
)

func newRbTreeOf[T constraints.Ordered](values ...T) *rbTree[T] {
	t := newRbTree[T](funcs.ValueCompare[T])
	for _, v := range values {
		t.Insert(v)
	}
	return t
}

func (t *rbTree[T]) Add(elem T) bool {
	t.Insert(elem)
	return true
}

func (t *rbTree[T]) Contains(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Iterator() types.Iterator[T] {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Size() int {
	return t.size
}

func (t *rbTree[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Remove(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) First() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Last() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) RemoveFirst() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) RemoveLast() (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) HeadSet(toElem T) types.SortedSet[T] {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) TailSet(fromElem T) types.SortedSet[T] {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Higher(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Lower(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Ceiling(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) Floor(elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (t *rbTree[T]) ReverseSet() types.SortedSet[T] {
	//TODO implement me
	panic("implement me")
}
