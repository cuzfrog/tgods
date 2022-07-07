[![CI](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml/badge.svg)](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/cuzfrog/tgods/branch/master/graph/badge.svg?token=XIEG8JLDDW)](https://codecov.io/gh/cuzfrog/tgods)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c1532de0f9ff4fcd9f2ec7b63792b37d)](https://www.codacy.com/gh/cuzfrog/tgods/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuzfrog/tgods&amp;utm_campaign=Badge_Grade)
# Typesafe Go Data Structures

Your brand-new Golang collections implementation with generics.

Go version >= 1.18, currently depends on `golang.org/x/exp` for some type constraints.

## Interfaces

All implementation implements `Collection`

| Method\Interface                   | (Usual)Time | Collection         | IndexAccess        | Bag                | Stack              | List               | ArrayList          | Queue              | Deque              |
|------------------------------------|-------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| `Size() int`                       | O(1)        | :heavy_check_mark: |                    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Clear()`                          | O(1)        | :heavy_check_mark: |                    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Get(index int) (T, bool)`         | O(1)        |                    | :heavy_check_mark: |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `Set(index int, elem T) (T, bool)` | O(1)        |                    | :heavy_check_mark: |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `Swap(indexA, indexB int) bool`    | O(1)        |                    | :heavy_check_mark: |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `Add(elem T) bool`                 | O(1)        |                    |                    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Pop() (T, bool)`                  | O(1)        |                    |                    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Peek() (T, bool)`                 | O(1)        |                    |                    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Contains(elem T) bool`            | O(n)        |                    |                    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| `Head() (T, bool)`                 | O(1)        |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: |
| `Tail() (T, bool)`                 | O(1)        |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `AddHead(elem T) bool`             | O(1)        |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: |
| `PopHead() (T, bool)`              | O(1)        |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    | :heavy_check_mark: |

## Implementations

| Implementation\Interface | Iterable           | Bag                | Stack              | List               | ArrayList / IndexAccess | Queue              | Deque              |
|--------------------------|--------------------|--------------------|--------------------|--------------------|-------------------------|--------------------|--------------------|
| ArrayStack               | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |                         |                    |                    |
| CircularArrayList        | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:      | :heavy_check_mark: | :heavy_check_mark: |
| LinkedList               | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                         | :heavy_check_mark: | :heavy_check_mark: |
| HeapPriorityQueue        | :heavy_check_mark: | :heavy_check_mark: |                    |                    |                         | :heavy_check_mark: |                    |

## Author

Cause Chung(cuzfrog@gmail.com)