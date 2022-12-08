package collections

import (
	"github.com/cuzfrog/tgods/funcs"
)

type intStruct struct {
	v int
}

var compInt = funcs.ValueCompare[int]
var eqInt = funcs.ValueEqual[int]

func (s *intStruct) Hash() uint {
	return funcs.NumHash(s.v)
}

func (s *intStruct) Equal(other *intStruct) bool {
	return s.v == other.v
}
