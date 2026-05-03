package claudeds

// ----------------------------------------------------------------------------
// Heap — generic priority queue. Ordering is supplied by the caller so the
// same type can be used as a min-heap, max-heap, or keyed on any struct field.
// ----------------------------------------------------------------------------

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{less: less}
}

func (h *Heap[T]) Len() int { return len(h.data) }

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.siftUp(len(h.data) - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	var zero T
	n := len(h.data)
	if n == 0 {
		return zero, false
	}
	top := h.data[0]
	h.data[0] = h.data[n-1]
	h.data[n-1] = zero
	h.data = h.data[:n-1]
	if len(h.data) > 0 {
		h.siftDown(0)
	}
	return top, true
}

func (h *Heap[T]) Peek() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.less(h.data[i], h.data[parent]) {
			return
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *Heap[T]) siftDown(i int) {
	n := len(h.data)
	for {
		l, r := 2*i+1, 2*i+2
		best := i
		if l < n && h.less(h.data[l], h.data[best]) {
			best = l
		}
		if r < n && h.less(h.data[r], h.data[best]) {
			best = r
		}
		if best == i {
			return
		}
		h.data[i], h.data[best] = h.data[best], h.data[i]
		i = best
	}
}
