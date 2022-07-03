package core

type Collection interface {
	Size() int
	Clear()
	String() *string
}

type Iterator[T any] interface {
	Next() (T, error)
}

type Bag[T any] interface {
	Collection
	Add(elem T)
	Pop() (T, error)
	Contains(elem T) bool
}

type Queue[T any] interface {
	Bag[T]
	Head() (T, error)
	Tail() (T, error)
}

type Deque[T any] interface {
	Queue[T]
	AddTail(elem T)
	PopHead() (T, error)
}

type Stack[T any] interface {
	Bag[T]
}

type List[T any] interface {
	Deque[T]
	Get(index int) (T, error)
	Put(index int, elem T) (T, error)
}

type Map[K any, V any] interface {
	Collection
	Get(k K) V
	Put(k K, v V) V
	Remove(k K) V
	Values() Iterator[V]
	Keys() Iterator[K]
}
