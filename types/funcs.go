package types

// Compare is a func that return 0 when a == b, 1 when a > b, -1 when a < b
type Compare[T any] func(a, b T) int8

type Less[T any] func(a, b T) bool

type Equal[T any] func(a, b T) bool

type Hash[T any] func(a T) uint
