package maps

import (
	"container/list"
	"github.com/cuzfrog/tgods/core"
)

type LinkedMap struct {
	list list.List
}

func New[K comparable, V any]() core.Map[K, V] {
	return nil
}
