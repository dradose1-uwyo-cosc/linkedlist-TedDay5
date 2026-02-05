package ds

type Stack struct {
	list LinkedList
}

// Push adds a new head node
func (s *Stack) Push(value string) {
	_ = s.list.InsertAt(0, value)
}

// Pop removes the head node
func (s *Stack) Pop() (string, bool) {
	if s.list.IsEmpty() {
		return "", false
	}

	value := s.list.Head.data
	_ = s.list.RemoveAt(0)
	return value, true
}
