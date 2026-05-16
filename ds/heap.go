package ds

type Heap[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		less: less,
	}
}

func (h *Heap[T]) Len() int {
	return len(h.items)
}

func (h *Heap[T]) Peek() (T, bool) {
	var zero T
	if len(h.items) == 0 {
		return zero, false
	} else {
		return h.items[0], true
	}
}

func (h *Heap[T]) Push(value T) {
	h.items = append(h.items, value)
	h.bubbleUp(len(h.items) - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	var zero T
	n := len(h.items)

	if n == 0 {
		return zero, false
	}

	top := h.items[0]
	last := n - 1
	h.items[0] = h.items[last]

	var empty T
	h.items[last] = empty
	h.items = h.items[:last]
	if len(h.items) > 0 {
		h.bubbleDown(0)
	}

	return top, true
}

func (h *Heap[T]) bubbleUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.less(h.items[i], h.items[parent]) {
			return
		}
		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		i = parent
	}
}

func (h *Heap[T]) bubbleDown(i int) {
	n := len(h.items)

	for {
		left := 2*i + 1
		right := 2*i + 2
		best := i

		if left < n && h.less(h.items[left], h.items[best]) {
			best = left
		}
		if right < n && h.less(h.items[right], h.items[best]) {
			best = right
		}
		if best == i {
			return
		}
		h.items[i], h.items[best] = h.items[best], h.items[i]
		i = best
	}
}
