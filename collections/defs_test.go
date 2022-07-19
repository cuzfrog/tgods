package collections

import (
	"github.com/cuzfrog/tgods/funcs"
)

type intStruct struct {
	v int
}

var compInt = funcs.ValueCompare[int]
var eqInt = funcs.ValueEqual[int]
