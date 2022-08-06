package utils

import (
	"github.com/cuzfrog/tgods/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](1, 4, 3, 2)
	arr := SliceFrom[int](c)
	assert.Equal(t, []int{1, 4, 3, 2}, arr)
}

func TestSliceProjection(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](1, 4, 3, 2)
	arr := SliceProject[int, int](c, func(v int) int { return v + 1 })
	assert.Equal(t, []int{2, 5, 4, 3}, arr)
}
