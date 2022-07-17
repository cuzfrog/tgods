package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRbTree_Properties(t *testing.T) {
	tree := newRbTreeOfComp(funcs.ValueCompare[int])
	assert.False(t, tree.insert(10))
	assert.False(t, tree.insert(20))
	assert.True(t, tree.insert(20))
	r := tree.root
	/*
		        10b
			      20r
	*/
	assert.Equal(t, 10, r.v)
	assert.Equal(t, black, r.c)
	assert.Equal(t, 20, r.b.v)
	assert.Equal(t, red, r.b.c)

	assert.False(t, tree.insert(30))
	r = tree.root
	/*
	      20b
	   10r   30r
	*/
	assert.Equal(t, 20, r.v)
	assert.Equal(t, black, r.c)
	assert.Equal(t, 10, r.a.v)
	assert.Equal(t, red, r.a.c)
	assert.Equal(t, 30, r.b.v)
	assert.Equal(t, red, r.b.c)

	assert.False(t, tree.insert(40))
	r = tree.root
	/*
	      20b
	   10b   30b
	           40r
	*/
	assert.Equal(t, 20, r.v)
	assert.Equal(t, black, r.c)
	assert.Equal(t, 10, r.a.v)
	assert.Equal(t, black, r.a.c)
	assert.Equal(t, 30, r.b.v)
	assert.Equal(t, black, r.b.c)
	assert.Equal(t, 40, r.b.b.v)
	assert.Equal(t, red, r.b.b.c)

	assert.False(t, tree.insert(50))
	r = tree.root
	/*
	      20b
	   10b   40b
	       30r 50r
	*/
	assert.Equal(t, 30, r.b.a.v)
	assert.Equal(t, red, r.b.a.c)
	assert.Equal(t, 50, r.b.b.v)
	assert.Equal(t, red, r.b.b.c)

	assert.False(t, tree.insert(60))
	r = tree.root
	/*
	      20b
	   10b   40r
	       30b 50b
	             60r
	*/
	l := bfTraverse[int](r)
	assert.Equal(t, []int{20, 10, 40, 30, 50, 60}, utils.SliceFrom[int](l))

	assert.False(t, tree.insert(70))
	r = tree.root
	/*
	      20b
	   10b   40r
	       30b 60b
	         50r  70r

	        40b
	   20r      60b
	 10b 30b  50r 70r
	*/
	l = bfTraverse[int](r)
	assert.Equal(t, []int{40, 20, 60, 10, 30, 50, 70}, utils.SliceFrom[int](l))

	dd, found := tree.delete(5)
	assert.False(t, found)
	dd, found = tree.delete(40)
	assert.Equal(t, 40, dd)
	assert.Equal(t, 6, tree.size)
	assert.True(t, found)
	tree.delete(10)
	tree.delete(20)
	tree.delete(30)
	tree.delete(50)
	tree.delete(60)
	l = bfTraverse[int](tree.root)
	assert.Equal(t, []int{70}, utils.SliceFrom[int](l))
}
