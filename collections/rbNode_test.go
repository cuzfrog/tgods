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
	n70 := newRbNode(70, nil, false, false)
	n30 := newRbNode(30, n70, left, false)
	n20 := newRbNode(20, n30, left, false)
	n50 := newRbNode(50, n30, right, false)
	n40 := newRbNode(40, n50, left, false)
	newRbNode(60, n50, right, false)

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

func Test_rbNode_delete(t *testing.T) {
	t.Run("simple red", func(t *testing.T) {
		/*
				    30b
				20b      50b
			  15r  25r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, black)
		n25 := newRbNode(25, n20, right, red)
		n15 := newRbNode(15, n20, left, red)
		newRbNode(50, n30, right, black)
		nd, found := deleteNode(n30, 15, compInt)
		/*
				    30b
				20b      50b
			 (15r)  25r
		*/
		assert.Same(t, n15, nd)
		assert.True(t, found)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, red, n25.c)
		assert.Nil(t, n20.a)
		assert.Nil(t, n15.p)
	})
	t.Run("simple parent red", func(t *testing.T) {
		/*
				    30b
				20r      50b
			  15b  25b
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n25 := newRbNode(25, n20, right, black)
		n15 := newRbNode(15, n20, left, black)
		newRbNode(50, n30, right, black)
		nd, found := deleteNode(n30, 15, compInt)
		/*
				    30b
				20b      50b
			 (15b)  25r
		*/
		assert.Same(t, n15, nd)
		assert.True(t, found)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, red, n25.c)
		assert.Nil(t, n20.a)
		assert.Nil(t, n15.p)
	})
	t.Run("parent black", func(t *testing.T) {
		t.Run("sibling black, red child RR", func(t *testing.T) {
			/*
					    30b
					20b      50b
				  15b  25b      ...
					     27r
			*/
			n30 := newRbNode(30, nil, false, black)
			n20 := newRbNode(20, n30, left, black)
			n25 := newRbNode(25, n20, right, black)
			n15 := newRbNode(15, n20, left, black)
			newRbNode(50, n30, right, black)
			n27 := newRbNode(27, n25, right, red)
			nd, _ := deleteNode(n30, 15, compInt)
			/*
					    30b
					25b      50b
				 20b  27b       ...
			*/
			assert.Same(t, n15, nd)
			assert.Equal(t, black, n20.c)
			assert.Equal(t, black, n25.c)
			assert.Equal(t, black, n27.c)
			assert.Equal(t, n20, n25.a)
			assert.Equal(t, n25, n20.p)
			assert.Equal(t, n27, n25.b)
			assert.Equal(t, n25, n27.p)
		})
		t.Run("sibling black, red child RL", func(t *testing.T) {
			/*
					    30b
					20b      50b
				  15b  25b      ...
					  23r
			*/
			n30 := newRbNode(30, nil, false, black)
			n20 := newRbNode(20, n30, left, black)
			n25 := newRbNode(25, n20, right, black)
			n15 := newRbNode(15, n20, left, black)
			newRbNode(50, n30, right, black)
			n23 := newRbNode(23, n25, left, red)
			nd, _ := deleteNode(n30, 15, compInt)
			/*
					    30b
					23b      50b
				 20b  25b       ...
			*/
			assert.Same(t, n15, nd)
			assert.Equal(t, black, n20.c)
			assert.Equal(t, black, n25.c)
			assert.Equal(t, black, n23.c)
			assert.Equal(t, n23, n30.a)
			assert.Equal(t, n30, n23.p)
			assert.Equal(t, n20, n23.a)
			assert.Equal(t, n25, n23.b)
		})
		t.Run("sibling black, red child LL", func(t *testing.T) {
			/*
						    30b
						20b      50b
					  15b  25b      ...
				    13r
			*/
			n30 := newRbNode(30, nil, false, black)
			n20 := newRbNode(20, n30, left, black)
			n25 := newRbNode(25, n20, right, black)
			n15 := newRbNode(15, n20, left, black)
			newRbNode(50, n30, right, black)
			n13 := newRbNode(23, n15, left, red)
			nd, _ := deleteNode(n30, 25, compInt)
			/*
					    30b
					15b      50b
				 13b  20b        ...
			*/
			assert.Same(t, n25, nd)
			assert.Equal(t, black, n20.c)
			assert.Equal(t, black, n15.c)
			assert.Equal(t, black, n13.c)
			assert.Equal(t, n15, n30.a)
			assert.Equal(t, n13, n15.a)
			assert.Equal(t, n20, n15.b)
		})
		t.Run("sibling black, with black children", func(t *testing.T) {

		})
		t.Run("sibling red", func(t *testing.T) {

		})
	})
}

