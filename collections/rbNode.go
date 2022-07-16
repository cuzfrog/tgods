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
		compRes := comp(d, ni.v)
		if compRes < 0 {
			if ni.a == nil {
				ni.a = newRbNode(d, ni)
				nn = ni.a
				break
			} else {
				ni = ni.a
			}
		} else if compRes > 0 {
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

// delete removes a node with given value d returns:
//  r the new root
//  found true if there's a node deleted
//  nd the removed node
func delete[T any](n *rbNode[T], d T, comp types.Compare[T]) (r *rbNode[T], found bool, nd *rbNode[T]) {
	nd = n
	for true {
		if n == nil {
			return nil, false, nil
		}
		compRes := comp(nd.v, d)
		if compRes < 0 {
			nd = n.a
		} else if compRes > 0 {
			nd = n.b
		} else {
			found = true

			break
		}
	}
	return //TODO
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

func swapDown[T any](n *rbNode[T]) *rbNode[T] {
	ns := n
	for !ns.isLeaf() {
		if ns.b != nil { // TODO, here we prefer to to right branch, but we can use flag to save tree leaning info to decide which branch to go
			ns = swapInorderSuccessor(ns)
		} else {
			ns = swapInorderPredecessor(ns)
		}
	}
	return ns
}

func swapInorderSuccessor[T any](n *rbNode[T]) *rbNode[T] {
	if n.b == nil {
		return n
	}
	ns := n.b
	for ns.a != nil {
		ns = ns.a
	}
	n.v, ns.v = ns.v, n.v
	return ns
}

func swapInorderPredecessor[T any](n *rbNode[T]) *rbNode[T] {
	if n.a == nil {
		return n
	}
	ns := n.a
	for ns.b != nil {
		ns = ns.b
	}
	n.v, ns.v = ns.v, n.v
	return ns
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

func (n *rbNode[T]) isLeaf() bool {
	if n == nil {
		return true
	}
	return n.a == nil && n.b == nil
}
