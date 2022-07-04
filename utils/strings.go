package utils

import (
	"fmt"
	"github.com/cuzfrog/tgods/core"
	"strings"
)

func ToString[T any](b core.Iterable[T]) *string {
	str := "["
	var values []string
	iter := b.Iterator()
	for iter.Next() {
		v := iter.Value()
		values = append(values, fmt.Sprintf("%v", v))
	}
	str += strings.Join(values, ", ")
	str += "]"
	return &str
}