package collections

import "github.com/cuzfrog/tgods/types"

// NewTreeAdjacencyGraph creates a treeMap based adjacency list graph with a custom Compare function
func NewTreeAdjacencyGraph[V any, E any](comp types.Compare[V]) types.Graph[V, E] {
	return newTreeAdjacencyList[V, E](comp)
}

// NewTreeAdjacencyGraphC creates a treeMap based adjacency list graph with a constrained vertex type that implements custom Hash and Equal
func NewTreeAdjacencyGraphC[V types.WithCompare[V], E any]() types.Graph[V, E] {
	comp := func(v, o V) int8 { return v.Compare(o) }
	return newTreeAdjacencyList[V, E](comp)
}
