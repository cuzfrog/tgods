package types

type Collection[T any] interface {
	Add(elem T) bool       // adds elem into the collection, return true if succeeded. If it's a Map, elem is an entry, a newly added entry will replace the old entry if exists.
	Contains(elem T) bool  // checks if an elem is in the collection by a given Equal function. If it's a Map, elem is an entry, but only compares the key. Typical time complexity O(n) in array based, O(log(n)) in tree based, and O(c) in hash based implementations.
	Iterator() Iterator[T] // returns a semantic iterator whose behavior is based on the sub-interface
	Each(func(index int, elem T))
	Size() int // returns element size, if it's a Map returns entry size, if it's a MultiMap returns value element size.
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
	AddHead(elem T) bool   // prepends elem to the head, return true if succeeded.
	RemoveHead() (T, bool) // removes the first elem of the list
	Head() (T, bool)
	AddTail(elem T) bool // appends elem to the tail, return true if succeeded.
	RemoveTail() (T, bool)
	Remove() (T, bool) // alias for RemoveTail
	Tail() (T, bool)
}

type Set[T any] interface {
	Collection[T]
	Remove(elem T) bool // removes elem from the set, returns true if found the elem. O(log(n))
}

// SortedSet The elem order is decided by Compare func.
type SortedSet[T any] interface {
	Set[T]
	First() (T, bool)       // returns the first elem and true if it has, or Nil and false if it hasn't. The elem is not removed from the set. O(log(n))
	Last() (T, bool)        // returns the last elem and true if it has, or Nil and false if it hasn't. The elem is not removed from the set. O(log(n))
	RemoveFirst() (T, bool) // returns and removes the first elem and true if it has, or Nil and false if it hasn't. O(log(n))
	RemoveLast() (T, bool)  // returns and removes the last elem and true if it has, or Nil and false if it hasn't. O(log(n))
	//HeadSet(toElem T) SortedSet[T]          // returns a new set contains elements from first to toElem(exclusive). O(n*log(n))
	//TailSet(fromElem T) SortedSet[T]        // returns a new set contains elements from fromElem(inclusive) to the last. O(n*log(n))
	//SubSet(fromElem, toElem T) SortedSet[T] // returns a new set contains elements from fromElem(inclusive) to toElem(exclusive). O(n*log(n))
	//Higher(elem T) (T, bool)                // returns the least elem that is greater than the given elem, or Nil and false if no one can be found. O(log(n))
	//Lower(elem T) (T, bool)                 // returns the greatest elem that is less than the given elem, or Nil and false if no one can be found. O(log(n))
	//Ceiling(elem T) (T, bool)               // returns the least elem that is greater than or equal to the given elem, or Nil and false if no one can be found. O(log(n))
	//Floor(elem T) (T, bool)                 // returns the greatest elem that is less than or equal to the given elem, or Nil and false if no one can be found. O(log(n))
	//ReverseSet() SortedSet[T]               // returns a view of the same set with an inverted Compare func and reverted element order. O(c)
}

type Entry[K any, V any] interface {
	Key() K
	Value() V
}

type Map[K any, V any] interface {
	Collection[Entry[K, V]]
	Get(k K) (V, bool)      // gets the elem by the given key, returns Nil and false if not found
	Put(k K, v V) (V, bool) // puts a key value pair, returns the old value and true associated with the key if any, or Nil and false if not exists. If it's a MultiMap, the value is a collection, which will be directly saved without copying the elements, the implementation of the collection then may violate the multimap properties.
	Remove(k K) (V, bool)   // removes the value associated with the given key, returns the moved value and true if found, or Nil and false if not exists
	ContainsKey(k K) bool
}

// SortedMap The key order is decided by Compare func.
type SortedMap[K any, V any] interface {
	Map[K, V]
	First() Entry[K, V]       // returns the first entry if it has, or nil if it's empty. The entry is not removed from the map. O(log(n))
	Last() Entry[K, V]        // returns the last entry if it has, or nil if it's empty. The entry is not removed from the map. O(log(n))
	RemoveFirst() Entry[K, V] // returns and removes the first entry if it has, or nil if it's empty. O(log(n))
	RemoveLast() Entry[K, V]  // returns and removes the last entry if it has, or nil if it's empty. O(log(n))
}

// Graph a graph implementation supporting directional edges with properties
//
//	V - the vertex type
//	E - the edge type
type Graph[V any, E any] interface {
	Collection[V]
	Connect(from, to V, edge E) (E, bool)               // connects vertices 'from' to 'to' with provided edge property, add the vertices into the graph if not already exist. Returns old edge property and true if a connection was existing, or Nil and false if not existing. If 'from' == 'to' by Compare, no edge will be added.
	Disconnect(from, to V) (E, bool)                    // disconnects vertices 'from' to 'to', returns the existing edge and true if any, or Nil and false. If 'from' == 'to' by Compare, no effect.
	InwardCount(vertex V) int                           // returns inward edge count. If the graph does not have the vertex, returns 0
	OutwardCount(vertex V) int                          // returns outward edge count. If the graph does not have the vertex, returns 0
	InwardEdges(vertex V) (Iterator[Entry[V, E]], int)  // returns inward neighbour edges and connected vertices and the count. If the graph does not have the vertex, returns an empty iterator.
	OutwardEdges(vertex V) (Iterator[Entry[V, E]], int) // returns outward neighbour edges and connected vertices and the count. If the graph does not have the vertex, returns an empty iterator.
	Remove(vertex V) bool                               // removes the vertex and all connected edges from the graph. Returns true if found the vertex in the graph, or false if not.
}

type ListMultiMap[K any, V any] interface {
	Map[K, List[V]]
	PutSingle(k K, v V)
	KeySize() int
	AllValues() List[V]
}

type SetMultiMap[K any, V any] interface {
	Map[K, Set[V]]
	PutSingle(k K, v V)
	KeySize() int
	AllValues() Set[V]
}
