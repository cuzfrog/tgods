package core

type Cloneable[T any, C Collection[T]] interface {
	Clone() C
}
