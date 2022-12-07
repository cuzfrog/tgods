package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedHashTable_Add(t *testing.T) {
	h := newLinkedHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int], 1)
	h.Add(2)
	old, found := h.AddTail(6)
	assert.False(t, found)
	old, found = h.AddTail(6)
	assert.Equal(t, 6, old)
	assert.True(t, found)
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

	m := make(map[int]node[int], 3)
	for _, b := range h.arr {
		for b != nil {
			m[b.Value()] = b
			b = b.Next()
		}
	}
	assert.Same(t, h.head, m[2].External())
	assert.Same(t, h.head.Next(), m[6].External())
	assert.Same(t, h.tail, m[3].External())

	assert.Equal(t, 3, h.size)
	assert.Equal(t, []int{2, 6, 3}, utils.SliceFrom[int](h))
}

func TestLinkedHashTable_Remove(t *testing.T) {
	h := newLinkedHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int], 1)
	assert.False(t, h.Remove(6))
	h.Add(2)
	h.Add(6)
	h.Add(3)
	assert.True(t, h.Remove(6))
	assert.Equal(t, 2, h.size)
	assert.Same(t, h.head, h.tail.Prev())
	assert.Same(t, h.tail, h.head.Next())
	assert.Equal(t, []int{2, 3}, utils.SliceFrom[int](h))

	h.Remove(2)
	assert.Equal(t, 1, h.Size())
	assert.Same(t, h.head, h.tail)
	assert.Nil(t, h.head.Prev())
	assert.Nil(t, h.tail.Next())
	h.Remove(3)
	assert.Equal(t, 0, h.size)
	assert.Nil(t, h.head)
	assert.Nil(t, h.tail)

	h.Add(5)
	h.Clear()
	assert.Nil(t, h.head)
	assert.Nil(t, h.tail)
	assert.Equal(t, 0, h.size)
}

func TestLinkedHashTable_Head_Tail(t *testing.T) {
	h := newLinkedHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int], OriginalOrder)
	v, found := h.Head()
	assert.False(t, found)
	v, found = h.Tail()
	assert.False(t, found)
	v, found = h.RemoveHead()
	assert.False(t, found)
	v, found = h.RemoveTail()
	assert.False(t, found)

	h.AddHead(3)
	v, found = h.AddHead(2)
	assert.False(t, found)
	h.Add(4)
	v, found = h.AddHead(2)
	assert.True(t, found)
	assert.Equal(t, 2, v)
	assert.Equal(t, []int{2, 3, 4}, utils.SliceFrom[int](h))
	v, found = h.Head()
	assert.True(t, found)
	assert.Equal(t, 2, v)
	v, found = h.RemoveHead()
	assert.Equal(t, 2, v)
	assert.Equal(t, []int{3, 4}, utils.SliceFrom[int](h))

	v, found = h.Tail()
	assert.Equal(t, 4, v)
	h.AddTail(5)
	assert.Equal(t, []int{3, 4, 5}, utils.SliceFrom[int](h))
	v, found = h.RemoveTail()
	assert.Equal(t, 5, v)
	assert.Equal(t, []int{3, 4}, utils.SliceFrom[int](h))
}
