package utils

import (
	"github.com/cuzfrog/tgods/lists"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSlice(t *testing.T) {
	l := lists.NewLinkedList(1, 4, 3, 2)
	arr := SliceFrom(l.Iterator(), l.Size())
	assert.Equal(t, []int{1, 4, 3, 2}, arr)

	arr = SliceFrom(l.Iterator(), 3)
	assert.Equal(t, []int{1, 4, 3}, arr)

	arr = SliceFrom(l.Iterator(), 5)
	assert.Equal(t, 4, len(arr))
	assert.Equal(t, 5, cap(arr))
}

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a2 := make([]int, 5)
	copy(a2, a)
	Shuffle(a2, rand.Intn)
	assert.NotEqual(t, a, a2)
}
