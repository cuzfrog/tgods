package collections

import "github.com/cuzfrog/tgods/types"

const defaultHashTableInitSize = 12

const defaultHashTableExpandRatio = 2 // r - bitwise shift, if ((cap(arr) - size) << r) < size then expand
const defaultHashTableShrinkRatio = 3 // r - bitwise shift, if (size << r) < cap(arr) then shrink, cap cannot be less than defaultHashTableInitSize

type hashTable[T any] struct {
	arr  []bucket[T]
	size int
	hs   types.Hash[T]
	eq   types.Equal[T]
}

func newHashTable[T any](hs types.Hash[T], eq types.Equal[T]) *hashTable[T] {
	return &hashTable[T]{nil, 0, hs, eq}
}

func (h *hashTable[T]) Add(elem T) (found bool) {
	h.expandIfNeeded()
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
	if b == nil {
		b = newLinkedListBucketOf[T](elem) // interface pointer receiver cannot be nil
		found = false
	} else {
		b, _, found = b.Save(elem, h.eq)
	}
	h.arr[i] = b
	if !found {
		h.size++
	}
	return true
}

func (h *hashTable[T]) Contains(elem T) bool {
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
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
	i := hashToIndex(h.hs(elem), cap(h.arr))
	b := h.arr[i]
	if b == nil {
		return false
	}
	b, _, found := b.Delete(elem, h.eq)
	if found {
		h.size--
	}
	h.arr[i] = b
	h.shrinkIfNeeded()
	return found
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
			b = newLinkedListBucketOf(it.Value())
		} else {
			b, _, _ = a[i].Save(it.Value(), h.eq)
		}
		a[i] = b
	}
	h.arr = a
}
