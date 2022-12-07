package collections

import (
	"github.com/cuzfrog/tgods/types"
)

const defaultHashTableInitSize = 12

const defaultHashTableExpandRatio = 2 // r - bitwise shift, if ((cap(arr) - size) << r) < size then expand
const defaultHashTableShrinkRatio = 3 // r - bitwise shift, if (size << r) < cap(arr) then shrink, cap cannot be less than defaultHashTableInitSize

type hashTable[T any] struct {
	arr       []bucket[T]
	size      int
	hs        types.Hash[T]
	eq        types.Equal[T]
	newNodeOf func(elem T) node[T]
}

func newHashTable[T any](hs types.Hash[T], eq types.Equal[T]) *hashTable[T] {
	return newHashTableOfInitCap(0, hs, eq)
}

func newHashTableOfInitCap[T any](initCap int, hs types.Hash[T], eq types.Equal[T]) *hashTable[T] {
	newNodeOf := func(elem T) node[T] { return newSlNode[T](elem, nil) }
	var arr []bucket[T]
	if initCap > 0 {
		arr = make([]bucket[T], initCap)
	}
	return &hashTable[T]{arr, 0, hs, eq, newNodeOf}
}

func newHashTableOfSlxNode[T any](hs types.Hash[T], eq types.Equal[T]) *hashTable[T] {
	newNodeOf := func(elem T) node[T] { return newSlxNode[T](elem, nil, nil) }
	return &hashTable[T]{nil, 0, hs, eq, newNodeOf}
}

// Add inserts the elem and return true if succeeded
func (h *hashTable[T]) Add(elem T) bool {
	h.add(elem)
	return true
}

func (h *hashTable[T]) Replace(elem T) (T, bool) {
	_, old, found := h.add(elem)
	return old, found
}

// add inserts the elem and return:
//
//		n - the node containing the elem;
//		old - existing elem if found;
//	 found - if found an existing elem.
func (h *hashTable[T]) add(elem T) (n node[T], old T, found bool) {
	h.expandIfNeeded()
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
	if b == nil { // interface pointer receiver cannot be nil
		n = h.newNodeOf(elem)
		b = n
		found = false
	} else {
		b, n, old, found = saveElemIntoBucket(b, elem, h.eq, h.newNodeOf)
	}
	h.arr[i] = b
	if !found {
		h.size++
	}
	return
}

func (h *hashTable[T]) Contains(elem T) bool {
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
	if b == nil {
		return false
	}
	return b.Contains(elem, h.eq)
}

func (h *hashTable[T]) Size() int {
	return h.size
}

func (h *hashTable[T]) Clear() {
	h.size = 0
	h.arr = nil
}

func (h *hashTable[T]) Remove(elem T) bool {
	n := h.remove(elem)
	return n != nil
}

// remove returns old node if found
func (h *hashTable[T]) remove(elem T) node[T] {
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
	if b == nil {
		return nil
	}
	b, oldN := removeElemFromBucket[T](b, elem, h.eq)
	if oldN != nil {
		h.size--
	}
	h.arr[i] = b
	h.shrinkIfNeeded()
	return oldN
}

// hashToIndex
func hashToIndex(hash uint, cap int) int {
	return int(hash) & (cap - 1)
}

func (h *hashTable[T]) expandIfNeeded() {
	if h.arr == nil {
		h.arr = make([]bucket[T], defaultHashTableInitSize)
		return
	}
	c := cap(h.arr)
	emptyCap := c - h.size
	if emptyCap < h.size>>defaultHashTableExpandRatio {
		dc := c << 1
		if dc <= c {
			panic("Cannot expand underlying array.")
		}
		h.hashRedistribute(dc)
	}
}

func (h *hashTable[T]) shrinkIfNeeded() {
	if h.arr == nil || cap(h.arr) <= defaultHashTableInitSize<<1 {
		return
	}
	if h.size < cap(h.arr)>>defaultHashTableShrinkRatio {
		c := cap(h.arr) >> 1
		h.hashRedistribute(c)
	}
}

// hashRedistribute c - the new cap
func (h *hashTable[T]) hashRedistribute(c int) {
	a := make([]bucket[T], c)
	it := h.Iterator()
	for it.Next() {
		i := hashToIndex(h.hs(it.Value()), c)
		b := a[i]
		if b == nil {
			b = newSlBucketOf(it.Value())
		} else {
			b, _, _, _ = saveElemIntoBucket(a[i], it.Value(), h.eq, h.newNodeOf)
		}
		a[i] = b
	}
	h.arr = a
}

func (h *hashTable[T]) getNode(elem T) node[T] {
	if h.size == 0 {
		return nil
	}
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
	return findNodeFromBucket[T](b, elem, h.eq)
}
