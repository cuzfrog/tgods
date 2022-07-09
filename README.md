[![CI](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml/badge.svg)](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/cuzfrog/tgods/branch/master/graph/badge.svg?token=XIEG8JLDDW)](https://codecov.io/gh/cuzfrog/tgods)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c1532de0f9ff4fcd9f2ec7b63792b37d)](https://www.codacy.com/gh/cuzfrog/tgods/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuzfrog/tgods&amp;utm_campaign=Badge_Grade)
# Typesafe Go Data Structures

Your brand-new Golang collections implementation with generics. Currently, under implementation...

Go version >= 1.18, currently depends on `golang.org/x/exp` for `Ordered` type constraints.

* ArrayStack - fixed length first-in-last-out array based stack.
* CircularArrayList - variable length/cap list with fast add/remove at head or tail and random access with index.
* LinkedList - doubly linked list with fast add/remove.
* HeapPriorityQueue - binary heap based min or max priority queue.
* RbTree - red black tree with no recursion but with a parent pointer in the node

more...

## Interfaces

All implementation implements `Collection`
```go
type Collection interface {
	Size() int
	Clear()
}
```

| Method\Interface                   | Time* | Bag                | Stack              | List               | ArrayList          | Queue              | Deque              |
|------------------------------------|-------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| `Get(index int) (T, bool)`         | O(c)  |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `Set(index int, elem T) (T, bool)` | O(c)  |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `Swap(indexA, indexB int) bool`    | O(c)  |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `Add(elem T) bool`                 | O(c)  | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Pop() (T, bool)`                  | O(c)  | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Peek() (T, bool)`                 | O(c)  | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Contains(elem T) bool`            | O(n)  | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Head() (T, bool)`                 | O(c)  |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: |
| `Tail() (T, bool)`                 | O(c)  |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `AddHead(elem T) bool`             | O(c)  |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: |
| `PopHead() (T, bool)`              | O(c)  |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: |

* Time* - typical time
## Implementations

| Implementation\Interface | Bag                | Stack              | List               | ArrayList          | Queue              | Deque              |
|--------------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| ArrayStack               | :heavy_check_mark: | :heavy_check_mark: |                    |                    |                    |                    |
| CircularArrayList        | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| LinkedList               | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: | :heavy_check_mark: |
| HeapPriorityQueue        | :heavy_check_mark: |                    |                    |                    | :heavy_check_mark: |                    |

## Author

Cause Chung(cuzfrog@gmail.com)