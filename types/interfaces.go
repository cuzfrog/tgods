package types

type HashAndEqual[T any] interface {
	Hash() uint
	Equal(other T) bool
}
