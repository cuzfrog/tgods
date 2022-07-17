package collections

type role int

const (
	list role = iota
	stack
	queue
	deque
)

const left, right = true, false
