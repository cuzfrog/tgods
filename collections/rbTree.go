package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

// assert rbTree implementation
var _ tree[int] = (*rbTree[int])(nil)

type rbTree[T any] struct {
	root *rbNode[T]
	size int
	comp types.Compare[T]
}

func newRbTree[T any](comp types.Compare[T]) *rbTree[T] {
	return &rbTree[T]{nil, 0, comp}
}

func (t *rbTree[T]) Insert(d T) bool {
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

func (t *rbTree[T]) Delete(d T) (T, bool) {
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
