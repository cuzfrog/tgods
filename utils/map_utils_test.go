package utils

import (
	"github.com/cuzfrog/tgods/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapUtilsFunctions(t *testing.T) {
	m := mocks.NewMockMap[int, string]()
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(5, "f")
	keys := KeysFrom(m)
	assert.ElementsMatch(t, []int{1, 2, 5}, keys)

	values := ValuesFrom(m)
	assert.ElementsMatch(t, []string{"a", "b", "f"}, values)

	col := mocks.NewMockCollectionOf[string]()
	addCnt := ValuesTo[int, string](m, col)
	assert.Equal(t, addCnt, 3)
	assert.ElementsMatch(t, []string{"a", "b", "f"}, SliceFrom[string](col))

	v := Compute(m, 2, func(v string, found bool) string { return v + "-computed" })
	assert.Equal(t, "b-computed", v)
	v, _ = m.Get(2)
	assert.Equal(t, "b-computed", v)

	v, computed := ComputeIfAbsent(m, 5, func() string { return "computed" })
	assert.False(t, computed)
	assert.Equal(t, "f", v)
	v, computed = ComputeIfAbsent(m, 6, func() string { return "computed" })
	assert.True(t, computed)
	assert.Equal(t, "computed", v)
}

func TestMapUtilsFunctions_MultiMap(t *testing.T) {
	m := mocks.NewMockMultiMap[int, string]()
	m.PutSingle(1, "a")
	m.PutSingle(1, "b")
	m.PutSingle(2, "a")
	m.PutSingle(2, "b")
	assert.Equal(t, []string{"a", "b", "a", "b"}, ValuesFromMulti[int, string](m))

	col := mocks.NewMockCollectionOf[string]()
	addedCnt := MultiValuesTo[int, string](m, col)
	assert.Equal(t, 4, addedCnt)
	assert.Equal(t, []string{"a", "b", "a", "b"}, SliceFrom[string](col))
}
