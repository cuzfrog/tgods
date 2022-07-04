package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack_Add(t *testing.T) {
	s := NewArrayStack[int](3)
	s.Add(5)
	s.Add(3)
	assert.Equal(t, 2, s.Size())
	assert.ElementsMatch(t, []int{5, 3, 0}, s.arr)
	assert.Equal(t, 1, s.cur)
}

func TestArrayStack_Clear(t *testing.T) {
	s := NewArrayStack[int](3)
	s.Add(5)
	s.Add(3)
	s.Clear()
	assert.Equal(t, 0, s.Size())
	assert.Equal(t, -1, s.cur)
	assert.ElementsMatch(t, []int{0, 0, 0}, s.arr)
}

func TestArrayStack_Contains(t *testing.T) {
	s := NewArrayStack[int](3)
	assert.False(t, s.Contains(3))
	s.Add(5)
	s.Add(3)
	assert.True(t, s.Contains(5))
	assert.False(t, s.Contains(4))
}

func TestArrayStack_Peek(t *testing.T) {
	s := NewArrayStack[int](3)
	s.Add(5)
	s.Add(3)
	v, ok := s.Peek()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = s.Peek()
	assert.Equal(t, 3, v)
	assert.Equal(t, 2, s.Size())
}

func TestArrayStack_Pop(t *testing.T) {
	s := NewArrayStack[int](3)
	s.Add(5)
	s.Add(3)
	v, ok := s.Pop()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = s.Pop()
	assert.Equal(t, 5, v)
	v, ok = s.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, s.Size())
}

func TestArrayStack_Iterator(t *testing.T) {
	s := NewArrayStack[int](3)
	s.Add(5)
	s.Add(3)
	iter := s.Iterator()
	assert.True(t, iter.Next())
	assert.True(t, iter.Next())
	i, v := iter.Index(), iter.Value()
	assert.Equal(t, 1, i)
	assert.Equal(t, 3, v)
	assert.False(t, iter.Next())
	assert.PanicsWithValue(t, "index(2) out of range", func() { iter.Index() })
	assert.PanicsWithValue(t, "index(2) out of range", func() { iter.Value() })
}