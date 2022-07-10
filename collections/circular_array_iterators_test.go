package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularArrayList_IteratorForList(t *testing.T) {
	l := newCircularArrayOf(3, 5, 7)
	l.cl = list
	it := l.Iterator()
	assert.True(t, it.Next())
	assert.Equal(t, 0, it.Index())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Index())
	assert.Equal(t, 5, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Index())
	assert.Equal(t, 7, it.Value())
	assert.False(t, it.Next())
	assert.False(t, it.Next())

	l.Clear()
	it = l.Iterator()
	assert.False(t, it.Next())
}

func TestCircularArrayList_IteratorForStack(t *testing.T) {
	l := newCircularArrayOf[int]()
	l.cl = stack
	l.Enstack(3)
	l.Enstack(5)
	l.Enstack(2)
	it := l.Iterator()
	assert.True(t, it.Next())
	assert.Equal(t, 0, it.Index())
	assert.Equal(t, 2, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Index())
	assert.Equal(t, 5, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Index())
	assert.Equal(t, 3, it.Value())
	assert.False(t, it.Next())
}

func TestCircularArrayList_IteratorForQueue(t *testing.T) {
	l := newCircularArrayOf[int]()
	l.cl = queue
	l.Enqueue(3)
	l.Enqueue(5)
	l.Enqueue(2)
	it := l.Iterator()
	assert.True(t, it.Next())
	assert.Equal(t, 0, it.Index())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Index())
	assert.Equal(t, 5, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Index())
	assert.Equal(t, 2, it.Value())
	assert.False(t, it.Next())
}

func TestCircularArrayList_IteratorForDeque(t *testing.T) {
	l := newCircularArrayOf[int]()
	l.cl = deque
	l.EnqueueLast(3)
	l.Enqueue(7)
	l.Enqueue(1)
	l.EnqueueLast(5)
	l.EnqueueLast(2)
	it := l.Iterator()
	assert.True(t, it.Next())
	assert.Equal(t, 0, it.Index())
	assert.Equal(t, 2, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Index())
	assert.Equal(t, 5, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Index())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Index())
	assert.Equal(t, 7, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Index())
	assert.Equal(t, 1, it.Value())
	assert.False(t, it.Next())
}
