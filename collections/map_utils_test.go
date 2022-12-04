package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapUtilsFunctions(t *testing.T) {
	m := NewHashMapOfNumKey[int, string]()
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(5, "f")
	keys := utils.KeysFrom(m)
	assert.ElementsMatch(t, []int{1, 2, 5}, keys)

	values := utils.ValuesFrom(m)
	assert.ElementsMatch(t, []string{"a", "b", "f"}, values)

	v := utils.Compute(m, 2, func(v string, found bool) string { return v + "-computed" })
	assert.Equal(t, "b-computed", v)
	v, _ = m.Get(2)
	assert.Equal(t, "b-computed", v)

	v, computed := utils.ComputeIfAbsent(m, 5, func() string { return "computed" })
	assert.False(t, computed)
	assert.Equal(t, "f", v)
	v, computed = utils.ComputeIfAbsent(m, 6, func() string { return "computed" })
	assert.True(t, computed)
	assert.Equal(t, "computed", v)
}

func TestMapUtilsFunctions_MultiMap(t *testing.T) {
	m := NewArrayListMultiMapOfNumKey[int, string]()
	m.PutSingle(1, "a")
	m.PutSingle(1, "b")
	m.PutSingle(2, "a")
	m.PutSingle(2, "b")
	assert.Equal(t, []string{"a", "b", "a", "b"}, utils.ValuesFromMulti[int, string](m))

	col := NewArrayListOf[string]()
	addedCnt := utils.MultiValuesTo[int, string](m, col)
	assert.Equal(t, 4, addedCnt)
	assert.Equal(t, []string{"a", "b", "a", "b"}, utils.SliceFrom[string](col))
}
