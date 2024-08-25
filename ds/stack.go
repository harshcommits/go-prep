package ds

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() T {

	var popped T

	if len(s.Items) > 0 {
		lastIndex := len(s.Items) - 1
		popped = s.Items[lastIndex]
		s.Items = s.Items[:lastIndex]
	}

	return popped
}
