package utils

import (
	"github.com/cuzfrog/tgods/lists"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	l := lists.NewLinkedList(1, 4, 3, 2)
	arr := Slice(l.Iterator(), l.Size())
	assert.Equal(t, []int{1, 4, 3, 2}, arr)
}
