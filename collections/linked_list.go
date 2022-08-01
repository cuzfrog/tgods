package collections

import (
	"github.com/cuzfrog/tgods/funcs"
	"github.com/cuzfrog/tgods/types"
)

type linkedList[T any] struct {
	head *dlNode[T]
	tail *dlNode[T]
	size int
	comp types.Equal[T]
	r    role
}

func newLinkedListOf[T comparable](values ...T) *linkedList[T] {
	return newLinkedListOfEq(funcs.ValueEqual[T], values...)
}

func newLinkedListOfEq[T any](eq types.Equal[T], values ...T) *linkedList[T] {
	l := &linkedList[T]{nil, nil, 0, eq, list}
	length := len(values)
	if length == 0 {
		return l
	}
	first := values[0]
	l.head = &dlNode[T]{first, nil, nil}
	l.tail = l.head
	l.size = 1
	for i := 1; i < length; i++ {
		n := &dlNode[T]{values[i], l.tail, nil}
		l.tail.next = n
		l.tail = n
		l.size++
	}
	return l
}

func (l *linkedList[T]) withRole(r role) *linkedList[T] {
	l.r = r
	return l
}

func (l *linkedList[T]) Size() int {
	return l.size
}

func (l *linkedList[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

// Contains checks if elem is present, O(n)
func (l *linkedList[T]) Contains(elem T) bool {
	iter := l.Iterator()
	for iter.Next() {
		v := iter.Value()
		if l.comp(v, elem) { // TODO, check for optimize
			return true
		}
	}
	return false
}

func (l *linkedList[T]) Head() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.head.v
	return elem, true
}

func (l *linkedList[T]) First() (elem T, found bool) {
	return l.Head()
}

func (l *linkedList[T]) Tail() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.tail.v
	return elem, true
}

// Peek same as Tail
func (l *linkedList[T]) Peek() (elem T, found bool) {
	return l.Tail()
}

// AddHead prepends to the list
func (l *linkedList[T]) AddHead(elem T) bool {
	prevHead := l.head
	l.head = &dlNode[T]{elem, nil, prevHead}
	if l.size == 0 {
		l.tail = l.head
	} else {
		prevHead.prev = l.head
	}
	l.size++
	return true
}

func (l *linkedList[T]) Enqueue(elem T) bool {
	return l.Add(elem)
}

func (l *linkedList[T]) Push(elem T) bool {
	return l.AddHead(elem)
}

func (l *linkedList[T]) RemoveHead() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.head.v
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	if l.size == 1 {
		l.tail = nil
	}
	l.size--
	return elem, true
}

func (l *linkedList[T]) Pop() (elem T, found bool) {
	return l.RemoveHead()
}

func (l *linkedList[T]) DequeueFirst() (elem T, found bool) {
	return l.Remove()
}

// Add adds elem to the tail
func (l *linkedList[T]) Add(elem T) bool {
	prevTail := l.tail
	l.tail = &dlNode[T]{elem, prevTail, nil}
	if l.size == 0 {
		l.head = l.tail
	} else {
		prevTail.next = l.tail
	}
	l.size++
	return true
}

func (l *linkedList[T]) AddTail(elem T) bool {
	return l.Add(elem)
}

func (l *linkedList[T]) EnqueueLast(elem T) bool {
	return l.AddHead(elem)
}

// Remove gets and removes the last elem
func (l *linkedList[T]) Remove() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.tail.v
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	if l.size == 1 {
		l.head = nil
	}
	l.size--
	return elem, true
}

func (l *linkedList[T]) RemoveTail() (elem T, found bool) {
	return l.Remove()
}

func (l *linkedList[T]) Dequeue() (elem T, found bool) {
	return l.RemoveHead()
}
