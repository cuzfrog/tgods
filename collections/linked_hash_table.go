package collections

type linkedHashTable[T any] struct {
	*hashTable[T]
	head  *dlNode[T]
	tail  *dlNode[T] // new element is added to
	limit int
}

func (h *linkedHashTable[T]) Add(elem T) bool {
	if h.tail == nil {
		h.tail = &dlNode[T]{elem, nil, nil}
		h.head = h.tail
	}

	h.add(elem)
	return true
}

func (h *linkedHashTable[T]) Contains(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (h *linkedHashTable[T]) Remove(elem T) bool {
	//TODO implement me
	panic("implement me")
}
