package ds

type Queue[T any] struct {
	Entries []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.Entries = append(q.Entries, value)
}

func (q *Queue[T]) Dequeue() T {
	value := q.Entries[0]
	q.Entries = q.Entries[1:]

	return value
}
