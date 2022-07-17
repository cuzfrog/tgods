package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeMap_Contains(t *testing.T) {
	m := NewTreeMapOf[int, int]()
	assert.Panics(t, func() { m.Contains(nil) })
}
