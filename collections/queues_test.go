package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueProperties(t *testing.T) {
	tests := []struct {
		name string
		q    types.Queue[int]
	}{
		{"linkedList1", NewLinkedListQueue[int]()},
		{"linkedList2", NewLinkedListQueueOfEq[int](funcs.ValueEqual[int])},
		{"arrayList1", NewArrayListQueue[int]()},
		{"arrayList2", NewArrayListQueueOfSize[int](10)},
		{"arrayList3", NewArrayListQueueOfEq[int](10, funcs.ValueEqual[int])},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := test.q
			q.Enqueue(7)
			q.Enqueue(6)
			q.Enqueue(11)
			q.Enqueue(7)
			q.Enqueue(8)
			q.Enqueue(3)
			assert.Equal(t, []int{7, 6, 11, 7, 8, 3}, utils.SliceFrom[int](q))
			v, ok := q.Dequeue()
			assert.True(t, ok)
			assert.Equal(t, 7, v)
			v, ok = q.Dequeue()
			assert.Equal(t, 6, v)
			assert.Equal(t, []int{11, 7, 8, 3}, utils.SliceFrom[int](q))
		})
	}
}

func TestPriorityQueueProperties(t *testing.T) {
	tests := []struct {
		name string
		q    types.Queue[int]
	}{
		{"maxBinaryHeap", NewHeapMaxPriorityQueue[int]()},
		{"maxBinaryHeapWithComp", NewHeapPriorityQueue[int](funcs.ValueCompare[int])},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := test.q
			q.Enqueue(7)
			q.Enqueue(6)
			q.Enqueue(11)
			q.Enqueue(7)
			q.Enqueue(8)
			q.Enqueue(3)
			assert.Equal(t, []int{11, 8, 7, 7, 6, 3}, utils.SliceFrom[int](q))

			v, ok := q.Dequeue()
			assert.True(t, ok)
			assert.Equal(t, 11, v)
			v, ok = q.Dequeue()
			assert.Equal(t, 8, v)
			assert.Equal(t, []int{7, 7, 6, 3}, utils.SliceFrom[int](q))
		})
	}
	t.Run("maxBinaryHeap", func(t *testing.T) {
		q := NewHeapMinPriorityQueue[int]()
		q.Enqueue(7)
		q.Enqueue(6)
		q.Enqueue(11)
		q.Enqueue(7)
		q.Enqueue(8)
		q.Enqueue(3)

		v, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, 3, v)
		v, ok = q.Dequeue()
		assert.Equal(t, 3, v)

		q.Enqueue(1)
		assert.Equal(t, []int{1, 6, 7, 7, 8, 11}, utils.SliceFrom[int](q))
	})
}
