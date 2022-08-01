[![CI](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml/badge.svg)](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/cuzfrog/tgods/branch/master/graph/badge.svg?token=XIEG8JLDDW)](https://codecov.io/gh/cuzfrog/tgods)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c1532de0f9ff4fcd9f2ec7b63792b37d)](https://www.codacy.com/gh/cuzfrog/tgods/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuzfrog/tgods&amp;utm_campaign=Badge_Grade)

# Typesafe Go Data Structures

Your brand-new Golang collections implementation with generics. Go version >= [1.18](https://tip.golang.org/doc/go1.18).

### Interfaces

| Implementation\Interface | Stack              | List               | Queue              | Deque              | Set                | Map                |
|--------------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| `arrayStack`             | :heavy_check_mark: |                    |                    |                    |                    |                    |
| `circularArray`          | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `linkedList`             | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `binaryHeap`             |                    |                    | :heavy_check_mark: |                    |                    |                    |
| `rbTree`                 |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |
| `hashTable`              |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |
| `linkedHashTable`        |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |

Top interface `Collection` contains general methods, sub-interfaces like `ArrayList`, `SortedSet` provide more rich functionalities.
All interface definitions can be found: [here](./types/collection.go)

### Data Structures:

* `arrayStack` - fixed length first-in-last-out array based stack. Backing up `Stack`
* `circularArray` - variable length/cap array with fast add/remove at head or tail and random access with index. Backing up `Stack`, `ArrayList`, `Queue`, `Deque`
* `linkedList` - doubly linked list with fast add/remove. Backing up `Stack`, `LinkedList`, `Queue`, `Deque`
* `binaryHeap` - binary heap based min or max priority queue. Backing up `Queue`
* `rbTree` - recursion-free red black tree implementation. Backing up `SortedSet`, `SortedMap`
* `hashTable` - variable length/cap array based hash table. Backing up `Set`, `Map`
* `linkedHashTable` hashTable preserving inserting order. Can serve as an `LRU cache`. Backing up `Set`, `Map`

## Usage:

### Constructors:

```go
import "github.com/cuzfrog/tgods/collections"

list := collections.NewLinkedListOf(1, 2, 3) // List[int]
list := collections.NewArrayListOf(1, 2, 3) // List[int]
queue := collections.NewLinkedListQueue[int]() // Queue[int]
queue := collections.NewArrayListQueue[int]() // Queue[int]
// more...
```

### Functional Transformation

```go
import "github.com/cuzfrog/tgods/transform"

c := collections.NewCircularArrayList(1, 3, 4)
l := collections.NewLinkedList[string]()
transform.MapTo[int, string](c, l, func(elem int) string { return fmt.Sprint(elem) }) // l ["1", "2", "3"]
```

### Utils:

```go
import "github.com/cuzfrog/tgods/utils"
list := collections.NewLinkedListOf(1, 2, 3)
utils.StringFrom(list) // [1, 2, 3]
utils.SliceFrom(list)  // []int{1, 2, 3}
funcs.NewStrHash() // func(s string) uint
// more...
```

### Semantic Iteration:

Iteration behavior is based on which interface type is constructed.

```go
list := collections.NewLinkedListOf[int]() // interface List[int], implementation linkedList[int]
list.Add(1)
list.Add(2)
list.Add(3)
list.Each(func (i, v int) {fmt.Print(v)}) // 123

stack := collections.NewLinkedListStack[int]() // interface Stack[int], implementation linkedList[int]
stack.Push(1)
stack.Push(2)
stack.Push(3)
stack.Each(func (i, v int) {fmt.Print(v)}) // 321
```

## Author

Cause Chung(cuzfrog@gmail.com)
