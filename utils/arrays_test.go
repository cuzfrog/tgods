package utils

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/mocks"
	"github.com/cuzfrog/tgods/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type myInt int

func (i myInt) Less(other myInt) bool {
	return i < other
}

func TestSlice(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](1, 4, 3, 2)
	arr := SliceFrom[int](c)
	assert.Equal(t, []int{1, 4, 3, 2}, arr)
}

func TestAddSliceTo(t *testing.T) {
	c := mocks.NewMockCollection[int](3)
	AddSliceTo[int]([]int{1, 2, 3}, c)
	assert.Equal(t, []int{1, 2, 3}, c.Elems())

}

func TestSliceProjection(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](1, 4, 3, 2)
	arr := SliceProject[int, int](c, func(v int) int { return v + 1 })
	assert.Equal(t, []int{2, 5, 4, 3}, arr)
}

func TestSort(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](1, 4, 3, 2)
	less := funcs.ValueLess[int]
	Sort[int](c, less)
	assert.ObjectsAreEqual(less, c.GetFlag("SortLessFn"))
}

func TestSortC(t *testing.T) {
	c := mocks.NewMockCollectionOf[myInt](1, 4, 3, 2)
	SortC[myInt](c)
	var less types.Less[myInt]
	less = c.GetFlag("SortLessFn").(types.Less[myInt])

	assert.True(t, less(1, 2))
	assert.False(t, less(1, 1))
	assert.False(t, less(2, 1))
}

func TestSortOrderable(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](1, 4, 3, 2)
	SortOrderable[int](c)
	assert.ObjectsAreEqual(funcs.ValueLess[int], c.GetFlag("SortLessFn"))
}
