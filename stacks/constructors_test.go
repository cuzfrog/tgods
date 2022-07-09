package stacks

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/cuzfrog/tgods/utils"
	"testing"
)

func TestStackProperties(t *testing.T) {
	tests := []struct {
		name string
		s    core.Stack[int]
	}{
		{"arrayStack", NewArrayStack[int](5)},
		{"linkedList", NewLinkedStack[int]()},
		{"circularArray", NewCircularArrayStack[int]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.s.Enstack(1)
			test.s.Enstack(3)
			test.s.Enstack(4)
			utils.SliceFrom[int](test.s)

		})
	}
}