func Test_rbNode_rebalance(t *testing.T) {
	t.Run("noAction", func(t *testing.T) {
		n := newRbNode(3, nil, false, black)
		r := rebalance(n)
		assert.Nil(t, r) // no next node for rectification

		/*
			    30b
			20b      50r
		*/
		n30 := newRbNode(30, nil, false, black)
		newRbNode(20, n30, left, black)
		n50 := newRbNode(50, n30, right, false)
		n30.b = n50
		r = rebalance(n50)
		assert.Nil(t, r)
		assert.Equal(t, black, n30.c)
		assert.Equal(t, n50, n30.b)
	})

	t.Run("recolorRight", func(t *testing.T) {
		/*
			    30b
			20r      50r
			      40r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n50 := newRbNode(50, n30, right, red)
		n40 := newRbNode(40, n50, left, red)
		r := rebalance(n40)
		assert.Equal(t, n30, r)
		assert.Equal(t, red, n40.c)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, black, n50.c)
		assert.Equal(t, red, n30.c)
		l := bfTraverse[int](n30)
		assert.Equal(t, []int{30, 20, 50, 40}, utils.SliceFrom[int](l))
	})
	t.Run("recolorLeft", func(t *testing.T) {
		/*
			    30b
			20r      50r
			  25r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n25 := newRbNode(25, n20, right, red)
		n50 := newRbNode(50, n30, right, red)
		r := rebalance(n25)
		assert.Equal(t, n30, r)
		assert.Equal(t, red, n25.c)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, black, n50.c)
		assert.Equal(t, red, n30.c)
		l := bfTraverse[int](r)
		assert.Equal(t, []int{30, 20, 50, 25}, utils.SliceFrom[int](l))
	})

	t.Run("LL", func(t *testing.T) {
		/*
				    30b
				20r      50b
			  15r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n15 := newRbNode(15, n20, left, red)
		n50 := newRbNode(50, n30, right, black)

		r := rebalance(n15)
		/*
			    20b
			15r     30r
			            50b
		*/
		assert.Equal(t, n20, r)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, red, n15.c)
		assert.Equal(t, black, n50.c)
		assert.Equal(t, red, n30.c)
		l := bfTraverse[int](r)
		assert.Equal(t, []int{20, 15, 30, 50}, utils.SliceFrom[int](l))
	})
	t.Run("LR", func(t *testing.T) {
		/*
			    30b
			20r      50b
			  25r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n25 := newRbNode(25, n20, right, red)
		n50 := newRbNode(50, n30, right, black)

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
	})
	t.Run("RR", func(t *testing.T) {
		/*
			    30b
			20b      50r
			            55r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, black)
		n50 := newRbNode(50, n30, right, red)
		n55 := newRbNode(55, n50, right, red)
		n55.c = red
		n50.b = n55

		r := rebalance(n55)
		/*
				    50b
				30r      55r
			  20b
		*/
		assert.Equal(t, n50, r)
		assert.Equal(t, black, n50.c)
		assert.Equal(t, red, n30.c)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, red, n55.c)
		l := bfTraverse[int](r)
		assert.Equal(t, []int{50, 30, 55, 20}, utils.SliceFrom[int](l))
	})
	t.Run("RL", func(t *testing.T) {
		/*
			    30b
			20b      50r
			       45r
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, black)
		n50 := newRbNode(50, n30, right, red)
		n45 := newRbNode(45, n50, left, red)

		r := rebalance(n45)
		/*
				    45b
				30r      50r
			  20b
		*/
		assert.Equal(t, n45, r)
		assert.Equal(t, red, n50.c)
		assert.Equal(t, red, n30.c)
		assert.Equal(t, black, n20.c)
		assert.Equal(t, black, n45.c)
		l := bfTraverse[int](r)
		assert.Equal(t, []int{45, 30, 50, 20}, utils.SliceFrom[int](l))
	})
}

