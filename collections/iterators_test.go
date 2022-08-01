package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/mocks"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
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
		{"arrayStack", newArrayStack[int](3)},
		{"circularArray", newCircularArrayOf[int]().withRole(stack)},
		{"linkedList", newLinkedListOf[int]().withRole(stack)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
			l.Push(3)
			l.Push(5)
			l.Push(2)
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

func TestIteratorForSortedSet(t *testing.T) {
	tests := []struct {
		name string
		s    types.SortedSet[int]
	}{
		{"rbTree", newRbTreeOf[int]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.s
			s.Add(1)
			s.Add(3)
			s.Add(2)
			s.Add(4)
			ok := s.Add(5)
			assert.True(t, ok)

			it := s.Iterator()
			assert.True(t, it.Next())
			assert.Equal(t, 0, it.Index())
			assert.Equal(t, 1, it.Value())
			assert.True(t, it.Next())
			assert.Equal(t, 1, it.Index())
			assert.Equal(t, 2, it.Value())
			assert.True(t, it.Next())
			assert.Equal(t, 2, it.Index())
			assert.Equal(t, 3, it.Value())
			assert.True(t, it.Next())
			assert.Equal(t, 3, it.Index())
			assert.Equal(t, 4, it.Value())
			assert.True(t, it.Next())
			assert.Equal(t, 4, it.Index())
			assert.Equal(t, 5, it.Value())
			assert.False(t, it.Next())
			assert.False(t, it.Next())
			assert.Equal(t, 0, it.Value())
		})
	}
}

func TestIteratorForSet(t *testing.T) {
	tests := []struct {
		name string
		s    types.Set[int]
	}{
		{"rbTree", newRbTreeOf[int]()},
		{"hashTable", newHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int])},
		{"linkedHashTable", newLinkedHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int])},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.s
			s.Add(1)
			s.Add(3)
			s.Add(2)
			s.Add(4)
			ok := s.Add(5)
			assert.True(t, ok)
			assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, utils.SliceFrom[int](s))
			it := s.Iterator()
			arr := make([]int, 5)
			for i := 0; i < 5; i++ {
				assert.True(t, it.Next())
				assert.Equal(t, i, it.Index())
				arr[i] = it.Value()
			}
			assert.False(t, it.Next())
			assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, arr)
		})
	}
}

func TestIteratorForNode(t *testing.T) {
	tests := []struct {
		name string
		b    bucket[int]
	}{
		{"slNode", newSlNode[int](3, nil)},
		{"slxNode", newSlxNode[int](3, nil, nil)},
		{"dlNode", newDlNode[int](3, nil, nil)},
		{"dlxNode", newDlxNode[int](3, nil, nil, nil)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := test.b
			saveElemIntoBucket(b, 5, eqInt, func(elem int) node[int] { return newSlNode(elem, nil) })
			it := b.Iterator()
			assert.True(t, it.Next())
			assert.Equal(t, 0, it.Index())
			assert.Equal(t, 3, it.Value())
			assert.True(t, it.Next())
			assert.Equal(t, 1, it.Index())
			assert.Equal(t, 5, it.Value())
			assert.False(t, it.Next())
			assert.Equal(t, 0, it.Value())
		})
	}
}

func Test_forEach(t *testing.T) {
	c := mocks.NewMockCollectionOf(3, 4, 5)
	arr := make([]int, 3)
	forEach[int](c, func(index int, v int) {
		arr[index] = v
	})
	assert.Equal(t, arr, utils.SliceFrom[int](c))
}

func Test_Each(t *testing.T) {
	c1 := NewArrayStack[int](3)
	c1.Push(1)
	c1.Push(2)
	c1.Push(3)
	c2 := NewArrayListOf(1, 2, 3)
	c3 := NewLinkedListOf(1, 2, 3)
	c4 := NewHeapMinPriorityQueue[int]()
	c4.Enqueue(1)
	c4.Enqueue(2)
	c4.Enqueue(3)
	c5 := NewTreeSetOf(1, 2, 3)
	c6 := NewHashSet[int](funcs.NumHash[int], funcs.ValueEqual[int])
	c6.Add(1)
	c6.Add(2)
	c6.Add(3)
	c7 := NewLinkedHashSet(funcs.NumHash[int], funcs.ValueEqual[int])
	c7.Add(1)
	c7.Add(2)
	c7.Add(3)

	tests := []struct {
		name string
		c    types.Collection[int]
	}{
		{"ArrayStack", c1},
		{"CircularArrayList", c2},
		{"LinkedList", c3},
		{"HeapMinPriorityQueue", c4},
		{"TreeSet", c5},
		{"HashTable", c6},
		{"LinkedHashTable", c7},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := test.c
			arr := make([]int, 3)
			c.Each(func(i, v int) {
				arr[i] = v
			})
			assert.Equal(t, arr, utils.SliceFrom(c))
		})
	}
}
