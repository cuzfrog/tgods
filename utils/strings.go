package utils

import (
	"fmt"
	"github.com/cuzfrog/tgods/types"
	"strings"
)

func StringFrom[T any](b types.Collection[T]) *string {
	return StringFromf(b, func(elem T) string { return fmt.Sprintf("%v", elem) })
}

func StringFromf[T any](b types.Collection[T], toStringFn func(elem T) string) *string {
	str := "["
	var values []string
	iter := b.Iterator()
	for iter.Next() {
		v := iter.Value()
		values = append(values, toStringFn(v))
	}
	str += strings.Join(values, ", ")
	str += "]"
	return &str
}
