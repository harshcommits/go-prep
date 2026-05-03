package claudeds

// ----------------------------------------------------------------------------
// Queue — FIFO, circular buffer. Avoids the O(n) dequeue cost of a naive slice.
// ----------------------------------------------------------------------------

type Queue[T any] struct {
	data       []T
	head, tail int
	size       int
}

func (q *Queue[T]) Len() int { return q.size }

func (q *Queue[T]) Enqueue(v T) {
	if q.size == len(q.data) {
		q.grow()
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	v := q.data[q.head]
	q.data[q.head] = zero
	q.head = (q.head + 1) % len(q.data)
	q.size--
	return v, true
}

func (q *Queue[T]) grow() {
	newCap := len(q.data) * 2
	if newCap == 0 {
		newCap = 8
	}
	buf := make([]T, newCap)
	for i := 0; i < q.size; i++ {
		buf[i] = q.data[(q.head+i)%len(q.data)]
	}
	q.data = buf
	q.head = 0
	q.tail = q.size
}
