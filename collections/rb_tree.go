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

func rbTreeInsert[T any](t *rbTree[T], d T) (T, bool) {
	r, found, nn, old := insertNode(t.root, d, t.comp)
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
	return old, found
}

func (t *rbTree[T]) insert(d T) bool {
	_, found := rbTreeInsert(t, d)
	return found
}

// delete returns the old value and true if found an existing entry
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

func (t *rbTree[T]) searchNode(elem T) *rbNode[T] {
	return searchNode(t.root, elem, t.comp)
}

func (t *rbTree[T]) firstNode() *rbNode[T] {
	if t.size == 0 {
		return nil
	}
	n := t.root
	for n.a != nil {
		n = n.a
	}
	return n
}

func (t *rbTree[T]) lastNode() *rbNode[T] {
	if t.size == 0 {
		return nil
	}
	n := t.root
	for n.b != nil {
		n = n.b
	}
	return n
}

// ======== SortedSet ========

func (t *rbTree[T]) Add(elem T) bool {
	t.insert(elem)
	return true
}

func (t *rbTree[T]) Contains(elem T) bool {
	n := t.searchNode(elem)
	return n != nil
}

func (t *rbTree[T]) Size() int {
	return t.size
}

func (t *rbTree[T]) Clear() {
	t.size = 0
	t.root = nil
}

func (t *rbTree[T]) Remove(elem T) bool {
	_, found := t.delete(elem)
	return found
}

func (t *rbTree[T]) First() (T, bool) {
	if t.size == 0 {
		return utils.Nil[T](), false
	}
	n := t.firstNode()
	return n.v, true
}

func (t *rbTree[T]) Last() (T, bool) {
	if t.size == 0 {
		return utils.Nil[T](), false
	}
	n := t.lastNode()
	return n.v, true
}
func (t *rbTree[T]) RemoveFirst() (T, bool) {
	n := t.removeFirstNode()
	if n != nil {
		return n.v, true
	}
	return utils.Nil[T](), false
}
func (t *rbTree[T]) removeFirstNode() *rbNode[T] {
	if t.size == 0 {
		return nil
	}
	if t.size == 1 {
		r := t.root
		t.Clear()
		return r
	}
	n := t.firstNode()
	nd := swapAndRemoveNode(n)
	t.size--
	return nd
}

func (t *rbTree[T]) RemoveLast() (T, bool) {
	n := t.removeLastNode()
	if n != nil {
		return n.v, true
	}
	return utils.Nil[T](), false
}
func (t *rbTree[T]) removeLastNode() *rbNode[T] {
	if t.size == 0 {
		return nil
	}
	if t.size == 1 {
		r := t.root
		t.Clear()
		return r
	}
	n := t.lastNode()
	nd := swapAndRemoveNode(n)
	t.size--
	return nd
}

//
//func (t *rbTree[T]) HeadSet(toElem T) types.SortedSet[T] {
//	nt := newRbTreeOfComp(t.comp)
//	it := t.Iterator()
//	for it.Next() {
//		v := it.Value()
//		if t.comp(v, toElem) < 0 {
//			nt.Add(v)
//		} else {
//			break
//		}
//	}
//	return nt
//}
//
//func (t *rbTree[T]) TailSet(fromElem T) types.SortedSet[T] {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *rbTree[T]) Higher(elem T) (T, bool) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *rbTree[T]) Lower(elem T) (T, bool) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *rbTree[T]) Ceiling(elem T) (T, bool) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *rbTree[T]) Floor(elem T) (T, bool) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *rbTree[T]) ReverseSet() types.SortedSet[T] {
//	//TODO implement me
//	panic("implement me")
//}
