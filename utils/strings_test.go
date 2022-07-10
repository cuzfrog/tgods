package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StringFrom(t *testing.T) {
	c := newMockCollectionOf[int](3, 7)
	assert.Equal(t, "[3, 7]", *StringFrom[int](c))
}
