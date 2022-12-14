package utils

import "github.com/cuzfrog/tgods/types"

// AddAll adds all elems from src collection into target collection, return number of elems actually added
//
// Deprecated: please use AddAllTo, will be removed in future releases
func AddAll[T any](src types.Collection[T], target types.Collection[T]) int {
	return AddAllTo[T](src, target)
}

// AddAllTo adds all elems from src collection into target collection, return number of elems actually added
func AddAllTo[T any](src types.Collection[T], target types.Collection[T]) int {
	cnt := 0
	itor := src.Iterator()
	for itor.Next() {
		added := target.Add(itor.Value())
		if added {
			cnt++
		}
	}
	return cnt
}
