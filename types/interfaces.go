package types

// WithHashAndEqual Constraint type for hash based data structures
type WithHashAndEqual[T any] interface {
	Hash() uint
	Equal(other T) bool
}

// WithCompare Constraint type for order based data structures, e.g. treeMap
type WithCompare[T any] interface {
	Compare(other T) int8
}

// WithEqual Constraint type for equality comparison, e.g. Collection.Contains
type WithEqual[T any] interface {
	Equal(other T) bool
}
