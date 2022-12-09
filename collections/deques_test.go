package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeProperties(t *testing.T) {
	tests := []struct {
		name string
		q    types.Deque[int]
	}{
		{"linkedList1", NewLinkedListDeque[int]()},
		{"linkedList2", NewLinkedListDequeOfEq[int](funcs.ValueEqual[int])},
		{"arrayList1", NewArrayListDeque[int]()},
		{"arrayList2", NewArrayListDequeOfSize[int](10)},
		{"arrayList3", NewArrayListDequeOfEq[int](10, funcs.ValueEqual[int])},
		{"arrayList4", NewArrayListDequeOfSizeP[int](10, NoAutoSizing)},
		{"arrayList5", NewArrayListDequeOfEqP[int](10, funcs.ValueEqual[int], NoAutoSizing)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := test.q
			q.Enqueue(7)
			q.Enqueue(6)
			q.Enqueue(11)
			q.Enqueue(7)
			q.Enqueue(8)
			q.EnqueueLast(3)
			assert.Equal(t, []int{3, 7, 6, 11, 7, 8}, utils.SliceFrom[int](q))
			v, ok := q.Dequeue()
			assert.True(t, ok)
			assert.Equal(t, 3, v)
			v, ok = q.Peek()
			assert.Equal(t, 7, v)
			v, ok = q.Dequeue()
			assert.Equal(t, 7, v)
			v, ok = q.DequeueFirst()
			assert.Equal(t, 8, v)

			assert.Equal(t, []int{6, 11, 7}, utils.SliceFrom[int](q))

			q.Push(33)
			assert.Equal(t, []int{33, 6, 11, 7}, utils.SliceFrom[int](q))
			v, ok = q.Pop()
			assert.Equal(t, 33, v)
		})
	}
}
