package utils

import (
	"fmt"
	"github.com/cuzfrog/tgods/types"
	"strings"
)

func StringFrom[T any](b types.Collection[T]) *string {
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
