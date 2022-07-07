package utils

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
