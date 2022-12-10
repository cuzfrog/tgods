package types

type WithHashAndEqual[T any] interface {
	Hash() uint
	Equal(other T) bool
}

type WithCompare[T any] interface {
	Compare(other T) int8
}

type WithEqual[T any] interface {
	Equal(other T) bool
}
