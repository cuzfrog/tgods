package lists

import "errors"

type node[T any] struct {
	v    T
	prev *node[T]
	next *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

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
func (l *LinkedList[T]) Pop() (elem T, err error) {
	if l.size == 0 {
		err = errors.New("no elem")
		return
	}
	elem = l.tail.v
	l.tail = l.tail.prev
	if l.size == 1 {
		l.head = nil
	}
	l.size--
	return
}

func (l *LinkedList[T]) Head() T {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Tail() T {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) AddTail(elem T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) PopHead() T {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Get(index int) T {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Put(index int) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0}
}
