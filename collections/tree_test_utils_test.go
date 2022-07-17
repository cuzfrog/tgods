package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func nodeValueEqual[T comparable](a, b *rbNode[T]) bool { return a.v == b.v }

func bfTraverse[T comparable](n *rbNode[T]) types.List[T] {
	l := NewLinkedList[T]()
	if n == nil {
		return l
	}
	nl := newLinkedListOfEq[*rbNode[T]](nodeValueEqual[T])
	nl.Add(n)
	for nl.Size() > 0 {
		next, _ := nl.RemoveHead()
		l.Add(next.v)
		if next.a != nil {
			nl.Add(next.a)
		}
		if next.b != nil {
			nl.Add(next.b)
		}
	}
	return l
}

func dfInorderTraverse[T comparable](n *rbNode[T]) types.List[T] {
	l := NewLinkedList[T]()
	if n == nil {
		return l
	}
	s := NewCircularArrayStackOfEq[*rbNode[T]](0, nodeValueEqual[T])
	s.Push(n)
	cur := n.a
	for cur != nil || s.Size() > 0 {
		for cur != nil {
			s.Push(cur)
			cur = cur.a
		}
		np, _ := s.Pop()
		l.Add(np.v)
		cur = np.b
	}
	return l
}

func Test_dfInorderTraverse(t *testing.T) {
	tree := newRbTreeOf(1, 2, 3, 4, 5)
	l := dfInorderTraverse(tree.root)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, utils.SliceFrom[int](l))
}
