package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack() //Stack content: []

	if s.IsEmpty() == false {
		t.Error("Stack is not initialized as empty")
	}

	s.Push("1") //Stack content: [1]
	s.Push("2") //Stack content: [1, 2]
	s.Push("3") //Stack content: [1, 2, 3]

	if s.IsEmpty() == true {
		t.Error("Push failed")
	}

	if s.Peek() != "3" {
		t.Error("Peek failed")
	}

	elem := s.Pop() //Stack content: [1, 2]

	if elem != "3" {
		t.Error("Pop failed")
	}

	s.Pop() //Stack content: [1]
	s.Pop() //Stack content: []

	if s.Pop() != "" {
		t.Error("Pop failed")
	}
}
