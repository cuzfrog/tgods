package lists

import (
	"fmt"
	"github.com/cuzfrog/tgods/core"
)

const DefaultInitSize = 12
const DefaultShrinkThreshold = 3 // bitwise shift

// assert circularArrayList implementation
var _ core.ArrayList[int] = (*circularArrayList[int])(nil)
var _ core.List[int] = (*circularArrayList[int])(nil)
var _ core.Stack[int] = (*circularArrayList[int])(nil)
var _ core.Queue[int] = (*circularArrayList[int])(nil)
var _ core.Deque[int] = (*circularArrayList[int])(nil)

type circularArrayList[T comparable] struct {
	start int //inclusive
	end   int //exclusive
	arr   []T
	size  int
}

// NewCircularArrayListOf creates an auto expandable circular array based list, auto shrinkable, but will not shrink if the length is <= DefaultInitSize,
// the underlying array will be lazily created, if init values are provided, the init arr size is the same as init values'
func NewCircularArrayListOf[T comparable](values ...T) *circularArrayList[T] {
	var arr []T
	var size, start int
	length := len(values)
	if length == 0 {
		arr = nil
		size = 0
		start = -1
	} else {
		arr = values
		size = length
		start = 0
	}
	return &circularArrayList[T]{start, size, arr, size}
}

// NewCircularArrayListWithInitSize creates underlying array eagerly with the size, see NewCircularArrayListOf
func NewCircularArrayListWithInitSize[T comparable](size int) *circularArrayList[T] {
	return &circularArrayList[T]{-1, 0, make([]T, size), 0}
}

func (l *circularArrayList[T]) Size() int {
	return l.size
}

// Clear drop the underlying arr, O(1)
func (l *circularArrayList[T]) Clear() {
	l.arr = nil
	l.start = -1
	l.end = 0
	l.size = 0
}

// Add appends to the tail of the list, same as AddTail
func (l *circularArrayList[T]) Add(elem T) bool {
	l.expandIfNeeded()
	if l.end >= len(l.arr) {
		l.end = 0
	}
	l.arr[l.end] = elem
	l.end++
	l.size++
	if l.start < 0 {
		l.start = 0
	}
	return true
}

// AddTail appends to the tail of the list, same as Add
func (l *circularArrayList[T]) AddTail(elem T) bool {
	return l.Add(elem)
}

func (l *circularArrayList[T]) Enstack(elem T) bool {
	return l.Add(elem)
}

func (l *circularArrayList[T]) EnqueueLast(elem T) bool {
	return l.Add(elem)
}

// RemoveTail removes the last elem of the list, same as Remove
func (l *circularArrayList[T]) RemoveTail() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	if l.end <= 0 {
		l.end = len(l.arr)
	}
	elem = l.arr[l.end-1]
	l.end--
	l.size--
	l.shrinkIfNeeded()
	return elem, true
}

// Remove removes the last elem of the list, same as RemoveTail
func (l *circularArrayList[T]) Remove() (elem T, found bool) {
	return l.RemoveTail()
}

func (l *circularArrayList[T]) Pop() (elem T, found bool) {
	return l.RemoveTail()
}

func (l *circularArrayList[T]) Peek() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	return l.arr[l.end-1], true
}

func (l *circularArrayList[T]) Dequeue() (elem T, found bool) {
	return l.RemoveTail()
}

// Contains checks if elem exists, O(n)
func (l *circularArrayList[T]) Contains(elem T) bool {
	it := l.Iterator()
	for it.Next() {
		v := it.Value()
		if v == elem {
			return true
		}
	}
	return false
}

// Head retrieves the first elem
func (l *circularArrayList[T]) Head() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	return l.arr[l.start], true
}

func (l *circularArrayList[T]) First() (elem T, found bool) {
	return l.Head()
}

// Tail retrieves the last elem, same as Peek
func (l *circularArrayList[T]) Tail() (T, bool) {
	return l.Peek()
}

