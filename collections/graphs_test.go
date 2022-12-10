package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphConstructors(t *testing.T) {
	tests := []struct {
		name string
		g    types.Graph[*intStruct, int]
	}{
		{"treeMapGraph1", NewTreeAdjacencyGraph[*intStruct, int](intStructCompare)},
		{"treeMapGraph2", NewTreeAdjacencyGraphC[*intStruct, int]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := test.g
			v3 := &intStruct{3}
			v5 := &intStruct{5}
			g.Add(v3)
			g.Add(v5)
			g.Connect(v3, v5, 35)
			assert.Equal(t, 1, g.OutwardCount(v3))
		})
	}
}
