package collections

import (
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
)

type bucket[T any] node[T]

func newSlBucketOf[T any](v T) bucket[T] {
	return newSlNode[T](v, nil)
}

//
//func (n *slNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
//	newNodeOf := func(elem T) node[T] { return newSlNode(elem, nil) }
//	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
//}
//
//func (n *slxNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
//	newNodeOf := func(elem T) node[T] { return newSlxNode(elem, nil, nil) }
//	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
//}
//
//func (n *dlNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
//	newNodeOf := func(elem T) node[T] { return newDlNode(elem, nil, nil) }
//	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
//}

//func (n *dlxNode[T]) Save(elem T, eq types.Equal[T]) (bucket[T], node[T], T, bool) {
//	newNodeOf := func(elem T) node[T] { return newDlxNode(elem, nil, nil, nil) }
//	return saveElemIntoBucket[T](elem, n, eq, newNodeOf)
//}

//saveElemIntoBucket puts the elem into the bucket, returns:
//  bucket - the either this or a changed bucket for performance based on size of the elem
//  node - the current node saving this elem
//  T - the old elem
//  bool - if found an existing elem by the eq func
func saveElemIntoBucket[T any](b bucket[T], elem T, eq types.Equal[T], newNodeOf func(elem T) node[T]) (bucket[T], node[T], T, bool) {
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
func removeElemFromBucket[T any](b bucket[T], elem T, eq types.Equal[T]) (bucket[T], T, bool) {
	if eq(elem, b.Value()) {
		v := b.Value()
		return nil, v, true
	}
	h := b
	for b.Next() != nil {
		v := b.Next().Value()
		if eq(elem, v) {
			b.SetNext(b.Next().Next())
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
