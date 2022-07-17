package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"golang.org/x/exp/constraints"
)

type rbTree[T any] struct {
	root *rbNode[T]
	size int
	comp types.Compare[T]
}

func newRbTreeOfComp[T any](comp types.Compare[T]) *rbTree[T] {
	return &rbTree[T]{nil, 0, comp}
}

func newRbTreeOf[T constraints.Ordered](values ...T) *rbTree[T] {
	t := newRbTreeOfComp[T](funcs.ValueCompare[T])
	for _, v := range values {
		t.insert(v)
	}
	return t
}

func (t *rbTree[T]) insert(d T) bool {
	r, found, nn := insertNode(t.root, d, t.comp)
	for true {
		nn = insertionRebalance(nn)
		if nn == nil {
			break
		}
	}
	for r.p != nil { // in case r is rotated down
		r = r.p
	}
	t.root = r
	if !found {
		t.size++
	}
	return found
}

func (t *rbTree[T]) delete(d T) (T, bool) {
	nd, found := deleteNode[T](t.root, d, t.comp)
	if found {
		t.size--
		if t.root.p != nil {
			t.root = t.root.p
		}
		return nd.v, true
	}
	return utils.Nil[T](), false
}

// ======== SortedSet ========

func (t *rbTree[T]) Add(elem T) bool {
	t.insert(elem)
	return true
}

func (t *rbTree[T]) Contains(elem T) bool {
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
