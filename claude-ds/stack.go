// ----------------------------------------------------------------------------
// Stack — LIFO, slice-backed. Zero value ready to use.
// ----------------------------------------------------------------------------

package claudeds

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Len() int { return len(s.data) }

func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	n := len(s.data)
	if n == 0 {
		return zero, false
	}
	v := s.data[n-1]
	s.data[n-1] = zero // prevent GC from retaining popped references
	s.data = s.data[:n-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}
