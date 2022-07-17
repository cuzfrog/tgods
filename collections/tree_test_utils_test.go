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
	nl := NewLinkedListQueueOfEq[*rbNode[T]](nodeValueEqual[T])
	nl.Enqueue(n)
	for nl.Size() > 0 {
		next, _ := nl.Dequeue()
		l.Add(next.v)
		if next.a != nil {
			nl.Enqueue(next.a)
		}
		if next.b != nil {
			nl.Enqueue(next.b)
		}
	}
	return l
}

func Test_bfTraverse(t *testing.T) {
	tree := newRbTreeOf(1, 2, 3, 4, 5)
	l := bfTraverse(tree.root)
	assert.Equal(t, []int{2, 1, 4, 3, 5}, utils.SliceFrom[int](l))
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
