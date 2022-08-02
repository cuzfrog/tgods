package collections

type role int

const (
	list role = iota
	stack
	queue
	deque
)

const left, right = true, false

type AccessOrder byte

const PutOrder, GetOrder AccessOrder = 1, 2
