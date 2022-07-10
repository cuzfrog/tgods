package removed

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
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

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a2 := make([]int, 5)
	copy(a2, a)
	Shuffle(a2, rand.Intn)
	assert.NotEqual(t, a, a2)
}
