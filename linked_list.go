package generic_collections

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (l *LinkedList[T]) Add(elem T) {
	prevTail := l.tail
	l.tail = &Node[T]{elem, prevTail, nil}
	if l.size == 0 {
		l.head = l.tail
	} else {
		prevTail.next = l.tail
	}
	l.size++
}

func (l *LinkedList[T]) Pop() T {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
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

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{nil, nil, 0}
}
