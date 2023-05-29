package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedHashMap_Limit(t *testing.T) {
	t.Run("limit 3, no access order", func(t *testing.T) {
		m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 3, 0)
		m.Put(1, 100)
		m.Put(2, 200)
		m.Put(3, 300)
		m.Put(4, 400)
		m.Get(1)
		assert.Equal(t, 3, m.Size())
		assert.False(t, m.ContainsKey(1))
		assert.Equal(t, 2, m.head.Value().Key())
		assert.Equal(t, 4, m.tail.Value().Key())
	})
}

func TestLinkedHashMap_AccessOrder(t *testing.T) {
	t.Run("no access order", func(t *testing.T) {
		m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 0, 0)
		m.Put(1, 100)
		m.Put(2, 200)
		m.Put(3, 300)
		m.Put(2, 210)
		m.Get(1)
		l := make([]int, 3)
		m.linkedHashTable.Each(func(index int, e types.Entry[int, int]) {
			l[index] = e.Key()
		})
		assert.Equal(t, []int{1, 2, 3}, l)
	})

	t.Run("put access order", func(t *testing.T) {
		m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 0, PutOrder)
		m.Put(1, 100)
		m.Put(2, 200)
		m.Put(3, 300)
		m.Put(2, 210)
		m.Get(1)
		l := make([]int, 3)
		m.linkedHashTable.Each(func(index int, e types.Entry[int, int]) {
			l[index] = e.Key()
		})
		assert.Equal(t, []int{1, 3, 2}, l)

		m.Put(1, 110)
		assert.Equal(t, 3, m.head.Value().Key())
		assert.Equal(t, 1, m.tail.Value().Key())
	})

	t.Run("get access order", func(t *testing.T) {
		m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 0, GetOrder)
		m.Put(1, 100)
		m.Put(2, 200)
		m.Put(3, 300)
		m.Put(2, 210)
		m.Get(1)
		l := make([]int, 3)
		m.linkedHashTable.Each(func(index int, e types.Entry[int, int]) {
			l[index] = e.Key()
		})
		assert.Equal(t, []int{2, 3, 1}, l)
	})

	t.Run("put and get access order", func(t *testing.T) {
		m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 0, GetOrder+PutOrder)
		m.Put(1, 100)
		m.Put(2, 200)
		m.Put(3, 300)
		m.Put(2, 210)
		m.Get(1)
		l := make([]int, 3)
		m.linkedHashTable.Each(func(index int, e types.Entry[int, int]) {
			l[index] = e.Key()
		})
		assert.Equal(t, []int{3, 2, 1}, l)
	})
}

func TestLinkedHashMap_Head_Tail(t *testing.T) {
	m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 0, OriginalOrder)
	k, v, found := m.Head()
	assert.False(t, found)
	k, v, found = m.Tail()
	assert.False(t, found)
	k, v, found = m.RemoveHead()
	assert.False(t, found)
	k, v, found = m.RemoveTail()
	assert.False(t, found)

	v, found = m.PutHead(3, 33)
	assert.False(t, found)
	v, found = m.PutHead(3, 333)
	assert.True(t, found)
	assert.Equal(t, 33, v)

	v, found = m.PutTail(4, 44)
	assert.False(t, found)
	v, found = m.PutTail(4, 444)
	assert.True(t, found)
	assert.Equal(t, 44, v)

	v, found = m.PutTail(5, 55)
	v, found = m.PutHead(2, 22)

	k, v, found = m.Head()
	assert.True(t, found)
	assert.Equal(t, 2, k)
	assert.Equal(t, 22, v)
	k, v, found = m.Tail()
	assert.True(t, found)
	assert.Equal(t, 5, k)
	assert.Equal(t, 55, v)
	assert.Equal(t, 4, m.Size())

	k, v, found = m.RemoveHead()
	assert.True(t, found)
	assert.Equal(t, 2, k)
	assert.Equal(t, 22, v)
	assert.Equal(t, 3, m.Size())
	k, v, found = m.RemoveTail()
	assert.True(t, found)
	assert.Equal(t, 5, k)
	assert.Equal(t, 55, v)
	assert.Equal(t, 2, m.Size())

	m.Clear()
	m.PutTail(1, 11)
	k, v, found = m.RemoveHead()
	assert.Equal(t, 1, k)
	assert.Equal(t, 11, v)
}

func TestLinkedHashMap_Remove_After_Expanding(t *testing.T) {
	m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 30, OriginalOrder)
	for i := 1; i <= 20; i++ {
		m.Put(i, 10+i)
	}
	assert.Equal(t, 20, m.Size())
	k, v, _ := m.RemoveHead()
	assert.Equal(t, 19, m.Size())
	assert.Equal(t, 1, k)
	assert.Equal(t, 11, v)
	k, v, _ = m.RemoveHead()
	assert.Equal(t, 18, m.Size())
	assert.Equal(t, 2, k)
	assert.Equal(t, 12, v)
	k, v, _ = m.RemoveHead()
	assert.Equal(t, 17, m.Size())
	assert.Equal(t, 3, k)
	assert.Equal(t, 13, v)
	k, v, _ = m.RemoveTail()
	assert.Equal(t, 16, m.Size())
	assert.Equal(t, 20, k)
	assert.Equal(t, 30, v)
	k, v, _ = m.RemoveTail()
	assert.Equal(t, 15, m.Size())
	assert.Equal(t, 19, k)
	assert.Equal(t, 29, v)
}
