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
		m.h.Each(func(index int, e types.Entry[int, int]) {
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
		m.h.Each(func(index int, e types.Entry[int, int]) {
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
		m.h.Each(func(index int, e types.Entry[int, int]) {
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
		m.h.Each(func(index int, e types.Entry[int, int]) {
			l[index] = e.Key()
		})
		assert.Equal(t, []int{3, 2, 1}, l)
	})
}
