package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHeapPriorityQueue(t *testing.T) {
	type obj struct {
		v string
	}

	comp := func(a, b *obj) int8 {
		if a.v > b.v {
			return 1
		} else if a.v < b.v {
			return -1
		}
		return 0
	}
	q := newBinaryHeap[*obj](comp)
	q.Enqueue(&obj{"1"})
	assert.Equal(t, 1, q.Size())
	v, _ := q.Peek()
	assert.Equal(t, "1", v.v)
}

func TestHeapPriorityQueue(t *testing.T) {
	q := NewHeapMaxPriorityQueue[int]()
	q.Enqueue(6)
	q.Add(3)
	assert.Equal(t, []int{6, 3}, utils.SliceFrom[int](q))
	q.Enqueue(7)
	assert.Equal(t, []int{7, 6, 3}, utils.SliceFrom[int](q))
	q.Enqueue(3)
	assert.Equal(t, []int{7, 6, 3, 3}, utils.SliceFrom[int](q))
	assert.Equal(t, 4, q.Size())
	assert.True(t, q.Contains(6))
	assert.False(t, q.Contains(5))

	v, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 7, v)
	v, _ = q.Dequeue()
	assert.Equal(t, 7, v)
	assert.Equal(t, []int{6, 3, 3}, utils.SliceFrom[int](q))
	v, _ = q.Dequeue()
	assert.Equal(t, 6, v)
	assert.Equal(t, []int{3, 3}, utils.SliceFrom[int](q))

	v, _ = q.Dequeue()
	assert.Equal(t, 3, v)
	v, _ = q.Dequeue()
	assert.Equal(t, 3, v)
	v, ok = q.Dequeue()
	assert.False(t, ok)
	q.Enqueue(2)
	q.Clear()
	v, ok = q.Peek()
	assert.False(t, ok)
}

func TestHeapPriorityQueue_swim(t *testing.T) {
	arr := newCircularArrayOf("t", "s", "r", "p", "n", "o", "a", "e", "i", "q", "w")
	q := &binaryHeap[string]{arr, funcs.ValueCompare[string]}
	q.swim()
	assert.Equal(t, []string{"w", "t", "r", "p", "s", "o", "a", "e", "i", "q", "n"}, utils.SliceFrom[string](arr))

	q.Enqueue("s")
	q.swim()
	assert.Equal(t, []string{"w", "t", "s", "p", "s", "r", "a", "e", "i", "q", "n", "o"}, utils.SliceFrom[string](arr))

}

func TestHeapPriorityQueue_sink(t *testing.T) {
	arr := newCircularArrayOf("t", "s", "r", "p", "n", "o", "a", "e", "i", "h", "g")
	q := &binaryHeap[string]{arr, funcs.ValueCompare[string]}
	q.sink()
	assert.Equal(t, []string{"s", "p", "r", "i", "n", "o", "a", "e", "g", "h", "t"}, utils.SliceFrom[string](arr))
	arr.RemoveTail()
	q.sink()
	assert.Equal(t, []string{"r", "p", "o", "i", "n", "h", "a", "e", "g", "s"}, utils.SliceFrom[string](arr))
}
