package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueComparator(t *testing.T) {
	assert.Equal(t, int8(-1), CompareOrdered(1, 2))
	assert.Equal(t, int8(1), CompareOrdered("c", "b"))
	assert.Equal(t, int8(0), CompareOrdered(2, 2))
	assert.True(t, 0 == CompareOrdered(2, 2))
}
