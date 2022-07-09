package trees

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var comp = utils.CompareOrdered[int]

func Test_insert_root(t *testing.T) {
	r, found, nn := insert(nil, 3, comp)
	assert.False(t, found)
	assert.Equal(t, rbNode[int]{3, nil, nil, nil, true}, *r)
	assert.Equal(t, r, nn)
}

func Test_insert(t *testing.T) {
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

	r, found, n20 := insert(n70, 20, comp)
	assert.True(t, found)
	assert.Equal(t, n70, r)
	assert.Equal(t, n20, n20)

	r, found, n10 := insert(n70, 10, comp)
	assert.Equal(t, n10, n20.a)
	assert.Equal(t, n20, n10.p)
	assert.Equal(t, 10, n10.v)
	/*         70
	    	30
		20      50
	   10     40 60
	*/
	r, found, n15 := insert(n20, 15, comp)
	assert.Equal(t, n15, n10.b)
	assert.Equal(t, n10, n15.p)

	r, found, n45 := insert(n30, 45, comp)
	assert.Equal(t, n45, n40.b)
	assert.Equal(t, n40, n45.p)

	r, found, n35 := insert(n70, 35, comp)
	assert.Equal(t, n40, n35.p)
	assert.Equal(t, n35, n40.a)
}

func Test_rotate(t *testing.T) {
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
	/*        7
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

	/*       7
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