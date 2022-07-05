package utils

import (
	"github.com/cuzfrog/tgods/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNil(t *testing.T) {
	assert.Equal(t, 0, Nil[int]())
	assert.Equal(t, "", Nil[string]())
	assert.Equal(t, nil, Nil[core.Bag[int]]())
}
