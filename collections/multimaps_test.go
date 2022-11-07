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
