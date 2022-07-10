package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularArrayList_Clone(t *testing.T) {
	l := newCircularArrayOf(3, 5, 7)
	nl := l.Clone()
	assert.NotSame(t, l, nl)
	assert.Equal(t, utils.SliceFrom[int](l), utils.SliceFrom[int](nl))
}
