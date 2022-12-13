package types

// WithHashAndEqual Constraint type for hash based data structures
type WithHashAndEqual[T any] interface {
	Hash() uint
	Equal(other T) bool
}

// WithCompare Constraint type for order based data structures, e.g. treeMap
type WithCompare[T any] interface {
	Compare(other T) int8 // returns 0 when equal, 1 when this > other, -1 when this < other
}

// WithEqual Constraint type for equality comparison, e.g. Collection.Contains
type WithEqual[T any] interface {
	Equal(other T) bool
}

// WithLess Constraint type for sorting, e.g. utils.SortC
type WithLess[T any] interface {
	Less(other T) bool // returns true if this < other
}
