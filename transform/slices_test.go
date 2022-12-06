package transform

import (
	"fmt"
	"github.com/cuzfrog/tgods/collections"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapSliceTo(t *testing.T) {
	c := []int{1, 3, 4}
	l := collections.NewLinkedListOf[string]()
	n := MapSliceTo[int, string](c, l, func(elem int) string { return fmt.Sprint(elem) })
	assert.Equal(t, []string{"1", "3", "4"}, utils.SliceFrom[string](l))
	assert.Equal(t, 3, n)
}

func TestFilterMapSliceTo(t *testing.T) {
	c := []int{1, 3, 4}
	l := collections.NewLinkedListOf[string]()
	n := FilterMapSliceTo[int, string](c, l, func(elem int) bool { return elem > 2 }, func(elem int) string { return fmt.Sprint(elem) })
	assert.Equal(t, []string{"3", "4"}, utils.SliceFrom[string](l))
	assert.Equal(t, 2, n)
}

func TestFilterSliceTo(t *testing.T) {
	c := []int{1, 3, 4}
	l := collections.NewLinkedListOf[int]()
	n := FilterSliceTo[int](c, l, func(elem int) bool { return elem > 2 })
	assert.Equal(t, []int{3, 4}, utils.SliceFrom[int](l))
	assert.Equal(t, 2, n)
}

func stringToChar(elem string) []uint8 {
	ss := make([]uint8, len(elem))
	for i := 0; i < len(elem); i++ {
		ss[i] = elem[i]
	}
	return ss
}

func TestFlatMapSliceTo(t *testing.T) {
	c := []string{"12", "3", "456"}
	l := collections.NewLinkedListOf[uint8]()
	n := FlatMapSliceTo[string, uint8](c, l, stringToChar)
	assert.Equal(t, []uint8{'1', '2', '3', '4', '5', '6'}, utils.SliceFrom[uint8](l))
	assert.Equal(t, 6, n)
}

func TestFilterFlatMapSliceTo(t *testing.T) {
	c := []string{"12", "3", "456"}
	l := collections.NewLinkedListOf[uint8]()
	n := FilterFlatMapSliceTo[string, uint8](
		c, l,
		func(elem string) bool { return len(elem) >= 2 },
		stringToChar)
	assert.Equal(t, []uint8{'1', '2', '4', '5', '6'}, utils.SliceFrom[uint8](l))
	assert.Equal(t, 5, n)
}

func TestReduceSlice(t *testing.T) {
	c := []int{1, 3, 4}
	res := ReduceSlice[int, string](c, "", func(acc string, next int) string { return acc + strconv.Itoa(next) })
	assert.Equal(t, "134", res)
}

func TestCountSlice(t *testing.T) {
	c := []int{1, 3, 4}
	res := CountSlice[int](c, func(elem int) bool {
		return elem > 2
	})
	assert.Equal(t, 2, res)
}
