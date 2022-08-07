package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumSet_Props(t *testing.T) {
	s := newEnumSet[myEnum](myEnumC)
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.Add(myEnumA))
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Contains(myEnumA))
	assert.False(t, s.Contains(myEnumC))

	assert.False(t, s.Add(myEnumA))
	assert.False(t, s.Remove(myEnumC))
	assert.True(t, s.Remove(myEnumA))

	v, found := s.First()
	assert.False(t, found)
	v, found = s.Last()
	assert.False(t, found)
	v, found = s.RemoveFirst()
	assert.False(t, found)
	v, found = s.RemoveLast()
	assert.False(t, found)

	s.Add(myEnumB)
	v, found = s.First()
	assert.True(t, found)
	assert.Equal(t, myEnumB, v)
	v, found = s.Last()
	assert.Equal(t, myEnumB, v)
	s.Add(myEnumC)
	v, found = s.Last()
	assert.Equal(t, myEnumC, v)
	v, found = s.RemoveFirst()
	assert.True(t, found)
	assert.Equal(t, myEnumB, v)
	v, found = s.RemoveLast()
	assert.True(t, found)
	assert.Equal(t, myEnumC, v)
	assert.Equal(t, 0, s.Size())

	s = newEnumSet[myEnum](myEnumC, myEnumA)
	assert.True(t, s.Contains(myEnumA))
	s.Clear()
	assert.Equal(t, 0, s.Size())
}
