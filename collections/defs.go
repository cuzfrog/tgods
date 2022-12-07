package collections

type role int

const (
	list role = iota
	stack
	queue
	deque
)

const left, right = true, false

// AccessOrder 1 - put order, 2 - get order, 3 - both get and put order, 0 - original put order (no explicit order change)
type AccessOrder byte

const (
	OriginalOrder AccessOrder = 0 // original put order (no explicit order change)
	PutOrder      AccessOrder = 1 // newly put element will be at the tail (the elem at head will be next for eviction).
	GetOrder      AccessOrder = 2 // newly get element will be at the tail (the elem at head will be next for eviction).
)
