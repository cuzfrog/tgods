package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListProperties(t *testing.T) {
	tests := []struct {
		name string
		l    types.List[int]
	}{
		{"circularArrayList1", NewArrayListOf[int]()},
		{"circularArrayList2", NewArrayListOfSize[int](0)},
		{"circularArrayList3", NewArrayListOfEq[int](0, funcs.ValueEqual[int])},
		{"circularArrayList4", NewArrayListOfSizeP[int](0, AutoExpand)},
		{"circularArrayList5", NewArrayListOfEqP[int](0, funcs.ValueEqual[int], AutoExpand)},
		{"linkedList1", NewLinkedListOf[int]()},
		{"linkedList2", NewLinkedListOfEq[int](funcs.ValueEqual[int])},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
			l.Add(1)
			l.Add(5)
			l.Add(3)
			l.Add(3)
			l.Remove()
			assert.Equal(t, []int{1, 5, 3}, utils.SliceFrom[int](l))
			l.AddHead(6)
			l.AddHead(7)
			assert.Equal(t, []int{7, 6, 1, 5, 3}, utils.SliceFrom[int](l))
			l.RemoveHead()
			assert.Equal(t, []int{6, 1, 5, 3}, utils.SliceFrom[int](l))
			l.Clear()
			assert.Empty(t, utils.SliceFrom[int](l))
		})
	}
}

func TestListConstraintTypeConstructors(t *testing.T) {
	tests := []struct {
		name string
		q    types.List[*intStruct]
	}{
		{"arrayList1", NewArrayListOfC[*intStruct]()},
		{"arrayList2", NewArrayListOfSizeC[*intStruct](10)},
		{"arrayList3", NewArrayListOfSizePC[*intStruct](10, NoAutoSizing)},
		{"linkedList1", NewLinkedListOfC[*intStruct]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := test.q
			v1 := &intStruct{1}
			q.Add(v1)
			assert.True(t, q.Contains(v1))
		})
	}
}
