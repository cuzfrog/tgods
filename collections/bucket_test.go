package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListBucket(t *testing.T) {
	var l *slNode[int]
	b, v, found := l.Save(1, eqInt)
	assert.False(t, found)
	assert.Equal(t, &slNode[int]{1, nil}, b)
	b, v, found = l.Delete(1, eqInt)
	assert.False(t, found)
	assert.Nil(t, b)

	l = &slNode[int]{3, nil}
	b, v, found = l.Save(3, eqInt)
	assert.True(t, found)
	assert.Equal(t, 3, v)
	assert.Equal(t, 3, l.v)
	assert.Nil(t, l.next)
	b, v, found = l.Delete(3, eqInt)
	assert.True(t, found)
	assert.Equal(t, 3, v)
	assert.Nil(t, b)
	b, v, found = l.Delete(4, eqInt)
	assert.False(t, found)
	assert.Equal(t, l, b)

	l.Save(4, eqInt)
	assert.Equal(t, 4, l.next.Value())
	b, v, found = l.Save(5, eqInt)
	assert.False(t, found)
	assert.Equal(t, 5, l.next.Next().Value())

	assert.True(t, l.Contains(3, eqInt))
	assert.False(t, l.Contains(8, eqInt))

	v, found = l.Get(4, eqInt)
	assert.True(t, found)
	assert.Equal(t, 4, v)
	v, found = l.Get(7, eqInt)
	assert.False(t, found)

	b, v, found = l.Delete(4, eqInt)
	assert.True(t, found)
	assert.Equal(t, 4, v)
	assert.Equal(t, l, b)
	assert.Equal(t, 3, l.v)
	assert.Equal(t, 5, l.next.Value())
	assert.Nil(t, l.next.Next())
}
