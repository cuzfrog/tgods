package trees

import "github.com/cuzfrog/tgods/utils"

const red, black = true, false

// assert rbTree implementation
var _ Tree[int] = (*rbTree[int])(nil)

type rbTree[T any] struct {
	root *rbNode[T]
	size int
	comp utils.Compare[T]
}

type rbNode[T any] struct {
	v T
	a *rbNode[T]
	b *rbNode[T]
	p *rbNode[T]
	c bool
}

func (t *rbTree[T]) Insert(d T) bool {
	r, found, nn := insert(t.root, d, t.comp)
	t.root = r
	for true {
		nn = rebalance(nn)
		if nn == nil {
			break
		}
	}
	if !found {
		t.size++
	}
	return found
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
func insert[T any](n *rbNode[T], d T, comp utils.Compare[T]) (r *rbNode[T], found bool, nn *rbNode[T]) {
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
	if n.p == nil {
		n.c = black
		return nil
	}
	if n.p.c == black {
		return nil
	}
	// when np is red, ngp must exist
	ngp := n.p.p
	nu := n.uncle()
	if nu.color() == red {
		n.p.setColor(black)
		nu.setColor(black)
		ngp.setColor(red)
		r = ngp
	} else {
		if n == n.p.a { // left child
			r = rotateRight(ngp)
		} else { // right child
			r = rotateLeft(ngp)
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
