package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedHashTable_Add(t *testing.T) {
	h := newLinkedHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int])
	h.Add(2)
	h.Add(6)
	h.Add(3)
	assert.Equal(t, 2, h.head.Value())
	assert.Equal(t, 6, h.head.Next().Value())
	assert.Equal(t, 3, h.head.Next().Next().Value())
	assert.Equal(t, 3, h.tail.Value())
	assert.Equal(t, 6, h.tail.Prev().Value())
	assert.Equal(t, 2, h.tail.Prev().Prev().Value())
	assert.Same(t, h.head.Next(), h.tail.Prev())
	assert.Same(t, h.head, h.tail.Prev().Prev())
	assert.Same(t, h.tail, h.head.Next().Next())

	assert.Equal(t, 3, h.size)
	assert.Equal(t, []int{2, 6, 3}, utils.SliceFrom[int](h))
}
