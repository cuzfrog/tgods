package utils

import (
	"github.com/cuzfrog/tgods/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAll(t *testing.T) {
	c1 := mocks.NewMockCollectionOf(1, 2, 3)
	c2 := mocks.NewMockCollection[int](10)
	addedCnt := AddAll[int](c1, c2)
	assert.Equal(t, 3, addedCnt)
	assert.ElementsMatch(t, SliceFrom[int](c2), []int{1, 2, 3})
}
