package types

type Collection[T any] interface {
	Add(elem T) bool      // add elem into the collection, return true if succeeded
	Contains(elem T) bool // typical time complexity O(n) in array based, O(log(n)) in tree based, and O(c) in hash based implementations.
	Iterator() Iterator[T]
	Each(func(index int, elem T))
	Size() int
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
}

type LinkedList[T any] interface {
	List[T]
}

type List[T any] interface {
	Collection[T]
	AddHead(elem T) bool
	RemoveHead() (T, bool) // removes the first elem of the list
	Head() (T, bool)
	AddTail(elem T) bool
	RemoveTail() (T, bool)
	Remove() (T, bool) // alias for RemoveTail
	Tail() (T, bool)
}

type Set[T any] interface {
	Collection[T]
	Remove(elem T) bool
}

// SortedSet The elem order is decided by Compare func.
type SortedSet[T any] interface {
	Set[T]
	First() (T, bool)                // returns the first elem and true if it has, or Nil and false if it hasn't. The elem is not removed from the set.
	Last() (T, bool)                 // returns the last elem and true if it has, or Nil and false if it hasn't. The elem is not removed from the set.
	RemoveFirst() (T, bool)          // returns the first elem and true if it has, or Nil and false if it hasn't.
	RemoveLast() (T, bool)           // returns the last elem and true if it has, or Nil and false if it hasn't.
	HeadSet(toElem T) SortedSet[T]   // returns a new set contains elements from first to toElem(exclusive).
	TailSet(fromElem T) SortedSet[T] // returns a new set contains elements from fromElem(inclusive) to the last.
	Higher(elem T) (T, bool)         // returns the least elem that is greater than the given elem, or Nil and false if no one can be found.
	Lower(elem T) (T, bool)          // returns the greatest elem that is less than the given elem, or Nil and false if no one can be found.
	Ceiling(elem T) (T, bool)        // returns the least elem that is greater than or equal to the given elem, or Nil and false if no one can be found.
	Floor(elem T) (T, bool)          // returns the greatest elem that is less than or equal to the given elem, or Nil and false if no one can be found.
	ReverseSet() SortedSet[T]        // returns a view of the same set with an inverted Compare func thus reverted element order.
}

type Map[K any, V any] interface {
	Get(k K) V
	Put(k K, v V) V
	Remove(k K) V
}
