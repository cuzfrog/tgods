package utils

import (
	"github.com/cuzfrog/tgods/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StringFrom(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](3, 7)
	assert.Equal(t, "[3, 7]", *StringFrom[int](c))
}
