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
