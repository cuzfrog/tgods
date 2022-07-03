package generic_collections

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Bag[T any] interface {
	Add(elem T)
	Pop() T
	Size() int
}

type Queue[T any] interface {
	Add(elem T)
	Pop() T
	Head() T
	Tail() T
	Size() int
}

type Deque[T any] interface {
	Queue[T]
	AddTail(elem T)
	PopHead() T
}

type Stack[T any] interface {
	Bag[T]
}

type List[T any] interface {
	Deque[T]
	Get(index int) T
	Put(index int) (T, bool)
}

type Map[K any, V any] interface {
	Get(k K) V
	Put(k K, v V) V
	Remove(k K) V
	Size() int
	Values() Iterator[V]
	Keys() Iterator[K]
}
