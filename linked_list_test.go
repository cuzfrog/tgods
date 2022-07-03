package generic_collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(5)
	assert.Equal(t, l.Size(), 2)
	assert.Equal(t, l.head, 1)
	assert.Equal(t, l.tail, 5)
}
