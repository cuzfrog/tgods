package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var entry1 = EntryOf("a", 1)

func TestMapProperties(t *testing.T) {
	tests := []struct {
		name string
		m    types.Map[string, int]
	}{
		{"treeMap1", NewTreeMapOf[string, int](entry1)},
		{"treeMap2", NewTreeMapOfComp[string, int](funcs.ValueCompare[string], entry1)},
		{"hashMap1", NewHashMapOfStrKey[int](entry1)},
		{"hashMap2", NewHashMapOf[string, int](funcs.NewStrHash(), funcs.ValueEqual[string], entry1)},
		{"linkedHashMap1", NewLinkedHashMapOfStrKey[int](entry1)},
		{"linkedHashMap2", NewLinkedHashMapOf[string, int](funcs.NewStrHash(), funcs.ValueEqual[string], entry1)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			m.Add(EntryOf("a", 1))
			m.Put("b", 2)
			assert.Equal(t, 2, m.Size())
			v, found := m.Get("b")
			assert.True(t, found)
			assert.Equal(t, 2, v)
			_, found = m.Get("c")
			assert.False(t, found)

			assert.True(t, m.ContainsKey("a"))
			assert.False(t, m.ContainsKey("aa"))

			old, exist := m.Remove("5")
			assert.False(t, exist)
			assert.Equal(t, 2, m.Size())
			old, exist = m.Remove("a")
			assert.True(t, exist)
			assert.Equal(t, 1, old)
			assert.Equal(t, 1, m.Size())

			old, exist = m.Put("b", 3)
			assert.True(t, exist)
			assert.Equal(t, 2, old)
			v, found = m.Get("b")
			assert.True(t, found)
			assert.Equal(t, 3, v)

			m.Clear()
			assert.Equal(t, 0, m.Size())
			v, found = m.Get("b")
			assert.False(t, found)
		})
	}
}

func TestNewHashMapOfNumKey(t *testing.T) {
	m1 := NewLinkedHashMap[int, string](funcs.NumHash[int], funcs.ValueEqual[int], 0, OriginalOrder)
	m1.Put(1, "a")
	m1.Put(2, "b")

	tests := []struct {
		name string
		m    types.Map[int, string]
	}{
		{"linkedHashMap1", m1},
		{"hashMap1", NewHashMapOfNumKey(EntryOf(1, "a"), EntryOf(2, "b"))},
		{"linkedHashMap2", NewLinkedHashMapOfNumKey(EntryOf(1, "a"), EntryOf(2, "b"))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			assert.Equal(t, 2, m.Size())
			old, found := m.Put(1, "aa")
			assert.True(t, found)
			assert.Equal(t, "a", old)
		})
	}
}
