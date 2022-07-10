package types

type Cloneable[T any, C Collection[T]] interface {
	Clone() C
}
