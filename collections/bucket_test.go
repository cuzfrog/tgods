package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Bucket(t *testing.T) {
	tests := []struct {
		name      string
		l         bucket[int]
		newNodeOf func(v int) node[int]
	}{
		{"slNode", newSlNode(3, nil), func(elem int) node[int] { return newSlNode(elem, nil) }},
		{"slxNode", newSlxNode(3, nil, nil), func(elem int) node[int] { return newSlxNode(elem, nil, nil) }},
		{"dlNode", newDlNode(3, nil, nil), func(elem int) node[int] { return newDlNode(elem, nil, nil) }},
		{"dlxNode", newDlxNode(3, nil, nil, nil), func(elem int) node[int] { return newDlxNode(elem, nil, nil, nil) }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := test.l
			b, n, v, found := saveElemIntoBucket(l, 3, eqInt, test.newNodeOf)
			assert.True(t, found)
			assert.Same(t, n, b)
			assert.Equal(t, 3, v)
			assert.Equal(t, 3, l.Value())
			assert.Nil(t, l.Next())
			b, v, found = removeElemFromBucket[int](l, 3, eqInt)
			assert.True(t, found)
			assert.Equal(t, 3, v)
			assert.Nil(t, b)
			b, v, found = removeElemFromBucket[int](l, 4, eqInt)
			assert.False(t, found)
			assert.Equal(t, l, b)

			saveElemIntoBucket(l, 4, eqInt, test.newNodeOf)
			assert.Equal(t, 4, l.Next().Value())
			b, n, v, found = saveElemIntoBucket(l, 5, eqInt, test.newNodeOf)
			assert.Same(t, l, b)
			assert.False(t, found)
			assert.Equal(t, 5, l.Next().Next().Value())
			assert.Equal(t, 5, n.Value())

			assert.True(t, l.Contains(3, eqInt))
			assert.False(t, l.Contains(8, eqInt))

			v, found = l.Get(4, eqInt)
			assert.True(t, found)
			assert.Equal(t, 4, v)
			v, found = l.Get(7, eqInt)
			assert.False(t, found)

			b, v, found = removeElemFromBucket[int](l, 4, eqInt)
			assert.True(t, found)
			assert.Equal(t, 4, v)
			assert.Equal(t, l, b)
			assert.Equal(t, 3, l.Value())
			assert.Equal(t, 5, l.Next().Value())
			assert.Nil(t, l.Next().Next())
		})
	}
}
