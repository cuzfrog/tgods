package collections

import (
	"github.com/cuzfrog/tgods/types"
)

const red, black = true, false

type rbNode[T any] struct {
	v T
	a *rbNode[T]
	b *rbNode[T]
	p *rbNode[T]
	c bool
}

func newRbNode[T any](d T, p *rbNode[T]) *rbNode[T] {
	return &rbNode[T]{d, nil, nil, p, red}
}

/*
insert returns:
	r - the top node after insertion
	found - if found an existing node
	nn - the newly created or found node
*/
func insert[T any](n *rbNode[T], d T, comp types.Compare[T]) (r *rbNode[T], found bool, nn *rbNode[T]) {
	if n == nil {
		r, found = newRbNode(d, nil), false
		nn = r
		return
	}
	ni := n
	for true {
		if comp(d, ni.v) < 0 {
			if ni.a == nil {
				ni.a = newRbNode(d, ni)
				nn = ni.a
				break
			} else {
				ni = ni.a
			}
		} else if comp(d, ni.v) > 0 {
			if ni.b == nil {
				ni.b = newRbNode(d, ni)
				nn = ni.b
				break
			} else {
				ni = ni.b
			}
		} else {
			ni.v = d
			found = true
			nn = ni
			break
		}
	}
	return n, found, nn
}

// rebalance recolors and/or rotates when necessary, returns next rectifiable node or nil if finishes
func rebalance[T any](n *rbNode[T]) (r *rbNode[T]) {
	np := n.p
	if np == nil {
		n.c = black
		return nil
	}
	if np.c == black {
		return nil
	}
	// when np is red, ngp must exist
	ngp := np.p
	nu := n.uncle()
	if nu.color() == red {
		np.setColor(black)
		nu.setColor(black)
		ngp.setColor(red)
		r = ngp
	} else {
		if n.p == ngp.a { // left branch
			if n == n.p.b {
				rotateLeft(np)
			}
			r = rotateRight(ngp)
			swapColor(r, r.b)
		} else { // right branch
			if n == n.p.a {
				rotateRight(np)
			}
			r = rotateLeft(ngp)
			swapColor(r, r.a)
		}
	}
	return r
}

func (n *rbNode[T]) uncle() *rbNode[T] {
	gp := n.p.p
	if gp == nil {
		return nil
	}
	if n.p == gp.a {
		return gp.b
	} else {
		return gp.a
	}
}

func swapColor[T any](a, b *rbNode[T]) {
	a.c, b.c = b.c, a.c
}

func rotateLeft[T any](n *rbNode[T]) (r *rbNode[T]) {
	if n.b == nil {
		return n
	}
	r = n.b
	n.b = r.a
	if n.b != nil {
		n.b.p = n
	}
	r.a = n
	updateParentChild(n, r)
	//r.c, n.c = n.c, red
	return
}

func rotateRight[T any](n *rbNode[T]) (r *rbNode[T]) {
	if n.a == nil {
		return n
	}
	r = n.a
	n.a = r.b
	if n.a != nil {
		n.a.p = n
	}
	r.b = n
	updateParentChild(n, r)
	//r.c, n.c = n.c, red
	return
}

// updateParentChild update original parent's child after rotation
//   n - the original node
//   nn - the new node rotated up
func updateParentChild[T any](n *rbNode[T], nn *rbNode[T]) {
	if n.p == nil {
		// ignore when n is the original root
	} else if n == n.p.a {
		n.p.a = nn
	} else {
		n.p.b = nn
	}
	nn.p = n.p
	n.p = nn
}

func (n *rbNode[T]) color() bool {
	if n == nil {
		return black
	}
	return n.c
}

func (n *rbNode[T]) setColor(c bool) {
	if n != nil {
		n.c = c
	}
}
