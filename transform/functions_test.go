package transform

import (
	"fmt"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/mocks"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapTo(t *testing.T) {
	c := collections.NewArrayList(1, 3, 4)
	l := collections.NewLinkedList[string]()
	n := MapTo[int, string](c, l, func(elem int) string { return fmt.Sprint(elem) })
	assert.Equal(t, []string{"1", "3", "4"}, utils.SliceFrom[string](l))
	assert.Equal(t, 3, n)
}

func TestFilterMapTo(t *testing.T) {
	c := mocks.NewMockCollectionOf(1, 3, 4)
	l := collections.NewLinkedList[string]()
	n := FilterMapTo[int, string](c, l, func(elem int) bool { return elem > 2 }, func(elem int) string { return fmt.Sprint(elem) })
	assert.Equal(t, []string{"3", "4"}, utils.SliceFrom[string](l))
	assert.Equal(t, 2, n)
}

func TestFilterTo(t *testing.T) {
	c := mocks.NewMockCollectionOf(1, 3, 4)
	l := collections.NewLinkedList[int]()
	n := FilterTo[int](c, l, func(elem int) bool { return elem > 2 })
	assert.Equal(t, []int{3, 4}, utils.SliceFrom[int](l))
	assert.Equal(t, 2, n)
}

func TestFlatMapTo(t *testing.T) {
	c := collections.NewLinkedListOfEq(nil, []int{1, 2}, []int{3}, []int{4, 5, 6})
	l := collections.NewLinkedList[string]()
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
	l := collections.NewLinkedList[string]()
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
