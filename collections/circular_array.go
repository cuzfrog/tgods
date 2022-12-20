package collections

import (
	"fmt"
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
	"github.com/cuzfrog/tgods/utils"
	"sort"
)

const defaultArrInitSize = 12
const defaultArrShrinkThreshold = 3 // bitwise shift

type circularArray[T any] struct {
	start          int //inclusive
	end            int //exclusive
	arr            []T
	size           int
	eq             types.Equal[T]
	r              role
	autoSizingFlag AutoSizingFlag
}
type circularArrayForSort[T any] struct {
	l    *circularArray[T]
	less types.Less[T]
}

// newCircularArrayOf creates an auto expandable circular array based list, auto shrinkable, but will not shrink if the length is <= defaultArrInitSize,
// the underlying array will be lazily created unless init values are provided, the init arr size is the same as init values'
func newCircularArrayOf[T comparable](values ...T) *circularArray[T] {
	return newCircularArrayOfP[T](funcs.ValueEqual[T], AutoExpand+AutoShrink, values...)
}

func newCircularArrayOfP[T any](eq types.Equal[T], flag AutoSizingFlag, values ...T) *circularArray[T] {
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
	return &circularArray[T]{start, size, arr, size, eq, list, flag}
}

// newCircularArray creates underlying array eagerly with the init cap
func newCircularArray[T comparable](initCap int, flag AutoSizingFlag) *circularArray[T] {
	return &circularArray[T]{-1, 0, make([]T, initCap), 0, funcs.ValueEqual[T], list, flag}
}

// newCircularArrayOfEq creates underlying array eagerly with the init cap, AutoExpand + AutoShrink
func newCircularArrayOfEq[T any](initCap int, eq types.Equal[T]) *circularArray[T] {
	return &circularArray[T]{-1, 0, make([]T, initCap), 0, eq, list, AutoExpand + AutoShrink}
}

// newCircularArrayOfEqP creates underlying array eagerly with the init cap
func newCircularArrayOfEqP[T any](initCap int, eq types.Equal[T], autoSizingFlag AutoSizingFlag) *circularArray[T] {
	return &circularArray[T]{-1, 0, make([]T, initCap), 0, eq, list, autoSizingFlag}
}

func (l *circularArray[T]) withRole(r role) *circularArray[T] {
	l.r = r
	return l
}

func (l *circularArray[T]) Size() int {
	return l.size
}

// Clear drop the underlying arr, O(1)
func (l *circularArray[T]) Clear() {
	l.arr = nil
	l.start = -1
	l.end = 0
	l.size = 0
}

// Add appends to the tail of the list, same as AddTail
func (l *circularArray[T]) Add(elem T) bool {
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
func (l *circularArray[T]) AddTail(elem T) bool {
	return l.Add(elem)
}

func (l *circularArray[T]) Push(elem T) bool {
	return l.Add(elem)
}

func (l *circularArray[T]) EnqueueLast(elem T) bool {
	return l.Add(elem)
}

// RemoveTail removes the last elem of the list, same as Remove
func (l *circularArray[T]) RemoveTail() (elem T, found bool) {
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
func (l *circularArray[T]) Remove() (elem T, found bool) {
	return l.RemoveTail()
}

func (l *circularArray[T]) Pop() (elem T, found bool) {
	return l.RemoveTail()
}

func (l *circularArray[T]) Peek() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	idx := l.end - 1
	if idx < 0 {
		idx = len(l.arr) - 1
	}
	return l.arr[idx], true
}

func (l *circularArray[T]) Dequeue() (elem T, found bool) {
	return l.RemoveTail()
}

// Contains checks if elem exists, O(n)
func (l *circularArray[T]) Contains(elem T) bool {
	it := l.Iterator()
	for it.Next() {
		v := it.Value()
		if l.eq(v, elem) {
			return true
		}
	}
	return false
}

// Head retrieves the first elem
func (l *circularArray[T]) Head() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	return l.arr[l.start], true
}

func (l *circularArray[T]) First() (elem T, found bool) {
	return l.Head()
}

