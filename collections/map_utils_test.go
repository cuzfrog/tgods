package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapUtilsFunctions(t *testing.T) {
	m := NewHashMapOfNumKey[int, string]()
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(5, "f")
	keys := utils.KeysFrom(m)
	assert.ElementsMatch(t, []int{1, 2, 5}, keys)

	values := utils.ValuesFrom(m)
	assert.ElementsMatch(t, []string{"a", "b", "f"}, values)

}
