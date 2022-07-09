[![CI](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml/badge.svg)](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/cuzfrog/tgods/branch/master/graph/badge.svg?token=XIEG8JLDDW)](https://codecov.io/gh/cuzfrog/tgods)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c1532de0f9ff4fcd9f2ec7b63792b37d)](https://www.codacy.com/gh/cuzfrog/tgods/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuzfrog/tgods&amp;utm_campaign=Badge_Grade)
# Typesafe Go Data Structures

Your brand-new Golang collections implementation with generics.

Go version >= 1.18, currently depends on `golang.org/x/exp` for `Ordered` type constraints.

* arrayStack - fixed length first-in-last-out array based stack.
* circularArrayList - variable length/cap list with fast add/remove at head or tail and random access with index.
* linkedList - doubly linked list with fast add/remove.
* binaryHeap - heap based min or max priority queue.
* rbTree (under implementation..) - red black tree with no recursion but with a parent pointer in the node

more...

## Implementations

All implementations implement `Collection`
```go
type Collection[T any] interface {
	Size() int
	Contains(elem T) bool
	Iterator() Iterator[T]
	Clear()
}
```

| Implementation\Interface | Stack              | List               | ArrayList          | Queue              | Deque              |
|--------------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| arrayStack               | :heavy_check_mark: |                    |                    |                    |                    |
| circularArrayList        | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| linkedList               | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: | :heavy_check_mark: |
| heapPriorityQueue        |                    |                    |                    | :heavy_check_mark: |                    |

## Author

Cause Chung(cuzfrog@gmail.com)
