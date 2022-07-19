package collections

import "github.com/cuzfrog/tgods/types"

type hashTable[T any] struct {
	arr types.ArrayList[bucket[T]]
}
