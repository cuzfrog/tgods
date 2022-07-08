package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate(t *testing.T) {
	n := &node[int]{3, nil, nil, true}
	assert.Equal(t, n, rotateLeft(n))
	assert.Equal(t, n, rotateRight(n))

	/*
		    3r
		2b     5r
		      4b 6b
	*/
	n = &node[int]{
		3,
		&node[int]{2, nil, nil, false},
		&node[int]{
			5,
			&node[int]{4, nil, nil, false},
			&node[int]{6, nil, nil, false},
			true,
		},
		true,
	}
	/*
		    5r
		  3r   6b
		2b  4b
	*/
	r := rotateLeft(n)
	assert.Equal(t, 5, r.v)
	assert.Equal(t, red, r.c)
	assert.Equal(t, 3, r.a.v)
	assert.Equal(t, red, r.a.c)
	assert.Equal(t, 2, r.a.a.v)
	assert.Equal(t, black, r.a.a.c)
	assert.Nil(t, r.a.a.a)
	assert.Nil(t, r.a.a.b)
	assert.Equal(t, 6, r.b.v)
	assert.Equal(t, black, r.b.c)
	assert.Nil(t, r.b.b)
	assert.Equal(t, 4, r.a.b.v)
	assert.Equal(t, black, r.a.b.c)
	assert.Nil(t, r.a.b.a)
	assert.Nil(t, r.a.b.b)

	/*
		    3r
		2b     5r
		      4b 6b
	*/
	r = rotateRight(r)
	assert.Equal(t, 3, r.v)
	assert.Equal(t, red, r.c)
	assert.Equal(t, 2, r.a.v)
	assert.Equal(t, black, r.a.c)
	assert.Nil(t, r.a.a)
	assert.Nil(t, r.a.b)
	assert.Equal(t, 5, r.b.v)
	assert.Equal(t, 4, r.b.a.v)
	assert.Equal(t, black, r.b.a.c)
	assert.Nil(t, r.b.a.a)
	assert.Nil(t, r.b.a.b)
	assert.Equal(t, 6, r.b.b.v)
	assert.Equal(t, black, r.b.b.c)
	assert.Nil(t, r.b.b.a)
	assert.Nil(t, r.b.b.b)
}
