[![CI](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml/badge.svg)](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/cuzfrog/tgods/branch/master/graph/badge.svg?token=XIEG8JLDDW)](https://codecov.io/gh/cuzfrog/tgods)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c1532de0f9ff4fcd9f2ec7b63792b37d)](https://www.codacy.com/gh/cuzfrog/tgods/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuzfrog/tgods&amp;utm_campaign=Badge_Grade)

# Typesafe Go Data Structures

Your brand-new Golang collections implementation with generics. Go version >= [1.18](https://tip.golang.org/doc/go1.18).

### Data Structures:

* `arrayStack` - fixed length first-in-last-out array based stack.
* `circularArray` - variable length/cap array with fast add/remove at head or tail and random access with index.
* `linkedList` - doubly linked list with fast add/remove.
* `binaryHeap` - binary heap based min or max priority queue.
* `rbTree` - red black tree implementation with no recursion.

### Interfaces

| Implementation\Interface | Stack              | List               | ArrayList          | Queue              | Deque              | Set                | SortedSet          |
|--------------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| `arrayStack`             | :heavy_check_mark: |                    |                    |                    |                    |                    |                    |
| `circularArray`          | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `linkedList`             | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `binaryHeap`             |                    |                    |                    | :heavy_check_mark: |                    |                    |                    |
| `treeSet`                |                    |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |

All implement `Collection`, all interface methods can be found: [here](./types/collection.go)

```go
type Collection[T any] interface {
    Add(elem T) bool
    Contains(elem T) bool
    Each(func (index int, elem T))
    Iterator() Iterator[T]
    Size() int
    Clear()
}
```

### Constructors:

```go
import "github.com/cuzfrog/tgods/collections"

list := collections.NewLinkedListOf(1, 2, 3) // List[int]
list := collections.NewCircularArrayListOf(1, 2, 3) // List[int]
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