// AddHead prepends to the list, see Add
func (l *circularArrayList[T]) AddHead(elem T) bool {
	l.expandIfNeeded()
	l.start--
	if l.start < 0 {
		l.start = len(l.arr) - 1
	}
	l.arr[l.start] = elem
	l.size++
	return true
}

func (l *circularArrayList[T]) Enqueue(elem T) bool {
	return l.AddHead(elem)
}

func (l *circularArrayList[T]) RemoveHead() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.arr[l.start]
	l.start++
	if l.start >= len(l.arr) {
		l.start = 0
	}
	l.size--
	l.shrinkIfNeeded()
	return elem, true
}

func (l *circularArrayList[T]) DequeueFirst() (elem T, found bool) {
	return l.RemoveHead()
}

func (l *circularArrayList[T]) Get(index int) (elem T, found bool) {
	arrIndex, ok := l.toArrIndex(index)
	if ok {
		return l.arr[arrIndex], true
	}
	return elem, false
}

// Set sets value by index and returns the old value, will not expand the list
func (l *circularArrayList[T]) Set(index int, elem T) (oldElem T, found bool) {
	arrIndex, ok := l.toArrIndex(index)
	if ok {
		oldElem = l.arr[arrIndex]
		l.arr[arrIndex] = elem
		return oldElem, true
	}
	return oldElem, false
}

// Swap exchanges values of provided indices, if one of the indices is invalid, returns false
func (l *circularArrayList[T]) Swap(indexA, indexB int) bool {
	arrIndexA, okA := l.toArrIndex(indexA)
	arrIndexB, okB := l.toArrIndex(indexB)
	if !okA || !okB {
		return false
	}
	l.arr[arrIndexA], l.arr[arrIndexB] = l.arr[arrIndexB], l.arr[arrIndexA]
	return true
}

func (l *circularArrayList[T]) toArrIndex(index int) (int, bool) {
	if index >= l.size {
		return -1, false
	}
	arrIndex := l.start + index
	length := len(l.arr)
	if arrIndex >= length {
		arrIndex -= length
	}
	return arrIndex, true
}

func (l *circularArrayList[T]) Iterator() core.Iterator[T] {
	return &calIterator[T]{l, -1, -1}
}

type calIterator[T comparable] struct {
	l        *circularArrayList[T]
	index    int
	arrIndex int
}

func (it *calIterator[T]) Next() bool {
	it.index++
	arrIndex, ok := it.l.toArrIndex(it.index) // TODO: optimize
	it.arrIndex = arrIndex
	return ok
}

// Index returns current index, will not fail when invalid, should be guarded by Next()
func (it *calIterator[T]) Index() int {
	return it.index
}

func (it *calIterator[T]) Value() T {
	return it.l.arr[it.arrIndex]
}

func (l *circularArrayList[T]) expandIfNeeded() {
	if l.arr == nil {
		l.arr = make([]T, DefaultInitSize)
	} else if l.size >= len(l.arr) {
		newLength := l.size << 1
		if newLength <= l.size {
			panic(fmt.Sprintf("cannot expand arr of size(%d)", l.size))
		}
		newArr := make([]T, l.size<<1)
		iter := l.Iterator() // TODO: try to optimize
		for iter.Next() {
			i, v := iter.Index(), iter.Value()
			newArr[i] = v
		}
		l.arr = newArr
		l.start = 0
		l.end = l.size
	}
}

func (l *circularArrayList[T]) shrinkIfNeeded() {
	if l.arr == nil {
		return
	}
	if l.size == 0 {
		l.arr = nil
		return
	}
	newLength := len(l.arr) >> DefaultShrinkThreshold
	if newLength <= l.size || newLength <= DefaultInitSize {
		return
	}
	newArr := make([]T, newLength)
	iter := l.Iterator() // TODO: try to optimize
	for iter.Next() {
		i, v := iter.Index(), iter.Value()
		newArr[i] = v
	}
	l.arr = newArr
	l.start = 0
	l.end = l.size
}
