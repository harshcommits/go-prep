package ds

type Stack struct {
	Items []string
}

func (s *Stack) Push(item string) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() string {
	if len(s.Items) == 0 {
		return "-1"
	}

	lastIndex := len(s.Items) - 1
	popped := s.Items[lastIndex]
	s.Items = s.Items[:lastIndex]

	return popped
}
