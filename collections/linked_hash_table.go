package collections

type linkedHashTable[T any] struct {
	*hashTable[T]
	head *dlNode[T] // new element is added to
	tail *dlNode[T]
}

func (h *linkedHashTable[T]) Add(elem T) bool {
	h.add()
}

func (h *linkedHashTable[T]) Contains(elem T) bool {
	//TODO implement me
	panic("implement me")
}

func (h *linkedHashTable[T]) Remove(elem T) bool {
	//TODO implement me
	panic("implement me")
}

