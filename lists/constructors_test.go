package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type aStruct struct {
	v int
}

func TestNewCircularArrayListOfType(t *testing.T) {
	l := NewCircularArrayListOfType[*aStruct](3, func(a, b *aStruct) bool { return a.v > b.v })
	assert.False(t, l.Contains(&aStruct{2}))
	l.Add(&aStruct{3})
	assert.True(t, l.Contains(&aStruct{2}))
	assert.False(t, l.Contains(&aStruct{5}))
}
