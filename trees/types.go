package trees

type Tree[T any] interface {
	Insert(d T) bool // Insert returns true if the entry is found in the tree by func Compare
}
