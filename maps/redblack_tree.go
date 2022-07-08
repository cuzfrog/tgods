package maps

import "github.com/cuzfrog/tgods/utils"

const red, black = true, false

type color bool

type rbTree[T any] struct {
	root *node[T]
	size int
	comp utils.Comparator[T]
}

type node[T any] struct {
	v T
	a *node[T]
	b *node[T]
	c bool
}

func rotateLeft[T any](n *node[T]) (r *node[T]) {
	if n.b == nil {
		return n
	}
	r = n.b
	n.b = r.a
	r.a = n
	r.c, n.c = n.c, r.c
	return
}

func rotateRight[T any](n *node[T]) (r *node[T]) {
	if n.a == nil {
		return n
	}
	r = n.a
	n.a = r.b
	r.b = n
	r.c, n.c = n.c, r.c
	return
}

func isRed[T any](n *node[T]) bool {
	return n != nil && n.c
}
