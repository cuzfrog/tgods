[![CI](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml/badge.svg)](https://github.com/cuzfrog/tgods/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/cuzfrog/tgods/branch/master/graph/badge.svg?token=XIEG8JLDDW)](https://codecov.io/gh/cuzfrog/tgods)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c1532de0f9ff4fcd9f2ec7b63792b37d)](https://www.codacy.com/gh/cuzfrog/tgods/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuzfrog/tgods&amp;utm_campaign=Badge_Grade)

# Typesafe Go Data Structures

Your brand-new Golang collections implementation with generics.
Go version >= [1.18](https://tip.golang.org/doc/go1.18). If facing compiler issue, please upgrade to [1.19](https://tip.golang.org/doc/go1.19).

```bash
go get github.com/cuzfrog/tgods@<version>
```

### Interfaces

| Implementation\Interface | Stack              | List               | Queue              | Deque              | Set                | Map                | MultiMap           | Graph              |
|--------------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
| `arrayStack`             | :heavy_check_mark: |                    |                    |                    |                    |                    |                    |                    |
| `circularArray`          | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |                    |                    |                    |
| `linkedList`             | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |                    |                    |                    |                    |
| `binaryHeap`             |                    |                    | :heavy_check_mark: |                    |                    |                    |                    |                    |
| `rbTree`                 |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `hashTable`              |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `linkedHashTable`        |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    |                    |
| `enumMap`                |                    |                    |                    |                    |                    | :heavy_check_mark: |                    |                    |
| `enumSet`                |                    |                    |                    |                    | :heavy_check_mark: |                    |                    |                    |
| `multiMap`               |                    |                    |                    |                    |                    | :heavy_check_mark: | :heavy_check_mark: |                    |
| `treeAdjacencyList`      |                    |                    |                    |                    |                    |                    |                    | :heavy_check_mark: |

Top interface `Collection` contains general methods, sub-interfaces like `ArrayList`, `SortedSet`, `SortedMap` provide more rich functionalities.
All interface definitions can be found: [here](./types/collection.go)

### Data Structures:

* `arrayStack` - fixed length first-in-last-out array based stack. Backing up `Stack`
* `circularArray` - variable length/cap array with fast add/remove at head or tail and random access with index. Backing up `Stack`, `ArrayList`, `Queue`, `Deque`
* `linkedList` - doubly linked list with fast add/remove. Backing up `Stack`, `LinkedList`, `Queue`, `Deque`
* `binaryHeap` - binary heap based min or max priority queue. Backing up `Queue`
* `rbTree` - recursion-free red black tree implementation. Backing up `SortedSet`, `SortedMap`
* `hashTable` - variable length/cap array based hash table, hash collision is handled by linked nodes. Backing up `Set`, `Map`
* `linkedHashTable` hashTable preserving inserting or configurable access order. Can serve as an `LRU cache`. Backing up `LinkedSet`, `LinkedMap`
* `enumMap` & `enumSet` fast array based map and set with `Integer` as the key. Implementing `SortedMap`, `SortedSet` respectively.
* `multiMap` multimap implementation with `arrayListMultiMap` and `hashSetMultiMap` constructors.
* `treeAdjacencyList` a treeMap based graph implementation with directional edge properties.
It has typical _O(log(n))_ time complexity for adding, searching, and removing vertices. Backing up `Graph`

## Usage:

### Rich Constructors:

```go
import "github.com/cuzfrog/tgods/collections"

list := collections.NewLinkedListOf(1, 2, 3) // List[int]
list := collections.NewArrayListOf(1, 2, 3) // List[int]
queue := collections.NewLinkedListQueue[int]() // Queue[int]
queue := collections.NewArrayListQueue[int]() // Queue[int]
hashMap := collections.NewHashMapOf[string, int](funcs.NewStrHash(), funcs.ValueEqual[string]) // Map[string, int]
hashMap := collections.NewHashMapOfStrKey[int](EntryOf("Earth", 3), EntryOf("Mars", 4)) // Map[string, int]
cache := collections.NewLRUCacheOfStrKey[int](40, collections.PutOrder + collections.GetOrder) // LRU cache[string, int] of size limit 40
// more...
```
Client implementation of [constraint interfaces](types/interfaces.go) to simplify collection creation:
```go
type myStruct struct {...}
func (s *myStruct) Hash() uint {...}
func (s *myStruct) Equal(other *myStruct) bool {...}
hashSet := collections.NewHashSetC[*myStruct]()
```

### Functional Transformation

```go
import "github.com/cuzfrog/tgods/transform"

c := collections.NewArrayListOf(1, 3, 4)
l := collections.NewLinkedListOf[string]()
transform.MapTo[int, string](c, l, func(elem int) string { return fmt.Sprint(elem) }) // l ["1", "2", "3"]

listOfList := ... // List[List[int]] [[1, 3], [2, 4]]
list := NewArrayListOf[int]()
transform.FlattenTo[types.List[int], int](listOfList, list) // list [1, 3, 2, 4]

c := collections.NewArrayListOf(1, 3, 4)
transform.Reduce[int, string](c, "", func(acc string, next int) string { return acc + strconv.Itoa(next) }) // "134"

c := []int{1, 3, 4}
transform.CountSlice[int](c, func(elem int) bool {return elem > 2}) // 2, yeah for slices, also has Count for Collection
// more...
```

### Built-in Utils & Functions:

```go
import "github.com/cuzfrog/tgods/utils"
list := collections.NewArrayListOf(1, 3, 2) // ArrayList 1, 3, 2
list.Sort(funcs.ValueLess[int]) // ArrayList 1, 2, 3
utils.StringFrom(list) // [1, 2, 3]
utils.SliceFrom(list)  // []int{1, 2, 3}
m := // Map [1->"a", 2->"b"]
utils.KeysFrom(m) // [1, 2]

// functions:
import "github.com/cuzfrog/tgods/funcs"
funcs.NewStrHash() // func(s string) uint, create a string hash function based on 64bit or 32bit platform
funcs.ValueEqual // (a, b comparable) bool

// HOFs:
funcs.CompToEq // func(comp Compare[T]) Equal[T]
funcs.InverseComp // func(comp Compare[T]) Compare[T]

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

## Development
Run test:
```cmd
go test --tags=test ./...
```

## Author

Cause Chung(cuzfrog@gmail.com)
