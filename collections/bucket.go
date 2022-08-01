package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type bucket[T any] node[T]

func newSlBucketOf[T any](v T) bucket[T] {
	return newSlNode[T](v, nil)
}

func (n *slNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
	newNodeOf := func(elem T) node[T] { return newSlNode(elem, nil) }
	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
}

func (n *slxNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
	newNodeOf := func(elem T) node[T] { return newSlxNode(elem, nil, nil) }
	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
}

func (n *dlNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
	newNodeOf := func(elem T) node[T] { return newDlNode(elem, nil, nil) }
	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
}

func (n *dlxNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
	newNodeOf := func(elem T) node[T] { return newDlxNode(elem, nil, nil, nil) }
	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
}

func saveElemIntoBucket[T any](elem T, b bucket[T], eq types.Equal[T], newNodeOf func(elem T) node[T]) (bucket[T], node[T], T, bool) {
	h := b
	var np node[T]
	var cur node[T] = b
	for cur != nil {
		if eq(elem, cur.Value()) {
			old := cur.Value()
			cur.SetValue(elem)
			return h, cur, old, true
		}
		np = cur
		cur = cur.Next()
	}
	np.SetNext(newNodeOf(elem))
	return h, np.Next(), utils.Nil[T](), false
}

func (n *slNode[T]) Get(elem T, eq types.Equal[T]) (T, bool) {
	var next node[T] = n
	for next != nil {
		if eq(elem, next.Value()) {
			return next.Value(), true
		}
		next = next.Next()
	}
	return utils.Nil[T](), false
}

// removeElemFromBucket removes the elem from the bucket, return the elem and true if found
func removeElemFromBucket[T any](n node[T], elem T, eq types.Equal[T]) (bucket[T], T, bool) {
	if eq(elem, n.Value()) {
		v := n.Value()
		return nil, v, true
	}
	h := n
	for n.Next() != nil {
		v := n.Next().Value()
		if eq(elem, v) {
			n.SetNext(n.Next().Next())
			return h, v, true
		}
	}
	return h, utils.Nil[T](), false
}

func (n *slNode[T]) Contains(elem T, eq types.Equal[T]) bool {
	var next node[T] = n
	for next != nil {
		if eq(elem, next.Value()) {
			return true
		}
		next = next.Next()
	}
	return false
}
