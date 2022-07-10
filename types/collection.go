package types

type Collection[T any] interface {
	Size() int
	Contains(elem T) bool // typical time complexity O(n)
	Iterator() Iterator[T]
	Clear()
}

type Iterator[T any] interface {
	Next() bool // checks if there's next elem, and move iterator state to next
	Index() int // returns current index, or undefined value if Next gives false
	Value() T   // returns current value, or undefined value if Next gives false
}

type IndexAccess[T any] interface {
	Get(index int) (T, bool)
	Set(index int, elem T) (T, bool)
	Swap(indexA, indexB int) bool
}

type Queue[T any] interface {
	Collection[T]
	Enqueue(elem T) bool
	Dequeue() (T, bool)
	Peek() (T, bool) // Peek retrieves the next elem of the queue, equivalent to Dequeue without removal
}

type Deque[T any] interface {
	Stack[T]
	Queue[T]
	EnqueueLast(elem T) bool
	DequeueFirst() (T, bool)
	First() (T, bool)
}

type Stack[T any] interface {
	Collection[T]
	Push(elem T) bool
	Pop() (T, bool)
	Peek() (T, bool)
}

type ArrayList[T any] interface {
	List[T]
	IndexAccess[T]
	Cloneable[T, ArrayList[T]]
}

type List[T any] interface {
	Collection[T]
	AddHead(elem T) bool
	RemoveHead() (T, bool) // removes the first elem of the list
	Head() (T, bool)
	AddTail(elem T) bool
	RemoveTail() (T, bool)
	Add(elem T) bool   // alias for AddTail
	Remove() (T, bool) // alias for RemoveTail
	Tail() (T, bool)
}

type Map[K any, V any] interface {
	Get(k K) V
	Put(k K, v V) V
	Remove(k K) V
}
