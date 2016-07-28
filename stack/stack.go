package stack

import (
	"errors"
)

// Stack is a FILO data structure.
// Stack supports two major methods: push and pop.
type Stack struct {
	data []interface{}
	len  int
}

// New returns an initial empty stack
func New() *Stack {
	return &Stack{}
}

// Length return how many elements are in a list
func (s *Stack) Length() int {
	return s.len
}

// IsEmpty returns if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

// Peek return the top-most element in stack, or nil if stack is empty
func (s *Stack) Peek() interface{} {
	return s.data[s.len-1]
}

// Push adds an element to stack
func (s *Stack) Push(element interface{}) {
	s.len++
	s.data = append(s.data, element)
}

// Pop removes and returns the toppest element in stack
func (s *Stack) Pop() (interface{}, error) {

	if s.IsEmpty() {
		return nil, errors.New("Pop an empty stack.")
	}
	element := s.Peek()
	s.data = s.data[:s.len-1]
	s.len--
	return element, nil
}
