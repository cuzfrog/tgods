package generic_collections

import "container/list"

type LinkedMap struct {
	list list.List
}

func New[K any, V any]() Map[K, V] {

}
