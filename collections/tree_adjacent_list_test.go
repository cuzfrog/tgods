package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeAdjacencyList_Properties(t *testing.T) {
	g := newTreeAdjacencyList[int, float64](funcs.ValueCompare[int])
	g.Add(1)
	g.Add(2)
	g.Add(3)
	assert.Equal(t, 3, g.Size())
	assert.True(t, g.Contains(2))

	assert.Equal(t, 0, g.InwardCount(1))
	assert.Equal(t, 0, g.InwardCount(11))
	assert.Equal(t, 0, g.OutwardCount(3))
	assert.Equal(t, 0, g.OutwardCount(44))
	it, size := g.OutwardEdges(3)
	assert.False(t, it.Next())
	assert.Equal(t, 0, size)
	it, size = g.OutwardEdges(333)
	assert.Equal(t, 0, size)
	it, size = g.InwardEdges(2)
	assert.False(t, it.Next())
	assert.Equal(t, 0, size)

	/*
		1 -> 2
		2
		3
		--
		2 <- 1
	*/
	oldE, exist := g.Connect(1, 2, 0.5)
	assert.False(t, g.inward.ContainsKey(1))
	assert.False(t, g.inward.ContainsKey(3))
	assert.False(t, exist)
	assert.Equal(t, 1, g.InwardCount(2))
	assert.Equal(t, 1, g.OutwardCount(1))
	it, size = g.InwardEdges(2)
	assert.Equal(t, 1, size)
	assert.True(t, it.Next())
	assert.Equal(t, 0.5, it.Value().Value())
	it, size = g.OutwardEdges(1)
	assert.Equal(t, 1, size)
	assert.True(t, it.Next())
	assert.Equal(t, 0.5, it.Value().Value())
	oldE, exist = g.Connect(1, 2, 0.8)
	assert.True(t, exist)
	assert.Equal(t, 0.5, oldE)
	it, size = g.InwardEdges(2)
	it.Next()
	assert.Equal(t, 0.8, it.Value().Value())
	it, size = g.OutwardEdges(1)
	it.Next()
	assert.Equal(t, 0.8, it.Value().Value())
	/*
			1 -> 2(0.8) 3(1.5)
			2
			3
			--
			2 <- 1(0.8)
		    3 <- 1(1.5)
	*/
	oldE, exist = g.Connect(1, 3, 1.5)
	it, size = g.OutwardEdges(1)
	assert.Equal(t, 2, size)
	assert.Equal(t, []float64{0.8, 1.5}, utils.SliceProjectIt(it, size, func(t types.Entry[int, float64]) float64 {
		return t.Value()
	}))
	assert.False(t, g.inward.ContainsKey(1))
	assert.True(t, g.inward.ContainsKey(2))
	assert.True(t, g.inward.ContainsKey(3))

	/*
			1 -> 2(0.8) 3(1.5)
			2 -> 1(0.7) 3(3.5)
			3
			--
			1 <- 2(0.7)
			2 <- 1(0.8)
		    3 <- 1(1.5) 2(3.5)
	*/
	oldE, exist = g.Connect(2, 3, 3.5)
	oldE, exist = g.Connect(2, 1, 0.7)
	assert.Equal(t, 3, g.Size())
	it, size = g.InwardEdges(3)
	assert.Equal(t, 2, size)
	assert.Equal(t, []float64{1.5, 3.5}, utils.SliceProjectIt(it, size, func(t types.Entry[int, float64]) float64 {
		return t.Value()
	}))
	it, size = g.OutwardEdges(2)
	assert.Equal(t, []float64{0.7, 3.5}, utils.SliceProjectIt(it, size, func(t types.Entry[int, float64]) float64 {
		return t.Value()
	}))

	assert.False(t, g.Remove(22))
	/*
			1 -> 3(1.5)
			3
			--
			1 <-
		    3 <- 1(1.5)
	*/
	assert.True(t, g.Remove(2))
	assert.Equal(t, 2, g.Size())
	assert.False(t, g.treeMap.ContainsKey(2))
	v, found := g.treeMap.Get(3)
	assert.True(t, found)
	assert.Equal(t, 0, v.Size())
	it, size = g.OutwardEdges(1)
	assert.Equal(t, 1, size)
	assert.True(t, it.Next())
	assert.Equal(t, 1.5, it.Value().Value())
	assert.False(t, it.Next())
	assert.False(t, g.inward.ContainsKey(2))
	v, found = g.inward.Get(1)
	assert.True(t, found)
	assert.Equal(t, 0, v.Size())
	it, size = g.InwardEdges(3)
	assert.Equal(t, 1, size)
	assert.True(t, it.Next())
	assert.Equal(t, 1.5, it.Value().Value())
	assert.False(t, it.Next())

	edge, exist := g.Disconnect(3, 1)
	assert.False(t, exist)
	edge, exist = g.Disconnect(11, 33)
	assert.False(t, exist)

	/*
			1 ->
			3
			--
			1 <-
		    3 <-
	*/
	edge, exist = g.Disconnect(1, 3)
	assert.True(t, exist)
	assert.Equal(t, 1.5, edge)
	assert.Equal(t, 2, g.Size())

	/*
			3
			--
		    3 <-
	*/
	g.Remove(1)
	assert.Equal(t, 1, g.Size())
	assert.False(t, g.Contains(1))

	g.Remove(3)
	assert.Equal(t, 0, g.Size())
	assert.Equal(t, 0, g.treeMap.Size())
	assert.Equal(t, 0, g.inward.Size())

	/*
				3
		        5 -> 3(0.6)
				--
			    3 < 5(0.6)
	*/
	g.Connect(5, 3, 0.6)
	assert.Equal(t, 2, g.Size())
	outwards, found := g.treeMap.Get(5)
	assert.True(t, found)
	f, found := outwards.Get(3)
	assert.Equal(t, 0.6, f)
	inwards, found := g.inward.Get(3)
	assert.True(t, found)
	f, found = inwards.Get(5)
	assert.True(t, found)
	assert.Equal(t, 0.6, f)
}
