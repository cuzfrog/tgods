package collections

import "github.com/cuzfrog/tgods/types"

// NewTreeAdjacencyGraph creates a treeMap based adjacency list graph with a custom Compare function
func NewTreeAdjacencyGraph[V any, E any](comp types.Compare[V]) types.Graph[V, E] {
	return newTreeAdjacencyList[V, E](comp)
}
