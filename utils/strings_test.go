package utils

import (
	"fmt"
	"github.com/cuzfrog/tgods/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StringFrom(t *testing.T) {
	c := mocks.NewMockCollectionOf[int](3, 7)
	assert.Equal(t, "[3, 7]", *StringFrom[int](c))
}

func Test_StringFromf(t *testing.T) {
	c := mocks.NewMockCollectionOf[uint8](65, 66)
	assert.Equal(t, "[A, B]", *StringFromf[uint8](c, func(elem uint8) string { return fmt.Sprintf("%c", elem) }))
}
