package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIteratorForList(t *testing.T) {
	tests := []struct {
		name string
		l    types.List[int]
	}{
		{"circularArray", newCircularArrayOf[int](3, 5, 7).withRole(list)},
		{"linkedList", newLinkedListOf[int](3, 5, 7).withRole(list)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
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
		})
	}
}

func TestForStack(t *testing.T) {
	tests := []struct {
		name string
		l    types.Stack[int]
	}{
		{"circularArray", newCircularArrayOf[int]().withRole(stack)},
		{"linkedList", newLinkedListOf[int]().withRole(stack)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
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
			assert.False(t, it.Next())
		})
	}
}

func TestIteratorForQueue(t *testing.T) {
	tests := []struct {
		name string
		l    types.Queue[int]
	}{
		{"circularArray", newCircularArrayOf[int]().withRole(queue)},
		{"linkedList", newLinkedListOf[int]().withRole(queue)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
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
			assert.False(t, it.Next())
		})
	}
}

func TestIteratorForDeque(t *testing.T) {
	tests := []struct {
		name string
		l    types.Deque[int]
	}{
		{"circularArray", newCircularArrayOf[int]().withRole(deque)},
		{"linkedList", newLinkedListOf[int]().withRole(deque)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
			l.EnqueueLast(3)
			l.Enqueue(7)
			l.Enqueue(1)
			l.EnqueueLast(5)
			l.EnqueueLast(2)
			v, ok := l.DequeueFirst()
			assert.True(t, ok)
			assert.Equal(t, 1, v)

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
			assert.False(t, it.Next())
			assert.False(t, it.Next())
		})
	}
}
