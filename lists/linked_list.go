package lists

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

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList[T]) String() *string {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Contains(elem T) bool {
	//TODO implement me
	panic("implement me")
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

func (l *LinkedList[T]) Get(index int) (T, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Put(index int, elem T) (T, bool) {
	//TODO implement me
	panic("implement me")
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

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0}
}
