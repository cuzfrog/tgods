package lists

type CircularArrayList[T comparable] struct {
	head int
	tail int
	arr  []T
	size int
}
