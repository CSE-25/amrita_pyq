package stack

// Stack implementation using Go slices.
type Stack struct {
	items []string
}

// Push adds an item to the stack.
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Peek returns the top item without removing it.
func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// NewStack creates a new stack.
func NewStack() *Stack {
	return &Stack{}
}

var STACK = NewStack()
