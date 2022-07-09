package trees

import "github.com/cuzfrog/tgods/utils"

const red, black = true, false

// assert rbTree implementation
var _ RbTree[int] = (*rbTree[int])(nil)

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
	r, found := insert(t.root, d, t.comp)
	t.root = r
	if !found {
		t.size++
	}
	return found
}

func newRbNode[T any](d T, p *rbNode[T]) *rbNode[T] {
	return &rbNode[T]{d, nil, nil, p, red}
}

// insert returns found or created rbNode
//   np - n's parent
//   ngp - n's grandparent
func insert[T any](n *rbNode[T], d T, comp utils.Compare[T]) (*rbNode[T], bool) {
	if n == nil {
		return newRbNode(d, nil), false
	}
	var found bool
	if comp(d, n.v) < 0 {
		n.a, found = insert(n.a, d, comp)
	} else if comp(d, n.v) > 0 {
		n.b, found = insert(n.b, d, comp)
	} else {
		n.v = d
		found = true
	}
	n = rectify(n)
	return n, found
}

// rectify recolors and/or rotates when necessary
func rectify[T any](n *rbNode[T]) *rbNode[T] {
	if n.p == nil {
		n.c = black
		return n
	}
	if n.p.c == black {
		return n
	}
	// when np is red, ngp must exist
	ngp := n.p.p
	nu := n.uncle()
	if nu.color() == red {
		n.p.setColor(black)
		nu.setColor(black)
		n.p.p.setColor(red)
	} else {
		if n == n.p.a { // left child
			rotateRight(ngp)
		} else { // right child

		}
	}
	return n
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

func rotateLeft[T any](n *rbNode[T]) {
	if n.b == nil {
		return
	}
	r := n.b
	n.b = r.a
	if n.b != nil {
		n.b.p = n
	}
	r.a = n
	updateParentChild(n, r)
}

func rotateRight[T any](n *rbNode[T]) {
	if n.a == nil {
		return
	}
	r := n.a
	n.a = r.b
	if n.a != nil {
		n.a.p = n
	}
	r.b = n
	updateParentChild(n, r)
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
