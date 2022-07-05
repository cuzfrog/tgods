package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularArrayList_shrinkIfNeeded(t *testing.T) {
	l := CircularArrayList[int]{5, 8, make([]int, 200), 3}
	l.arr[l.start] = 2
	l.arr[l.end-1] = 4
	l.shrinkIfNeeded()
	assert.Equal(t, 25, len(l.arr))
	assert.Equal(t, 0, l.start)
	assert.Equal(t, 3, l.end)
	assert.Equal(t, 3, l.size)
	assert.ElementsMatch(t, []int{2, 0, 4}, l.arr[0:3])

	l = CircularArrayList[int]{5, 8, make([]int, 12), 3}
	l.shrinkIfNeeded()
	assert.Equal(t, 12, len(l.arr))

	l = CircularArrayList[int]{5, 8, make([]int, 14), 3}
	l.shrinkIfNeeded()
	assert.Equal(t, 14, len(l.arr))
}

func TestCircularArrayList_expandIfNeeded(t *testing.T) {
	l := NewCircularArrayList(1, 3, 5)
	l.PopHead()
	l.Add(6)
	l.expandIfNeeded()
	assert.Equal(t, 6, len(l.arr))
	assert.ElementsMatch(t, []int{3, 5, 6}, l.arr[0:3])
	assert.Equal(t, 3, l.size)
	assert.Equal(t, 0, l.start)
	assert.Equal(t, 3, l.end)

	l = NewCircularArrayListWithInitSize[int](5)
	l.Add(4)
	l.AddHead(2)
	l.expandIfNeeded()
	assert.Equal(t, 5, len(l.arr))
	assert.Equal(t, 2, l.size)
	assert.Equal(t, 4, l.start)
	assert.Equal(t, 1, l.end)
}

func TestCircularArrayList_Add(t *testing.T) {
	l := NewCircularArrayList[int]()
	l.Add(3)
	assert.Equal(t, 0, l.start)
	assert.Equal(t, 1, l.end)
	assert.Equal(t, 1, l.size)
	l.AddHead(7)
	l.Add(6)
	assert.Equal(t, 3, l.size)
	assert.Equal(t, len(l.arr)-1, l.start)
	assert.Equal(t, 2, l.end)
	assert.Equal(t, 7, l.arr[l.start])
	assert.Equal(t, 3, l.arr[0])
	assert.Equal(t, 6, l.arr[1])
}

func TestCircularArrayList_AddHead(t *testing.T) {
	l := NewCircularArrayListWithInitSize[int](5)
	l.AddHead(3)
	assert.Equal(t, 1, l.size)
	assert.Equal(t, 4, l.start)
	assert.Equal(t, 0, l.end)
	assert.Equal(t, 3, l.arr[4])
	l.AddHead(5)
	assert.Equal(t, 2, l.size)
	assert.Equal(t, 3, l.start)
	assert.Equal(t, 0, l.end)
	assert.Equal(t, 5, l.arr[3])
	l.Add(33)
	assert.Equal(t, 3, l.size)
	assert.Equal(t, 3, l.start)
	assert.Equal(t, 1, l.end)
	assert.Equal(t, 33, l.arr[0])
}

func TestCircularArrayList_Clear(t *testing.T) {
	l := NewCircularArrayListWithInitSize[int](5)
	l.AddHead(5)
	l.Clear()
	assert.Equal(t, 0, l.size)
	assert.Equal(t, -1, l.start)
	assert.Equal(t, 0, l.end)
	assert.Nil(t, l.arr)
}

func TestCircularArrayList_Contains(t *testing.T) {
	l := NewCircularArrayListWithInitSize[int](5)
	l.AddHead(5)
	assert.True(t, l.Contains(5))
	assert.False(t, l.Contains(6))
	l.Clear()
	assert.False(t, l.Contains(5))
}

func TestCircularArrayList_Get(t *testing.T) {
	l := NewCircularArrayListWithInitSize[int](5)
	l.Add(100)
	for i := 1; i <= 12; i++ {
		l.AddHead(i)
	}
	v, ok := l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 12, v)
	v, ok = l.Get(12)
	assert.Equal(t, 100, v)
	v, ok = l.Get(13)
	assert.False(t, ok)
}

func TestCircularArrayList_Set(t *testing.T) {
	l := NewCircularArrayListWithInitSize[int](5)
	l.Add(100)
	v, ok := l.Set(0, 5)
	assert.True(t, ok)
	assert.Equal(t, 100, v)
	assert.Equal(t, 5, l.arr[0])

	v, ok = l.Set(1, 35)
	assert.False(t, ok)
}

func TestCircularArrayList_Peek(t *testing.T) {
	l := NewCircularArrayList(3, 5)
	v, ok := l.Peek()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	v, ok = l.Tail()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	assert.Equal(t, 2, l.size)

	l.Clear()
	v, ok = l.Tail()
	assert.False(t, ok)
}

func TestCircularArrayList_Pop(t *testing.T) {
	l := NewCircularArrayList(3, 5)
	v, ok := l.Pop()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	v, ok = l.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = l.Pop()
	assert.False(t, ok)
}

func TestCircularArrayList_PopHead(t *testing.T) {
	l := NewCircularArrayList(3, 5)
	v, ok := l.PopHead()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = l.PopHead()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	v, ok = l.PopHead()
	assert.False(t, ok)
}

func TestCircularArrayList_Iterator(t *testing.T) {
	l := NewCircularArrayList(3, 5, 7)
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
