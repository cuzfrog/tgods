package removed

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](l, r T) (m T) {
	if l < r {
		m = l
	} else {
		m = r
	}
	return
}

func Max[T constraints.Ordered](l, r T) (m T) {
	if l > r {
		m = l
	} else {
		m = r
	}
	return
}

// Shuffle redistributes elems in the slice a using Knuth algorithm: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func Shuffle[T any](a []T, randFunc func(n int) int) {
	for i := len(a) - 1; i >= 1; i-- {
		j := randFunc(i)
		a[i], a[j] = a[j], a[i]
	}
}
