package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCircularArrayListOfType(t *testing.T) {
	l := newCircularArrayOfEq[*intStruct](3, func(a, b *intStruct) bool { return a.v > b.v })
	assert.False(t, l.Contains(&intStruct{2}))
	l.Add(&intStruct{3})
	assert.True(t, l.Contains(&intStruct{2}))
	assert.False(t, l.Contains(&intStruct{5}))
}

func TestCircularArrayList_shrinkIfNeeded(t *testing.T) {
	l := &circularArray[int]{5, 8, make([]int, 200), 3, equalInt, list}
	l.arr[l.start] = 2
	l.arr[l.end-1] = 4
	l.shrinkIfNeeded()
	assert.Equal(t, 25, len(l.arr))
	assert.Equal(t, 0, l.start)
	assert.Equal(t, 3, l.end)
	assert.Equal(t, 3, l.Size())
	assert.Equal(t, []int{2, 0, 4}, l.arr[0:3])

	l = &circularArray[int]{5, 8, make([]int, 12), 3, equalInt, list}
	l.shrinkIfNeeded()
	assert.Equal(t, 12, len(l.arr))

	l = &circularArray[int]{5, 8, make([]int, 14), 3, equalInt, list}
	l.shrinkIfNeeded()
	assert.Equal(t, 14, len(l.arr))

	l = newCircularArrayOf[int]()
	l.shrinkIfNeeded()
}

func TestCircularArrayList_expandIfNeeded(t *testing.T) {
	l := newCircularArrayOf(1, 3, 5)
	l.RemoveHead()
	l.Add(6)
	l.expandIfNeeded()
	assert.Equal(t, 6, len(l.arr))
	assert.Equal(t, []int{3, 5, 6}, l.arr[0:3])
	assert.Equal(t, 3, l.size)
	assert.Equal(t, 0, l.start)
	assert.Equal(t, 3, l.end)

	l = newCircularArray[int](5)
	l.Add(4)
	l.AddHead(2)
	l.expandIfNeeded()
	assert.Equal(t, 5, len(l.arr))
	assert.Equal(t, 2, l.size)
	assert.Equal(t, 4, l.start)
	assert.Equal(t, 1, l.end)
}

func TestCircularArrayList_Add(t *testing.T) {
	l := newCircularArrayOf[int]()
	l.Add(3)
	assert.Equal(t, 0, l.start)
	assert.Equal(t, 1, l.end)
	assert.Equal(t, 1, l.size)
	l.AddHead(7)
	l.AddTail(6)
	assert.Equal(t, 3, l.size)
	assert.Equal(t, len(l.arr)-1, l.start)
	assert.Equal(t, 2, l.end)
	assert.Equal(t, 7, l.arr[l.start])
	assert.Equal(t, 3, l.arr[0])
	assert.Equal(t, 6, l.arr[1])
}

func TestCircularArrayList_AddHead(t *testing.T) {
	l := newCircularArray[int](5)
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
	l := newCircularArray[int](5)
	l.AddHead(5)
	l.Clear()
	assert.Equal(t, 0, l.size)
	assert.Equal(t, -1, l.start)
	assert.Equal(t, 0, l.end)
	assert.Nil(t, l.arr)
}

func TestCircularArrayList_Contains(t *testing.T) {
	l := newCircularArray[int](5)
	l.AddHead(5)
	assert.True(t, l.Contains(5))
	assert.False(t, l.Contains(6))
	l.Clear()
	assert.False(t, l.Contains(5))
}

func TestCircularArrayList_Get(t *testing.T) {
	l := newCircularArray[int](5)
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
	l := newCircularArray[int](5)
	l.Add(100)
	v, ok := l.Set(0, 5)
	assert.True(t, ok)
	assert.Equal(t, 100, v)
	assert.Equal(t, 5, l.arr[0])

	v, ok = l.Set(1, 35)
	assert.False(t, ok)
}

func TestCircularArrayList_Swap(t *testing.T) {
	l := newCircularArrayOf(3, 6, 7, 8)
	ok := l.Swap(0, 2)
	assert.True(t, ok)
	assert.Equal(t, []int{7, 6, 3, 8}, l.arr)
	ok = l.Swap(0, 7)
	assert.False(t, ok)
	ok = l.Swap(9, 1)
	assert.False(t, ok)
}

func TestCircularArrayList_Peek(t *testing.T) {
	l := newCircularArrayOf(3, 5)
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

func TestCircularArrayList_Head(t *testing.T) {
	l := newCircularArrayOf(3, 5)
	v, ok := l.Head()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = l.First()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	l = newCircularArrayOf[int]()
	v, ok = l.Head()
	assert.False(t, ok)
}

func TestCircularArrayList_Remove(t *testing.T) {
	l := newCircularArrayOf(3, 5)
	v, ok := l.RemoveTail()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	v, ok = l.RemoveTail()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = l.RemoveTail()
	assert.False(t, ok)

	l = &circularArray[int]{4, 1, make([]int, 6), 3, equalInt, list}
	l.arr[0] = 13
	l.arr[4] = 3
	l.arr[5] = 11
	v, ok = l.RemoveTail()
	assert.True(t, ok)
	assert.Equal(t, 13, v)
	assert.Equal(t, 2, l.size)
	v, ok = l.RemoveTail()
	assert.Equal(t, 11, v)
	assert.Equal(t, 5, l.end)
	assert.Equal(t, 1, l.size)
	v, ok = l.RemoveTail()
	assert.Equal(t, 3, v)
	assert.Equal(t, 4, l.end)
	assert.Equal(t, 0, l.size)
	v, ok = l.RemoveTail()
	assert.False(t, ok)
}

func TestCircularArrayList_RemoveHead(t *testing.T) {
	l := newCircularArrayOf(3, 5)
	v, ok := l.RemoveHead()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = l.RemoveHead()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	v, ok = l.RemoveHead()
	assert.False(t, ok)
}

func TestCircularArrayList_Clone(t *testing.T) {
	l := newCircularArrayOf(3, 5, 7)
	nl := l.clone()
	assert.NotSame(t, l, nl)
	assert.Equal(t, utils.SliceFrom[int](l), utils.SliceFrom[int](nl))
}

func TestCircularArray_toArrIndex(t *testing.T) {
	l := newCircularArrayOf(3, 5)
	l.AddHead(7)
	l.Add(13)
	i, ok := l.toArrIndex(0)
	assert.True(t, ok)
	assert.Equal(t, len(l.arr)-1, i)
	i, ok = l.toArrIndex(1)
	assert.Equal(t, 0, i)
	i, ok = l.toArrIndex(2)
	assert.Equal(t, 1, i)
	i, ok = l.toArrIndex(3)
	assert.Equal(t, 2, i)
	i, ok = l.toArrIndex(l.Size() - 1)
	assert.True(t, ok)
	assert.Equal(t, 2, i)

	i, ok = l.toArrIndex(-1)
	assert.False(t, ok)
	i, ok = l.toArrIndex(l.Size())
	assert.False(t, ok)
	i, ok = l.toArrIndex(4)
	assert.False(t, ok)
}
