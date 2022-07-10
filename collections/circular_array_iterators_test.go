package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularArrayList_Iterator(t *testing.T) {
	l := newCircularArrayOf(3, 5, 7)
	it := l.Iterator()
	assert.True(t, it.Next())
	assert.Equal(t, 0, it.Index())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Index())
	assert.Equal(t, 5, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Index())
	assert.Equal(t, 7, it.Value())
	assert.False(t, it.Next())
	assert.False(t, it.Next())

	l.Clear()
	it = l.Iterator()
	assert.False(t, it.Next())
}
