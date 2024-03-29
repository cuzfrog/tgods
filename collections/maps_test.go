package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
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

func TestSortedMapProperties(t *testing.T) {
	tests := []struct {
		name string
		m    types.SortedMap[int, int]
	}{
		{"treeMap1", NewTreeMapOf[int, int]()},
		{"treeMap2", NewTreeMapOfComp[int, int](funcs.ValueCompare[int])},
		{"enumMap", NewEnumMap[int, int](10)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			m.Put(3, 300)
			m.Put(6, 200)
			m.Put(2, 800)
			m.Put(1, 900)
			e := m.First()
			assert.Equal(t, 1, e.Key())
			assert.Equal(t, 900, e.Value())
			e = m.Last()
			assert.Equal(t, 6, e.Key())
			e = m.RemoveFirst()
			assert.Equal(t, 1, e.Key())
			assert.Equal(t, 3, m.Size())
			e = m.RemoveLast()
			assert.Equal(t, 6, e.Key())
			e = m.RemoveLast()
			assert.Equal(t, 3, e.Key())
			assert.Equal(t, 1, m.Size())
			e = m.RemoveLast()
			assert.Equal(t, 2, e.Key())
			assert.Equal(t, 0, m.Size())
			assert.Nil(t, m.RemoveLast())
			m.Put(5, 100)
			e = m.RemoveFirst()
			assert.Equal(t, 5, e.Key())
			assert.Nil(t, m.RemoveFirst())
			assert.Nil(t, m.First())
			assert.Nil(t, m.Last())
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

func TestLRUCache_Properties(t *testing.T) {
	tests := []struct {
		name string
		m    types.Map[int, string]
	}{
		{"LRU1", NewLRUCache[int, string](funcs.NumHash[int], funcs.ValueEqual[int], 3, GetOrder)},
		{"LRU2", NewLRUCacheOfNumKey[int, string](3, GetOrder)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			m.Put(1, "a")
			m.Put(3, "c")
			m.Put(2, "b")
			m.Get(1)
			m.Put(4, "d")
			assert.Equal(t, []int{2, 1, 4}, utils.KeysFrom(m))
		})
	}
}

func TestLRUCacheOfStrKey_Properties(t *testing.T) {
	tests := []struct {
		name string
		m    types.Map[string, string]
	}{
		{"LRU1", NewLRUCacheOfStrKey[string](3, GetOrder)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			m.Put("1", "a")
			m.Put("3", "c")
			m.Put("2", "b")
			m.Get("1")
			m.Put("4", "d")
			assert.Equal(t, []string{"2", "1", "4"}, utils.KeysFrom(m))
		})
	}
}

func TestMapConstraintTypeConstructors(t *testing.T) {
	tests := []struct {
		name string
		m    types.Map[*intStruct, int]
	}{
		{"treeMap1", NewTreeMapOfC[*intStruct, int]()},
		{"hashMap1", NewHashMapC[*intStruct, int]()},
		{"linkedHashMap1", NewLinkedHashMapC[*intStruct, int](0, OriginalOrder)},
		{"linkedHashMap2", NewLRUCacheC[*intStruct, int](0, OriginalOrder)},
		{"linkedHashMap3", NewLinkedHashMapOfC[*intStruct, int]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			k1 := &intStruct{3}
			m.Put(k1, 2)
			assert.Equal(t, 1, m.Size())
			k2 := &intStruct{3}
			assert.True(t, m.ContainsKey(k2))
		})
	}
}
