package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type obj struct {
	v int
	s int
}

func TestLinkedList_Add(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(5)
	assert.Equal(t, 2, l.Size())
	assert.Equal(t, 1, l.head.v)
	assert.Equal(t, 5, l.tail.v)

	l2 := NewLinkedList[*obj]()
	l2.Add(nil)
	assert.Nil(t, l2.head.v)
}

func TestLinkedList_Pop(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(5)
	v1, _ := l.Pop()
	assert.Equal(t, 5, v1)
	v2, _ := l.Pop()
	assert.Equal(t, 1, v2)
	assert.Equal(t, 0, l.Size())
	_, err := l.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	l2 := NewLinkedList[*obj]()
	v3, _ := l2.Pop()
	assert.Nil(t, v3)
}
