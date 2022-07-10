package utils

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSlice(t *testing.T) {
	c := &mockCollection[int]{}
	c.arr = []int{1, 4, 3, 2}
	c.size = 4
	arr := SliceFrom[int](c)
	assert.Equal(t, []int{1, 4, 3, 2}, arr)
}

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a2 := make([]int, 5)
	copy(a2, a)
	Shuffle(a2, rand.Intn)
	assert.NotEqual(t, a, a2)
}
