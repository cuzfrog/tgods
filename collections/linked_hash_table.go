package collections

type linkedHashTable[T any] struct {
	*hashTable[T]
	head  node[T]
	tail  node[T] // new element is added to
	limit int
}

func (h *linkedHashTable[T]) Add(elem T) bool {
	if h.tail == nil {
		h.tail = newDlNode(elem, nil, nil)
		h.head = h.tail
		x := h.tail
		n, _, _ := h.add(elem)
		n.SetExternal(x)
	} else {
		n, _, found := h.add(elem)
		if found {
			x := n.External()
			xPrev := x.Prev()
			xNext := x.Next()
			if xPrev != nil {
				xPrev.SetNext(x.Next())
			} else if xNext != nil {
				x.Next().SetPrev(xPrev)
			}
			x.SetPrev(nil)
			x.SetNext(nil)
			h.tail.SetNext(x)
		} else {
			x := newDlNode(elem, nil, nil)
			h.tail.SetNext(x)
			n.SetExternal(x)
		}
	}
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
