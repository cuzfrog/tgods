package trees

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/lists"
)

func bfTraverse[T comparable](n *rbNode[T]) core.List[T] {
	l := lists.NewLinkedList[T]()
	if n == nil {
		return l
	}
	nl := lists.NewLinkedListOfEqual[*rbNode[T]](func(a, b *rbNode[T]) bool { return a.v == b.v })
	nl.Add(n)
	for nl.Size() > 0 {
		next, _ := nl.PopHead()
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
