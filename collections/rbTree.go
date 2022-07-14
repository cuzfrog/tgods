package collections

import "github.com/cuzfrog/tgods/types"

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
	r, found, nn := insert(t.root, d, t.comp)
	for true {
		nn = rebalance(nn)
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
