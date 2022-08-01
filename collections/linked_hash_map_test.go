package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedHashMap_PutWithLimit(t *testing.T) {
	m := newLinkedHashMap[int, int](funcs.NumHash[int], funcs.ValueEqual[int], 3, 3)
	m.Put(1, 100)
	m.Put(2, 200)
	m.Put(3, 300)
	m.Put(4, 400)
	assert.Equal(t, 3, m.Size())
	assert.False(t, m.ContainsKey(1))
	assert.Equal(t, 2, m.head.Value().Key())
	assert.Equal(t, 4, m.tail.Value().Key())

}
