package generic_collections

type Node[T any] struct {
	v    T
	prev *Node[T]
	next *Node[T]
}
