package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type myEnum uint8

const (
	myEnumA myEnum = iota
	myEnumB
	myEnumC
)

func TestEnumMap(t *testing.T) {
	m := newEnumMap[myEnum, string](myEnumC)
	assert.Equal(t, 0, m.Size())
	old, found := m.Put(myEnumA, "a")
	assert.False(t, found)
	assert.Equal(t, 1, m.Size())
	old, found = m.Put(myEnumA, "aa")
	assert.True(t, found)
	assert.Equal(t, "a", old)
	m.Put(myEnumB, "b")
	assert.Equal(t, 2, m.Size())

	assert.False(t, m.ContainsKey(myEnumC))
	m.Put(myEnumC, "c")
	assert.True(t, m.ContainsKey(myEnumC))
	assert.Panics(t, func() { m.Contains(EntryOf(myEnumC, "cc")) })

	v, found := m.Get(myEnumA)
	assert.True(t, found)
	assert.Equal(t, "aa", v)
	old, found = m.Remove(myEnumA)
	assert.True(t, found)
	assert.Equal(t, "aa", old)
	assert.Equal(t, 2, m.Size())
	assert.False(t, m.ContainsKey(myEnumA))
	old, found = m.Remove(myEnumA)
	assert.False(t, found)
	v, found = m.Get(myEnumA)
	assert.False(t, found)

	m.Clear()
	assert.Equal(t, 0, m.size)
	for i := 0; i < 3; i++ {
		assert.Nil(t, m.arr[i])
	}
}