func Test_rbNode_rotate(t *testing.T) {
	t.Run("rotate left and right", func(t *testing.T) {
		n := &rbNode[int]{3, nil, nil, nil, red}
		assert.Equal(t, newRbNode(3, nil, false, red), rotateLeft(n))
		assert.Equal(t, newRbNode(3, nil, false, red), rotateRight(n))

		/*
			         7
				   3
				2     5
				     4 6
		*/
		n7 := newRbNode(7, nil, false, false)
		n3 := newRbNode(3, n7, left, false)
		n2 := newRbNode(2, n3, left, false)
		n5 := newRbNode(5, n3, right, false)
		n4 := newRbNode(4, n5, left, false)
		n6 := newRbNode(6, n5, right, false)
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
	})

	t.Run("swap LR", func(t *testing.T) {
		/*
					    30b
					20r      50b
				      25r
			         23 27
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n25 := newRbNode(25, n20, right, red)
		n23 := newRbNode(23, n25, left, false)
		newRbNode(50, n30, right, black)
		n27 := newRbNode(27, n25, right, false)
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
	})

	t.Run("swap RL", func(t *testing.T) {
		/*
					    30b
					20b      50r
				           40r
			              35  45
		*/
		n30 := newRbNode(30, nil, false, black)
		n20 := newRbNode(20, n30, left, red)
		n50 := newRbNode(50, n30, right, black)
		n40 := newRbNode(40, n50, left, red)
		n35 := newRbNode(35, n40, left, false)
		n45 := newRbNode(45, n40, right, false)
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
	})
}

func Test_rbNode_swapDown(t *testing.T) {
	t.Run("1 round to successor", func(t *testing.T) {
		/*
					     30b
					20b        50r
				            40r
			              35  45
		*/
		n30 := newRbNode(30, nil, false, black)
		newRbNode(20, n30, left, red)
		n50 := newRbNode(50, n30, right, black)
		n40 := newRbNode(40, n50, left, red)
		n35 := newRbNode(35, n40, left, false)
		newRbNode(45, n40, right, false)
		r := n30
		/*
					     35b
					20b        50r
				            40r
			              30  45
		*/
		ns := swapDown(r)
		assert.Same(t, ns, n35)
		l := bfTraverse[int](r)
		assert.Equal(t, []int{35, 20, 50, 40, 30, 45}, utils.SliceFrom[int](l))
	})

	t.Run("3 round to successor", func(t *testing.T) {
		/*
						     30b
						20b        50r
					            40r
				                  45
			                       46
		*/
		n30 := newRbNode(30, nil, false, black)
		newRbNode(20, n30, left, red)
		n50 := newRbNode(50, n30, right, black)
		n40 := newRbNode(40, n50, left, red)
		n45 := newRbNode(45, n40, right, false)
		n46 := newRbNode(46, n45, right, false)
		r := n30
		/*
						     40b
						20b        50r
					            45r
				                  46
			                       30
		*/
		ns := swapDown(r)
		assert.Same(t, ns, n46)
		l := bfTraverse[int](r)
		assert.Equal(t, []int{40, 20, 50, 45, 46, 30}, utils.SliceFrom[int](l))
	})
}
