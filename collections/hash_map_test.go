package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap(t *testing.T) {
	tests := []struct {
		name string
		h    types.Map[string, int]
	}{
		{"hashMap", newHashMap[string, int](funcs.NewStrHash(), funcs.ValueEqual[string])},
		{"linkedHashMap", newLinkedHashMap[string, int](funcs.NewStrHash(), funcs.ValueEqual[string], 0, 1)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := test.h
			assert.Equal(t, 0, h.Size())
			_, found := h.Put("a", 1)
			assert.False(t, found)
			old, found := h.Put("a", 2)
			assert.True(t, found)
			assert.Equal(t, 1, old)
			assert.Equal(t, 1, h.Size())
			assert.False(t, h.ContainsKey("b"))
			h.Put("b", 3)
			assert.True(t, h.ContainsKey("b"))
			v, found := h.Get("b")
			assert.True(t, found)
			assert.Equal(t, 3, v)
			v, found = h.Get("c")
			assert.False(t, found)
			v, found = h.Remove("c")
			assert.False(t, found)
			v, found = h.Remove("b")
			assert.True(t, found)
			assert.Equal(t, 3, v)
		})
	}

}
