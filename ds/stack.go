package ds

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}

	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return popped
}
