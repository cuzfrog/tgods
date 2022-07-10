package funcs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueComparator(t *testing.T) {
	assert.Equal(t, int8(-1), ValueCompare(1, 2))
	assert.Equal(t, int8(1), ValueCompare("c", "b"))
	assert.Equal(t, int8(0), ValueCompare(2, 2))
	assert.True(t, 0 == ValueCompare(2, 2))
}

func TestEqualComparable(t *testing.T) {
	assert.True(t, ValueEqual(1, 1))
	assert.False(t, ValueEqual("1", "2"))
}

func TestCompToEq(t *testing.T) {
	fn := CompToEq(ValueCompare[int])
	assert.True(t, fn(1, 1))
	assert.False(t, fn(1, 2))
}

func TestInverseComp(t *testing.T) {
	fn := InverseComp(ValueCompare[int])
	assert.Equal(t, int8(0), fn(1, 1))
	assert.Equal(t, int8(-1), fn(2, 1))
	assert.Equal(t, int8(1), fn(2, 3))
}
