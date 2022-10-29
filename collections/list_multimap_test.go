package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayListMultiMap_PutSingle_AllValues(t *testing.T) {
	mm := newArrayListMultiMap[string, int](funcs.NewStrHash(), funcs.ValueEqual[string])
	assert.Equal(t, 0, mm.Size())

	mm.PutSingle("a", 1)
	mm.PutSingle("a", 1)
	mm.PutSingle("a", 2)
	mm.PutSingle("b", 3)
	assert.Equal(t, 4, mm.Size())
	assert.Equal(t, 2, mm.KeySize())
	assert.ElementsMatch(t, []int{1, 1, 2, 3}, utils.SliceFrom[int](mm.AllValues()))

	assert.True(t, mm.Contains(EntryOf[string, types.List[int]]("a", NewArrayListOf(111))))
}

func TestArrayListMultiMap_Put_Add(t *testing.T) {
	mm := newArrayListMultiMap[string, int](funcs.NewStrHash(), funcs.ValueEqual[string])

	old, found := mm.Put("a", NewArrayListOf(1, 2))
	assert.False(t, found)
	assert.Equal(t, 2, mm.Size())
	assert.Equal(t, 1, mm.KeySize())

	old, found = mm.Put("a", NewArrayListOf(1))
	assert.True(t, found)
	assert.Equal(t, []int{1, 2}, utils.SliceFrom[int](old))
	assert.Equal(t, 1, mm.Size())

	mm.Put("b", NewArrayListOf(1, 3))
	assert.Equal(t, 3, mm.Size())
	assert.Equal(t, 2, mm.KeySize())

	mm.Add(EntryOf[string, types.List[int]]("b", NewArrayListOf(1, 3, 4)))
	assert.Equal(t, 4, mm.Size())
}
