package collections

import (
	"github.com/cuzfrog/tgods/types"
)

func bfTraverse[T comparable](n *rbNode[T]) types.List[T] {
	l := NewLinkedList[T]()
	if n == nil {
		return l
	}
	nl := newLinkedListOfEq[*rbNode[T]](func(a, b *rbNode[T]) bool { return a.v == b.v })
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
