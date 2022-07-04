package core

type Collection interface {
	Size() int
	Clear()
}

type Iterator[T any] interface {
	// Next checks if there's next elem, and move iterator state to next
	Next() bool
	Index() int
	Value() T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type IndexAccess[T any] interface {
	Get(index int) (T, bool)
	Put(index int, elem T) (T, bool)
}

type Bag[T comparable] interface {
	Collection
	Iterable[T]
	Add(elem T)
	Pop() (T, bool)
	Peek() (T, bool)
	Contains(elem T) bool
}

type Queue[T comparable] interface {
	Bag[T]
	Head() (T, bool)
	Tail() (T, bool)
}

type Deque[T comparable] interface {
	Queue[T]
	AddHead(elem T)
	PopHead() (T, bool)
}

type Stack[T comparable] interface {
	Bag[T]
}

type ArrayList[T comparable] interface {
	List[T]
	IndexAccess[T]
}

type List[T comparable] interface {
	Deque[T]
}

type Map[K comparable, V any] interface {
	Collection
	Get(k K) V
	Put(k K, v V) V
	Remove(k K) V
	Values() Iterable[V]
	Keys() Iterable[K]
}
