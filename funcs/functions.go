package funcs

import "github.com/cuzfrog/tgods/types"

func forEach[T any](c types.Collection[T], fn func(index int, v T)) {
	it := c.Iterator()
	for it.Next() {
		fn(it.Index(), it.Value())
	}
}
