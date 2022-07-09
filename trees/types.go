package trees

import "github.com/cuzfrog/tgods/core"

type Traversable[T any] interface {
	InorderIterator() core.Iterator[T]
}

type Tree[T any] interface {
	Insert(d T) bool // Insert returns true if the entry is found in the tree by func Compare
}
