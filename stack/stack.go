package stack

import (
	"errors"
)

type Stack struct {
	data []interface{}
	len  int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Length() int {
	return s.len
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Peek() interface{} {
	return s.data[s.len-1]
}

func (s *Stack) Push(element interface{}) {
	s.len += 1
	s.data = append(s.data, element)
}

func (s *Stack) Pop() (interface{}, error) {

	if s.IsEmpty() {
		return nil, errors.New("Pop an empty stack.")
	}
	element := s.Peek()
	s.data = s.data[:s.len-1]
	s.len -= 1
	return element, nil
}
