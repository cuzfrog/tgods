package lists

import (
	"github.com/cuzfrog/tgods/core"
)

// assert LinkedList implementation
var _ core.List[int] = (*LinkedList[int])(nil)

type node[T comparable] struct {
	v    T
	prev *node[T]
	next *node[T]
}

type LinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

// Contains checks if elem is present, O(n)
func (l *LinkedList[T]) Contains(elem T) bool {
	iter := l.Iterator()
	for iter.HasNext() {
		_, v := iter.Next()
		if elem == v {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) Head() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.head.v
	return elem, true
}

func (l *LinkedList[T]) Tail() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.tail.v
	return elem, true
}

// Peek same as Tail
func (l *LinkedList[T]) Peek() (elem T, found bool) {
	return l.Tail()
}

// AddTail same as Add
func (l *LinkedList[T]) AddTail(elem T) {
	l.Add(elem)
}

// PopHead removes elem from the head
func (l *LinkedList[T]) PopHead() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.head.v
	l.head = l.head.next
	if l.size == 1 {
		l.tail = nil
	}
	l.size--
	return elem, true
}

// Add adds elem to the tail
func (l *LinkedList[T]) Add(elem T) {
	prevTail := l.tail
	l.tail = &node[T]{elem, prevTail, nil}
	if l.size == 0 {
		l.head = l.tail
	} else {
		prevTail.next = l.tail
	}
	l.size++
}

// Pop gets and removes the last elem
func (l *LinkedList[T]) Pop() (elem T, found bool) {
	if l.size == 0 {
		return elem, false
	}
	elem = l.tail.v
	l.tail = l.tail.prev
	if l.size == 1 {
		l.head = nil
	}
	l.size--
	return elem, true
}

type iterator[T comparable] struct {
	index    int
	cur      *node[T]
	nextTurn bool
}

func (l *LinkedList[T]) Iterator() core.Iterator[T] {
	return &iterator[T]{0, l.head, false}
}

func (iter *iterator[T]) HasNext() bool {
	if iter.nextTurn {
		panic("can only call HasNext once before calling Next")
	}
	iter.nextTurn = true
	return iter.cur != nil
}

func (iter *iterator[T]) Next() (index int, elem T) {
	if iter.cur == nil {
		panic("cannot move to next, guard Next by checking HasNext")
	}
	index, elem = iter.index, iter.cur.v
	iter.cur = iter.cur.next
	iter.index++
	iter.nextTurn = false
	return
}

func NewLinkedList[T comparable](values ...T) *LinkedList[T] {
	l := &LinkedList[T]{nil, nil, 0}
	length := len(values)
	if length == 0 {
		return l
	}
	first := values[0]
	l.head = &node[T]{first, nil, nil}
	l.tail = l.head
	l.size = 1
	for i := 1; i < length; i++ {
		n := &node[T]{values[i], l.tail, nil}
		l.tail.next = n
		l.tail = n
		l.size++
	}
	return l
}
