package collections

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rbNode_insert_root(t *testing.T) {
	r, found, nn := insert(nil, 3, compInt)
	assert.False(t, found)
	assert.Equal(t, rbNode[int]{3, nil, nil, nil, true}, *r)
	assert.Equal(t, r, nn)
}

func Test_rbNode_insert(t *testing.T) {
	/*         70
	    30
	20      50
	      40 60
	*/
	n70 := newRbNode(70, nil)
	n30 := newRbNode(30, n70)
	n70.a = n30
	n20 := newRbNode(20, n30)
	n30.a = n20
	n50 := newRbNode(50, n30)
	n30.b = n50
	n40 := newRbNode(40, n50)
	n50.a = n40
	n60 := newRbNode(60, n50)
	n50.b = n60

	r, found, n20 := insert(n70, 20, compInt)
	assert.True(t, found)
	assert.Equal(t, n70, r)
	assert.Equal(t, n20, n20)

	r, found, n10 := insert(n70, 10, compInt)
	assert.Equal(t, n10, n20.a)
	assert.Equal(t, n20, n10.p)
	assert.Equal(t, 10, n10.v)
	/*         70
	    	30
		20      50
	   10     40 60
	*/
	r, found, n15 := insert(n20, 15, compInt)
	assert.Equal(t, n15, n10.b)
	assert.Equal(t, n10, n15.p)

	r, found, n45 := insert(n30, 45, compInt)
	assert.Equal(t, n45, n40.b)
	assert.Equal(t, n40, n45.p)

	r, found, n35 := insert(n70, 35, compInt)
	assert.Equal(t, n40, n35.p)
	assert.Equal(t, n35, n40.a)
}

func Test_rbNode_rebalance_noAction(t *testing.T) {
	n := newRbNode(3, nil)
	r := rebalance(n)
	assert.Equal(t, black, n.c) // root color as black
	assert.Nil(t, r)            // no next node for rectification

	/*
		    30b
		20b      50r
	*/
	n30 := newRbNode(30, nil)
	n30.c = black
	n20 := newRbNode(20, n30)
	n30.a = n20
	n20.c = black
	n50 := newRbNode(50, n30)
	n30.b = n50
	r = rebalance(n50)
	assert.Nil(t, r)
	assert.Equal(t, black, n30.c)
	assert.Equal(t, n50, n30.b)
}

func Test_rbNode_rebalance_recolorRight(t *testing.T) {
	/*
		    30b
		20r      50r
		      40r
	*/
	n30 := newRbNode(30, nil)
	n30.c = black
	n20 := newRbNode(20, n30)
	n20.c = red
	n30.a = n20
	n50 := newRbNode(50, n30)
	n50.c = red
	n30.b = n50
	n40 := newRbNode(40, n50)
	n40.c = red
	n50.a = n40
	r := rebalance(n40)
	assert.Equal(t, n30, r)
	assert.Equal(t, red, n40.c)
	assert.Equal(t, black, n20.c)
	assert.Equal(t, black, n50.c)
	assert.Equal(t, red, n30.c)
	l := bfTraverse[int](n30)
	assert.Equal(t, []int{30, 20, 50, 40}, utils.SliceFrom[int](l))
}

func Test_rbNode_rebalance_recolorLeft(t *testing.T) {
	/*
		    30b
		20r      50r
		  25r
	*/
	n30 := newRbNode(30, nil)
	n30.c = black
	n20 := newRbNode(20, n30)
	n20.c = red
	n25 := newRbNode(25, n20)
	n25.c = red
	n20.b = n25
	n30.a = n20
	n50 := newRbNode(50, n30)
	n50.c = red
	n30.b = n50
	r := rebalance(n25)
	assert.Equal(t, n30, r)
	assert.Equal(t, red, n25.c)
	assert.Equal(t, black, n20.c)
	assert.Equal(t, black, n50.c)
	assert.Equal(t, red, n30.c)
	l := bfTraverse[int](r)
	assert.Equal(t, []int{30, 20, 50, 25}, utils.SliceFrom[int](l))
}

func Test_rbNode_rebalance_rotateLR(t *testing.T) {
	/*
		    30b
		20r      50b
		  25r
	*/
	n30 := newRbNode(30, nil)
	n30.c = black
	n20 := newRbNode(20, n30)
	n20.c = red
	n25 := newRbNode(25, n20)
	n25.c = red
	n20.b = n25
	n30.a = n20
	n50 := newRbNode(50, n30)
	n50.c = black
	n30.b = n50

	r := rebalance(n25)
	/*
		    25b
		20r      30r
		            50b
	*/
	assert.Equal(t, n25, r)
	assert.Equal(t, black, n25.c)
	assert.Equal(t, red, n20.c)
	assert.Equal(t, black, n50.c)
	assert.Equal(t, red, n30.c)
	l := bfTraverse[int](r)
	assert.Equal(t, []int{25, 20, 30, 50}, utils.SliceFrom[int](l))
}

