package trees

import (
	"github.com/cuzfrog/tgods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var comp = utils.CompareOrdered[int]

func Test_insert(t *testing.T) {
	n, found := insert(nil, 3, comp)
	assert.False(t, found)
	assert.Equal(t, rbNode[int]{3, nil, nil, nil, true}, *n)

	/*
		    30b
		20b     50r
		      40b 60b
	*/
	//var newNodes = func() *rbNode[int] {
	//	return &rbNode[int]{
	//		30,
	//		&rbNode[int]{20, nil, nil, false},
	//		&rbNode[int]{
	//			50,
	//			&rbNode[int]{40, nil, nil, false},
	//			&rbNode[int]{60, nil, nil, false},
	//			true,
	//		},
	//		false,
	//	}
	//}
	//n, found = insert(newNodes(), 20, comp)
	//assert.True(t, found)
	//assert.Equal(t, *newNodes(), *n)
	//
	//n, found = insert(newNodes(), 10, comp)
	//assert.Equal(t, 10, n.a.a.v)
}

func Test_rotate(t *testing.T) {
	n := &rbNode[int]{3, nil, nil, nil, true}
	rotateLeft(n)
	assert.Equal(t, newRbNode(3, nil), n)
	rotateRight(n)
	assert.Equal(t, newRbNode(3, nil), n)

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
	rotateLeft(n3)
	r := n5
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
	rotateRight(n5)
	r = n3
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
