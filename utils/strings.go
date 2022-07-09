package utils

import (
	"fmt"
	"github.com/cuzfrog/tgods/core"
	"strings"
)

func StringFrom[T any](b core.Collection[T]) *string {
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