func Test_swapLR(t *testing.T) {
	/*
				    30b
				20r      50b
			      25r
		         23 27
	*/
	n30 := newRbNode(30, nil)
	n30.c = black
	n20 := newRbNode(20, n30)
	n20.c = red
	n25 := newRbNode(25, n20)
	n25.c = red
	n20.b = n25
	n30.a = n20
	n23 := newRbNode(23, n25)
	n25.a = n23
	n50 := newRbNode(50, n30)
	n50.c = black
	n30.b = n50
	n27 := newRbNode(27, n25)
	n25.b = n27
	/*
					    30b
					25r      50b
			     20r   27
		           23
	*/
	rotateLeft(n20)
	l := bfTraverse[int](n30)
	assert.Equal(t, []int{30, 25, 50, 20, 27, 23}, utils.SliceFrom[int](l))
	assert.Equal(t, n25, n30.a)
	assert.Equal(t, n30, n25.p)
	assert.Equal(t, n20, n25.a)
	assert.Equal(t, n25, n20.p)
	assert.Equal(t, n23, n20.b)
	assert.Equal(t, n20, n23.p)
	assert.Equal(t, n27, n25.b)
	assert.Equal(t, n25, n27.p)
}

func Test_swapRL(t *testing.T) {
	/*
				    30b
				20b      50r
			           40r
		              35  45
	*/
	n30 := newRbNode(30, nil)
	n30.c = black
	n20 := newRbNode(20, n30)
	n20.c = red
	n30.a = n20
	n50 := newRbNode(50, n30)
	n50.c = black
	n30.b = n50
	n40 := newRbNode(40, n50)
	n50.a = n40
	n40.c = red
	n35 := newRbNode(35, n40)
	n40.a = n35
	n45 := newRbNode(45, n40)
	n40.b = n45
	l := bfTraverse[int](n30)
	assert.Equal(t, []int{30, 20, 50, 40, 35, 45}, utils.SliceFrom[int](l))
	/*
					    30b
					20b        40b
			                35   50r
		                        45
	*/
	rotateRight(n50)
	l = bfTraverse[int](n30)
	assert.Equal(t, n20, n30.a)
	assert.Equal(t, n40, n30.b)
	assert.Equal(t, n30, n40.p)
	assert.Equal(t, n35, n40.a)
	assert.Equal(t, n40, n35.p)
	assert.Equal(t, n50, n40.b)
	assert.Equal(t, n40, n50.p)
	assert.Equal(t, n45, n50.a)
	assert.Equal(t, n50, n45.p)
	assert.Equal(t, []int{30, 20, 40, 35, 50, 45}, utils.SliceFrom[int](l))
}

func Test_rbNode_rotate(t *testing.T) {
	n := &rbNode[int]{3, nil, nil, nil, true}
	assert.Equal(t, newRbNode(3, nil), rotateLeft(n))
	assert.Equal(t, newRbNode(3, nil), rotateRight(n))

	/*
		         7
			   3
			2     5
			     4 6
	*/
	n7 := newRbNode(7, nil)
	n3 := newRbNode(3, n7)
	n7.a = n3
	n2 := newRbNode(2, n3)
	n3.a = n2
	n5 := newRbNode(5, n3)
	n3.b = n5
	n4 := newRbNode(4, n5)
	n5.a = n4
	n6 := newRbNode(6, n5)
	n5.b = n6
	n = n3
	/*    7
	    5
	  3   6
	2  4
	*/
	r := rotateLeft(n3)
	assert.Equal(t, 5, r.v)
	assert.Equal(t, 3, r.a.v)
	assert.Equal(t, 2, r.a.a.v)
	assert.Nil(t, r.a.a.a)
	assert.Nil(t, r.a.a.b)
	assert.Equal(t, 6, r.b.v)
	assert.Nil(t, r.b.b)
	assert.Equal(t, 4, r.a.b.v)
	assert.Nil(t, r.a.b.a)
	assert.Nil(t, r.a.b.b)
	// assert relationship:
	assert.Equal(t, n5, n7.a)
	assert.Equal(t, n7, n5.p)
	assert.Equal(t, n3, n5.a)
	assert.Equal(t, n5, n3.p)
	assert.Equal(t, n6, n5.b)
	assert.Equal(t, n5, n6.p)
	assert.Equal(t, n2, n3.a)
	assert.Equal(t, n3, n2.p)
	assert.Equal(t, n4, n3.b)
	assert.Equal(t, n3, n4.p)

	/*   7
	   3
	2     5
	     4 6
	*/
	r = rotateRight(n5)
	assert.Equal(t, 3, r.v)
	assert.Equal(t, 2, r.a.v)
	assert.Nil(t, r.a.a)
	assert.Nil(t, r.a.b)
	assert.Equal(t, 5, r.b.v)
	assert.Equal(t, 4, r.b.a.v)
	assert.Nil(t, r.b.a.a)
	assert.Nil(t, r.b.a.b)
	assert.Equal(t, 6, r.b.b.v)
	assert.Nil(t, r.b.b.a)
	assert.Nil(t, r.b.b.b)
	// assert relationship:
	assert.Equal(t, n3, n7.a)
	assert.Equal(t, n7, n3.p)
	assert.Equal(t, n2, n3.a)
	assert.Equal(t, n3, n2.p)
	assert.Equal(t, n5, n3.b)
	assert.Equal(t, n3, n5.p)
	assert.Equal(t, n4, n5.a)
	assert.Equal(t, n5, n4.p)
	assert.Equal(t, n6, n5.b)
	assert.Equal(t, n5, n6.p)
}
