package transform

import (
	"fmt"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/mocks"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapTo(t *testing.T) {
	c := collections.NewArrayListOf(1, 3, 4)
	l := collections.NewLinkedListOf[string]()
	n := MapTo[int, string](c, l, func(elem int) string { return fmt.Sprint(elem) })
	assert.Equal(t, []string{"1", "3", "4"}, utils.SliceFrom[string](l))
	assert.Equal(t, 3, n)
}

func TestFilterMapTo(t *testing.T) {
	c := mocks.NewMockCollectionOf(1, 3, 4)
	l := collections.NewLinkedListOf[string]()
	n := FilterMapTo[int, string](c, l, func(elem int) bool { return elem > 2 }, func(elem int) string { return fmt.Sprint(elem) })
	assert.Equal(t, []string{"3", "4"}, utils.SliceFrom[string](l))
	assert.Equal(t, 2, n)
}

func TestFilterTo(t *testing.T) {
	c := mocks.NewMockCollectionOf(1, 3, 4)
	l := collections.NewLinkedListOf[int]()
	n := FilterTo[int](c, l, func(elem int) bool { return elem > 2 })
	assert.Equal(t, []int{3, 4}, utils.SliceFrom[int](l))
	assert.Equal(t, 2, n)
}

func TestFlatMapTo(t *testing.T) {
	c := collections.NewLinkedListOfEq(nil, []int{1, 2}, []int{3}, []int{4, 5, 6})
	l := collections.NewLinkedListOf[string]()
	n := FlatMapTo[[]int, string](c, l, func(elem []int) []string {
		ss := make([]string, len(elem))
		for i, v := range elem {
			ss[i] = fmt.Sprint(v)
		}
		return ss
	})
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, utils.SliceFrom[string](l))
	assert.Equal(t, 6, n)
}

func TestFilterFlatMapTo(t *testing.T) {
	c := collections.NewLinkedListOfEq(nil, []int{1, 2}, []int{3}, []int{4, 5, 6})
	l := collections.NewLinkedListOf[string]()
	n := FilterFlatMapTo[[]int, string](
		c, l,
		func(elem []int) bool { return len(elem) >= 2 },
		func(elem []int) []string {
			ss := make([]string, len(elem))
			for i, v := range elem {
				ss[i] = fmt.Sprint(v)
			}
			return ss
		})
	assert.Equal(t, []string{"1", "2", "4", "5", "6"}, utils.SliceFrom[string](l))
	assert.Equal(t, 5, n)
}

func TestReduce(t *testing.T) {
	c := mocks.NewMockCollectionOf(1, 3, 4)
	res := Reduce[int, string](c, "", func(acc string, next int) string { return acc + strconv.Itoa(next) })
	assert.Equal(t, "134", res)
}

func TestCount(t *testing.T) {
	c := mocks.NewMockCollectionOf(1, 3, 4)
	res := Count[int](c, func(elem int) bool {
		return elem > 2
	})
	assert.Equal(t, 2, res)
}
