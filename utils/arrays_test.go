package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	c := &mockCollection[int]{}
	c.arr = []int{1, 4, 3, 2}
	c.size = 4
	arr := SliceFrom[int](c)
	assert.Equal(t, []int{1, 4, 3, 2}, arr)
}
