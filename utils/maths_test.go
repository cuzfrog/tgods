package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max(1, 3))
	assert.Equal(t, 3, Max(3, 2))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 3, Min(5, 3))
	assert.Equal(t, 3, Min(3, 6))
}
