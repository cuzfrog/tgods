package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRbTree_Add(t *testing.T) {
	s := newRbTreeOf(3, 4, 5)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, utils.SliceFrom[int](s))
}
