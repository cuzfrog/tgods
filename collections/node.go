package collections

import "github.com/cuzfrog/tgods/types"

type node[T any] interface {
	Value() T
	SetValue(v T) T
	Prev() node[T]
	Next() node[T]
	External() node[T]
	SetPrev(prev node[T]) node[T]
	SetNext(next node[T]) node[T]
	SetExternal(x node[T]) node[T]

	Contains(elem T, eq types.Equal[T]) bool // checks if the elem is in the bucket
	Iterator() types.Iterator[T]
}

// assert type
var _ node[int] = (*slNode[int])(nil)
var _ node[int] = (*dlNode[int])(nil)
var _ node[int] = (*slxNode[int])(nil)
var _ node[int] = (*dlxNode[int])(nil)

func newSlNode[T any](v T, next node[T]) node[T] {
	return &slNode[T]{v, next}
}

func newDlNode[T any](v T, prev, next node[T]) node[T] {
	return &dlNode[T]{&slNode[T]{v, next}, prev}
}

func newSlxNode[T any](v T, next, external node[T]) node[T] {
	return &slxNode[T]{&slNode[T]{v, next}, external}
}

func newDlxNode[T any](v T, prev, next, external node[T]) node[T] {
	return &dlxNode[T]{&dlNode[T]{&slNode[T]{v, next}, prev}, external}
}

type slNode[T any] struct {
	v    T
	next node[T]
}

type dlNode[T any] struct {
	*slNode[T]
	prev node[T]
}

type slxNode[T any] struct {
	*slNode[T]
	x node[T]
}

type dlxNode[T any] struct {
	*dlNode[T]
	x node[T]
}

// ======== Implementations ========

func (n *slNode[T]) Value() T {
	return n.v
}

func (n *slNode[T]) SetValue(v T) T {
	old := v
	n.v = v
	return old
}

func (n *slNode[T]) Prev() node[T] {
	return nil
}

func (n *dlNode[T]) Prev() node[T] {
	return n.prev
}

func (n *slNode[T]) Next() node[T] {
	return n.next
}

func (n *slNode[T]) External() node[T] {
	return nil
}

func (n *slxNode[T]) External() node[T] {
	return n.x
}

func (n *slNode[T]) SetPrev(_ node[T]) node[T] {
	return nil
}

func (n *dlNode[T]) SetPrev(prev node[T]) node[T] {
	old := n.prev
	n.prev = prev
	return old
}

func (n *slNode[T]) SetNext(next node[T]) node[T] {
	old := n.next
	n.next = next
	return old
}

func (n *slNode[T]) SetExternal(_ node[T]) node[T] {
	return nil
}

func (n *slxNode[T]) SetExternal(x node[T]) node[T] {
	old := n.x
	n.x = x
	return old
}
