package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Node(t *testing.T) {
	tests := []struct {
		name        string
		n           node[int]
		supportPrev bool
		supportX    bool
	}{
		{"slNode", newSlNode(1, nil), false, false},
		{"slxNode", newSlxNode(1, nil, nil), false, true},
		{"dlNode", newDlNode(1, nil, nil), true, false},
		{"dlxNode", newDlxNode(1, nil, nil, nil), true, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := test.n
			assert.Equal(t, 1, n.Value())
			n.SetValue(5)
			assert.Equal(t, 5, n.Value())
			assert.Nil(t, n.Prev())
			assert.Nil(t, n.Next())

			next := newSlNode(0, nil)
			n.SetNext(next)
			assert.Same(t, next, n.Next())

			prev := newSlNode(0, nil)
			n.SetPrev(prev)
			if test.supportPrev {
				assert.Same(t, prev, n.Prev())
			} else {
				assert.Nil(t, n.Prev())
			}

			x := newSlNode(3, nil)
			if test.supportX {
				n.SetExternal(x)
				assert.Same(t, x, n.External())
			} else {
				assert.Nil(t, n.External())
				assert.Panics(t, func() {
					n.SetExternal(x)
				})
			}
		})
	}
}

func Test_removeNodeFromList(t *testing.T) {
	n1 := newDlNode(2, nil, nil)
	n2 := newDlNode(6, n1, nil)
	n1.SetNext(n2)
	n3 := newDlNode(3, n2, nil)
	n2.SetNext(n3)
	removeNodeFromList(n2)
	assert.Equal(t, n1, n3.Prev())
	assert.Equal(t, n3, n1.Next())
	assert.Nil(t, n1.Prev())
	assert.Nil(t, n3.Next())
	assert.Nil(t, n2.Prev())
	assert.Nil(t, n2.Next())
}
