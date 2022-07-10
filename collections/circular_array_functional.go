package collections

import "github.com/cuzfrog/tgods/types"

func (l *circularArray[T]) Clone() types.ArrayList[T] {
	arr := make([]T, l.size)
	copy(arr, l.arr)
	return &circularArray[T]{l.start, l.end, arr, l.size, l.comp, l.r}
}
