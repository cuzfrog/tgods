package collections

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type obj struct {
	v int
	s int
}

func TestLinkedList_New(t *testing.T) {
	comp := func(a string, b string) bool { return strings.Contains(a, b) }
	l := newLinkedListOfEq(comp)
	l.Add("abc")
	assert.True(t, l.Contains("ab"))
	assert.False(t, l.Contains("bcd"))
}

func TestLinkedList_Clear(t *testing.T) {
	l := newLinkedListOf("a", "d")
	l.Clear()
	assert.Equal(t, 0, l.size)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestLinkedList_Contains(t *testing.T) {
	l := newLinkedListOf("a", "d")
	assert.True(t, l.Contains("d"))
	assert.False(t, l.Contains("ds"))
}

func TestLinkedList_Head_Tail(t *testing.T) {
	l := newLinkedListOf[int]()
	_, okh := l.Head()
	assert.False(t, okh)
	_, okt := l.Tail()
	assert.False(t, okt)

	l.Add(3)
	l.Add(5)
	v, okh2 := l.Head()
	assert.True(t, okh2)
	assert.Equal(t, 3, v)
	assert.Equal(t, 2, l.size)
	v, okh2 = l.First()
	assert.True(t, okh2)
	assert.Equal(t, 3, v)
	v, okt2 := l.Tail()
	assert.True(t, okt2)
	assert.Equal(t, 5, v)
	vp, okp := l.Peek()
	assert.Equal(t, v, vp)
	assert.Equal(t, okt2, okp)
}

func TestLinkedList_Add(t *testing.T) {
	l := newLinkedListOf[int]()
	l.Add(1)
	l.AddTail(5)
	assert.Equal(t, 2, l.Size())
	assert.Equal(t, 1, l.head.v)
	assert.Equal(t, 5, l.tail.v)
	assert.Equal(t, l.head.next, l.tail)
	assert.Equal(t, l.tail.prev, l.head)
	assert.Nil(t, l.head.prev)
	assert.Nil(t, l.tail.next)

	l2 := newLinkedListOf[*obj]()
	l2.Add(nil)
	assert.Nil(t, l2.head.v)
}

func TestLinkedList_AddHead(t *testing.T) {
	l := newLinkedListOf[int]()
	l.AddHead(5)
	l.Add(1)
	assert.Equal(t, 2, l.Size())
	assert.Equal(t, 5, l.head.v)
	assert.Equal(t, 1, l.tail.v)
	assert.Equal(t, l.head.next, l.tail)
	assert.Equal(t, l.tail.prev, l.head)
	assert.Nil(t, l.head.prev)
	assert.Nil(t, l.tail.next)

	l.AddHead(6)
	assert.Equal(t, 3, l.Size())
	assert.Equal(t, 6, l.head.v)
	assert.Equal(t, 5, l.head.next.v)
	assert.Nil(t, l.head.prev)
}

func TestLinkedList_Remove(t *testing.T) {
	l := newLinkedListOf[int]()
	l.Add(1)
	l.Add(5)
	v1, _ := l.Remove()
	assert.Equal(t, 5, v1)
	assert.Nil(t, l.tail.next)
	v2, _ := l.RemoveTail()
	assert.Equal(t, 1, v2)
	assert.Equal(t, 0, l.Size())
	_, ok := l.Remove()
	assert.False(t, ok)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	l2 := newLinkedListOf[*obj]()
	v3, _ := l2.Remove()
	assert.Nil(t, v3)
}

func TestLinkedList_PopHead(t *testing.T) {
	l := newLinkedListOf[int]()
	l.Add(1)
	l.Add(5)
	v, found := l.RemoveHead()
	assert.True(t, found)
	assert.Equal(t, 1, v)
	assert.Nil(t, l.head.prev)
	v, found = l.RemoveHead()
	assert.True(t, found)
	assert.Equal(t, 5, v)
	_, found = l.RemoveHead()
	assert.False(t, found)
}
