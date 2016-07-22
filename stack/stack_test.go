package stack_test

import (
	"testing"

	"github.com/cizixs/go-algorithms/stack"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := stack.New()

	if s.Length() != 0 {
		t.Error("New stack should have length 0, got %d", s.Length())
	}
}

func TestPushAndPeek(t *testing.T) {
	assert := assert.New(t)

	s := stack.New()

	s.Push("a")
	assert.Equal(s.Length(), 1, "Stack length should be 1 after pushing an element, got %d", s.Length())
	assert.Equal(s.Peek(), "a", "Expect element `a`, got %v", s.Peek())

	s.Push(42)
	assert.Equal(s.Length(), 2, "Stack length should be 2 after pushing another element, got %d", s.Length())
	assert.Equal(s.Peek(), 42, "Expect element `42`, got %v", s.Peek())
}

func testPop(t *testing.T) {
	assert := assert.New(t)

	s := stack.New()

	data := []interface{}{42, "universe", 13.2}
	for _, element := range data {
		s.Push(element)
	}

	for i, _ := range data {
		element, _ := s.Pop()
		assert.Equal(element, data[2-i], "Expect %v, got %v", data[2-i], element)
	}

	// pop from the empty stack
	_, err := s.Pop()
	assert.Error(err, "Pop from an empty stack expects an error, got nil")

}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)

	s := stack.New()
	assert.True(s.IsEmpty(), "New stack should be empty")

	s.Push(42)
	assert.False(s.IsEmpty(), "Stack should not be empty.")

	s.Pop()
	assert.True(s.IsEmpty(), "New stack should be empty")

}