// Tail retrieves the last elem, same as Peek
func (l *circularArray[T]) Tail() (T, bool) {
	return l.Peek()
}

// AddHead prepends to the list, see Add
func (l *circularArray[T]) AddHead(elem T) bool {
	l.expandIfNeeded()
	l.start--
	if l.start < 0 {
		l.start = len(l.arr) - 1
	}
	l.arr[l.start] = elem
	l.size++
	return true
}

func (l *circularArray[T]) Enqueue(elem T) bool {
	return l.AddHead(elem)
}

func (l *circularArray[T]) RemoveHead() (elem T, found bool) {
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

func (l *circularArray[T]) DequeueFirst() (elem T, found bool) {
	return l.RemoveHead()
}

func (l *circularArray[T]) Get(index int) (elem T, found bool) {
	arrIndex, ok := l.toArrIndex(index)
	if ok {
		return l.arr[arrIndex], true
	}
	return elem, false
}

func (l *circularArray[T]) MustGet(index int) T {
	elem, found := l.Get(index)
	if !found {
		panic(fmt.Sprintf("index %d is out of current size %d", index, l.size))
	}
	return elem
}

// Set sets value by index and returns the old value, will not expand the list
func (l *circularArray[T]) Set(index int, elem T) (oldElem T, found bool) {
	if index >= cap(l.arr) {
		return utils.Nil[T](), false
	}
	size := l.size
	if index >= size {
		for i := 0; i <= index-size; i++ {
			l.Add(utils.Nil[T]())
		}
	}
	arrIndex, _ := l.toArrIndex(index)
	oldElem = l.arr[arrIndex]
	l.arr[arrIndex] = elem
	return oldElem, true
}

// Swap exchanges values of provided indices, if one of the indices is invalid, returns false
func (l *circularArray[T]) Swap(indexA, indexB int) bool {
	arrIndexA, okA := l.toArrIndex(indexA)
	arrIndexB, okB := l.toArrIndex(indexB)
	if !okA || !okB {
		return false
	}
	l.arr[arrIndexA], l.arr[arrIndexB] = l.arr[arrIndexB], l.arr[arrIndexA]
	return true
}

func (l *circularArray[T]) clone() *circularArray[T] {
	arr := make([]T, l.size)
	copy(arr, l.arr)
	return &circularArray[T]{l.start, l.end, arr, l.size, l.eq, l.r, l.autoSizingFlag}
}

func (l *circularArray[T]) toArrIndex(index int) (int, bool) {
	if index >= l.size || index < 0 {
		return -1, false
	}
	arrIndex := l.start + index
	length := len(l.arr)
	if arrIndex >= length {
		arrIndex -= length
	}
	return arrIndex, true
}

func (l *circularArray[T]) expandIfNeeded() {
	c := cap(l.arr)
	if l.autoSizingFlag&AutoExpand == 0 {
		if l.size >= c {
			panic(fmt.Sprintf("AutoExpand disabled but current cap %d cannot contain size increment", c))
		}
		return
	}
	if c == 0 {
		l.arr = make([]T, defaultArrInitSize)
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

func (l *circularArray[T]) shrinkIfNeeded() {
	if l.arr == nil || l.autoSizingFlag&AutoShrink == 0 {
		return
	}
	newLength := len(l.arr) >> defaultArrShrinkThreshold
	if newLength <= l.size || newLength <= defaultArrInitSize {
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

func (l *circularArray[T]) Sort(less types.Less[T]) {
	c := &circularArrayForSort[T]{l, less}
	sort.Sort(c)
}

func (c *circularArrayForSort[T]) Len() int {
	return c.l.Size()
}

func (c *circularArrayForSort[T]) Less(i, j int) bool {
	elem1, found1 := c.l.Get(i)
	elem2, found2 := c.l.Get(j)
	if !found1 || !found2 {
		panic(fmt.Sprintf("invalid index %d, %d", i, j))
	}
	return c.less(elem1, elem2)
}

func (c *circularArrayForSort[T]) Swap(i, j int) {
	swapped := c.l.Swap(i, j)
	if !swapped {
		panic(fmt.Sprintf("invalid index %d, %d", i, j))
	}
}
