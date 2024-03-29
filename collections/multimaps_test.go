package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArrayListMultiMap_Properties(t *testing.T) {
	m1 := NewArrayListMultiMapOf[int, string](funcs.NumHash[int], funcs.ValueEqual[int])
	m1.PutSingle(1, "a")
	m2 := NewArrayListMultiMapOfNumKey[int, string]()
	m2.PutSingle(1, "a")

	tests := []struct {
		name string
		m    types.MultiMap[int, string]
	}{
		{"arrListMultiMap1", m1},
		{"arrListMultiMap1", m2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			m.PutSingle(1, "a")
			m.PutSingle(2, "b")
			m.PutSingle(2, "c")
			l, found := m.Get(0)
			assert.False(t, found)

			l, found = m.Get(1)
			assert.True(t, found)
			assert.Equal(t, []string{"a", "a"}, utils.SliceFrom[string](l))

			l, found = m.Get(2)
			assert.True(t, found)
			assert.Equal(t, []string{"b", "c"}, utils.SliceFrom[string](l))
		})
	}
}

func TestNewArrayListMultiMapOfStrKey_Properties(t *testing.T) {
	m := NewArrayListMultiMapOfStrKey[int]()
	m.PutSingle("a", 1)
	m.PutSingle("a", 2)

	l, found := m.Get("a")
	assert.True(t, found)
	assert.Equal(t, []int{1, 2}, utils.SliceFrom[int](l))
}

func TestNewHashSetMultiMap_Properties(t *testing.T) {
	m1 := NewHashSetMultiMapOf[int, string](funcs.NumHash[int], funcs.ValueEqual[int], funcs.NewStrHash(), funcs.ValueEqual[string])
	m1.PutSingle(1, "a")
	m2 := NewHashSetMultiMapOfNumKey[int, string](funcs.NewStrHash(), funcs.ValueEqual[string])
	m2.PutSingle(1, "a")

	tests := []struct {
		name string
		m    types.MultiMap[int, string]
	}{
		{"hashSetMultiMap1", m1},
		{"hashSetMultiMap1", m2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.m
			m.PutSingle(1, "a")
			m.PutSingle(2, "b")
			m.PutSingle(2, "c")
			l, found := m.Get(0)
			assert.False(t, found)

			l, found = m.Get(1)
			assert.True(t, found)
			assert.Equal(t, []string{"a"}, utils.SliceFrom[string](l))

			l, found = m.Get(2)
			assert.True(t, found)
			assert.ElementsMatch(t, []string{"b", "c"}, utils.SliceFrom[string](l))
		})
	}
}

func TestNewHashSetMultiMapOfStrKey_Properties(t *testing.T) {
	m := NewHashSetMultiMapOfStrKey[int](funcs.NumHash[int], funcs.ValueEqual[int])
	m.PutSingle("a", 1)
	m.PutSingle("a", 2)
	m.PutSingle("a", 2)

	l, found := m.Get("a")
	assert.True(t, found)
	assert.ElementsMatch(t, []int{1, 2}, utils.SliceFrom[int](l))
}

func TestArrayListMultiMapConstraintInterface(t *testing.T) {
	m := NewArrayListMultiMapC[*intStruct, int]()
	k1 := &intStruct{3}
	m.PutSingle(k1, 2)
	assert.Equal(t, 1, m.Size())
	k2 := &intStruct{3}
	assert.True(t, m.ContainsKey(k2))
}

func TestHashSetMultiMapConstraintInterface(t *testing.T) {
	m := NewHashSetMultiMapC[*intStruct, int](funcs.NumHash[int], funcs.ValueEqual[int])
	k1 := &intStruct{3}
	m.PutSingle(k1, 2)
	assert.Equal(t, 1, m.Size())
	k2 := &intStruct{3}
	assert.True(t, m.ContainsKey(k2))

	m2 := NewHashSetMultiMapCC[*intStruct, *intStruct]()
	m2.PutSingle(k1, &intStruct{33})
	m2.PutSingle(k1, &intStruct{33})
	assert.Equal(t, 1, m2.Size())
}
