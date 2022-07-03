package core

type Collection interface {
	Size() int
	Clear()
	String() *string
}

type Iterator[T any] interface {
	// HasNext checks if there's next elem
	HasNext() bool
	// Next returns index, value and move iterator state to next
	Next() (int, T)
}

type IndexAccess[T any] interface {
	Get(index int) (T, bool)
	Put(index int, elem T) (T, bool)
}

type Bag[T any] interface {
	Collection
	Add(elem T)
	Pop() (T, bool)
	Peek() (T, bool)
	Contains(elem T) bool
}

type Queue[T any] interface {
	Bag[T]
	Head() (T, bool)
	Tail() (T, bool)
}

type Deque[T any] interface {
	Queue[T]
	AddTail(elem T)
	PopHead() (T, bool)
}

type Stack[T any] interface {
	Bag[T]
}

type List[T any] interface {
	Deque[T]
}

type Map[K any, V any] interface {
	Collection
	Get(k K) V
	Put(k K, v V) V
	Remove(k K) V
	Values() Iterator[V]
	Keys() Iterator[K]
}
