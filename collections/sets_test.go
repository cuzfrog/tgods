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
		s    types.SortedSet[int]
	}{
		{"treeSet1", NewTreeSetOf[int]()},
		{"treeSet2", NewTreeSetOfComp[int](funcs.ValueCompare[int])},
		{"enumSet", NewEnumSet[int](10)},
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
			s.Remove(5)
			assert.Equal(t, 2, s.Size())
			assert.Equal(t, []int{2, 9}, utils.SliceFrom[int](s))

			v, found := s.RemoveFirst()
			assert.True(t, found)
			assert.Equal(t, 2, v)
			v, found = s.RemoveFirst()
			assert.Equal(t, 9, v)
			v, found = s.RemoveFirst()
			assert.False(t, found)
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
		{"linkedHashSet1", NewLinkedHashSetOfNum[int](2)},
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

func TestSetOfStrProperties(t *testing.T) {
	tests := []struct {
		name string
		s    types.Set[string]
	}{
		{"hashSet1", NewHashSetOfStr("1", "2", "3", "2")},
		{"linkedHashSet1", NewLinkedHashSetOfStr("1", "2", "3", "2")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.s
			assert.Equal(t, 3, s.Size())
			s.Remove("3")
			assert.ElementsMatch(t, []string{"1", "2"}, utils.SliceFrom[string](s))
		})
	}

}
