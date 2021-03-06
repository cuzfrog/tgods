package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSetProperties(t *testing.T) {
	tests := []struct {
		name string
		s    types.Set[int]
	}{
		{"treeSet1", NewTreeSetOf[int]()},
		{"treeSet2", NewTreeSetOfComp[int](funcs.ValueCompare[int])},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.s
			s.Add(2)
			s.Add(2)
			assert.Equal(t, 1, s.Size())
			s.Add(3)
			s.Add(9)
			s.Add(5)
			s.Remove(3)
			assert.Equal(t, []int{2, 5, 9}, utils.SliceFrom[int](s))
		})
	}
}

func TestSetProperties(t *testing.T) {
	tests := []struct {
		name string
		s    types.Set[int]
	}{
		{"treeSet1", NewTreeSetOf[int]()},
		{"treeSet2", NewTreeSetOfComp[int](funcs.ValueCompare[int])},
		{"hashSet1", NewHashSetOfNum[int](2)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.s
			s.Add(2)
			s.Add(2)
			assert.Equal(t, 1, s.Size())
			s.Add(3)
			s.Add(9)
			s.Add(5)
			s.Remove(3)
			assert.ElementsMatch(t, []int{2, 5, 9}, utils.SliceFrom[int](s))
		})
	}
}

func TestSetStr(t *testing.T) {
	s := NewHashSetOfStr("1", "2", "3", "2")
	assert.Equal(t, 3, s.Size())
	s.Remove("3")
	assert.ElementsMatch(t, []string{"1", "2"}, utils.SliceFrom[string](s))
}
