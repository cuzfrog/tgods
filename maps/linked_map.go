package maps

import (
	"container/list"
	"github.com/cuzfrog/ggods/core"
)

type LinkedMap struct {
	list list.List
}

func New[K any, V any]() core.Map[K, V] {
	return nil
}
