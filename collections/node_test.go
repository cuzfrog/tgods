package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
