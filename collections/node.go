package collections

type node[T any] interface {
	Value() T
	SetValue(v T) T
	Prev() node[T]
	Next() node[T]
	External() node[T]
	SetPrev(prev node[T]) node[T]
	SetNext(next node[T]) node[T]
	SetExternal(x node[T]) node[T]
}

type slNode[T any] struct {
	v    T
	next node[T]
}

type dlNode[T any] struct {
	v    T
	prev *dlNode[T]
	next *dlNode[T]
}

type slxNode[T any] struct {
	v    T
	next *slxNode[T]
	x    node[T]
}

func (n *slNode[T]) Value() T {
	return n.v
}

func (n *slNode[T]) SetValue(v T) T {
	old := v
	n.v = v
	return old
}

func (n *slNode[T]) Prev() node[T] {
	panic("not supported")
}

func (n *slNode[T]) Next() node[T] {
	return n.next
}

func (n *slNode[T]) External() node[T] {
	return nil
}

func (n *slNode[T]) SetPrev(prev node[T]) node[T] {
	panic("not supported")
}

func (n *slNode[T]) SetNext(next node[T]) node[T] {
	old := n.next
	n.next = next
	return old
}

func (n *slNode[T]) SetExternal(x node[T]) node[T] {
	panic("not supported")
}
