package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackProperties(t *testing.T) {
	tests := []struct {
		name string
		s    types.Stack[int]
	}{
		{"arrayStack", NewArrayStack[int](5)},
		{"linkedList", NewLinkedStack[int]()},
		{"circularArray", NewCircularArrayStack[int]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.s
			s.Push(1)
			s.Push(3)
			s.Push(4)
			assert.Equal(t, []int{4, 3, 1}, utils.SliceFrom[int](s))
			s.Pop()
			assert.Equal(t, []int{3, 1}, utils.SliceFrom[int](s))
			s.Clear()
			assert.Empty(t, utils.SliceFrom[int](s))
		})
	}
}
