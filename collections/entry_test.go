package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryOf(t *testing.T) {
	e := EntryOf("a", 1)
	assert.Equal(t, "a", e.Key())
	assert.Equal(t, 1, e.Value())
}

func TestKeyEntry(t *testing.T) {
	e := keyEntry[int, int]{5}
	assert.Equal(t, 5, e.Key())
	assert.Panics(t, func() { e.Value() })
}
