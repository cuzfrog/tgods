package utils

import (
	"github.com/cuzfrog/tgods/lists"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_String(t *testing.T) {
	l := lists.NewLinkedList(3, 7)
	assert.Equal(t, "[3, 7]", *StringFrom[int](l))
}
