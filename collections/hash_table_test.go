package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable(t *testing.T) {
	h := newHashTable[int](funcs.NumHash[int], funcs.ValueEqual[int])
	assert.False(t, h.Contains(7))
	h.Add(1)
	assert.False(t, h.Contains(7))
	h.Add(2)
	h.Add(3)
	h.Add(4)
	assert.True(t, h.Add(5))
	assert.Equal(t, 5, h.size)
	assert.True(t, h.Contains(5))
	assert.False(t, h.Contains(7))
	assert.False(t, h.Remove(8))
	assert.True(t, h.Remove(5))
	assert.False(t, h.Contains(5))
	assert.Equal(t, 4, h.size)
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, utils.SliceFrom[int](h))
	h.Clear()
	assert.Equal(t, 0, h.size)
	assert.Nil(t, h.arr)
}

func Test_newHashTableOfInitCap(t *testing.T) {
	h := newHashTableOfInitCap(10, funcs.NumHash[int], funcs.ValueEqual[int])
	assert.Equal(t, 10, cap(h.arr))
}

var newSlNodeOf = func(v int) node[int] { return newSlBucketOf[int](v) }

func TestHashTable_expand(t *testing.T) {
	a := make([]bucket[int], 5)
	a[0] = newSlBucketOf(3)
	a[1] = newSlBucketOf(5)
	a[2] = newSlBucketOf(6)
	a[3] = newSlBucketOf(7)
	a[4] = newSlBucketOf(1)
	h := &hashTable[int]{a, 5, funcs.NumHash[int], funcs.ValueEqual[int], newSlNodeOf}
	h.expandIfNeeded()
	assert.Equal(t, 10, cap(h.arr))
	assert.ElementsMatch(t, []int{1, 3, 5, 6, 7}, utils.SliceFrom[int](h))
}

func TestHashTable_shrink(t *testing.T) {
	a := make([]bucket[int], 32)
	a[0] = newSlBucketOf(3)
	a[1] = newSlBucketOf(5)
	h := &hashTable[int]{a, 2, funcs.NumHash[int], funcs.ValueEqual[int], newSlNodeOf}
	h.shrinkIfNeeded()
	assert.Equal(t, 16, cap(h.arr))
	assert.ElementsMatch(t, []int{3, 5}, utils.SliceFrom[int](h))
}
