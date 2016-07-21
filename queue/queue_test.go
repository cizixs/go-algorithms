package queue_test

import (
	"testing"

	"github.com/cizixs/go-algorithms/queue"
)


func TestNew(t *testing.T){
	q := queue.New()

	if !q.IsEmpty() {
		t.Error("New Queue should be empty!")
	}

	if q.Length() != 0 {
		t.Error("New Queue should be length of 0, got %d", q.Length())
	}
}


func TestPush(t *testing.T){
	q := queue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Length() != 3{
		t.Error("Queue should be length 3, but got %d", q.Length())
	}
}

func TestPop(t *testing.T){
	q := queue.New()
	elements := []int{1, 2, 3}
	for _, elem := range elements {
		q.Push(elem)
	}

	for index, elem := range elements {
		if data, _ := q.Pop(); data != elem && q.Length() != 2-index {
			t.Error("Poped value should be %d, but got %v", elem, data)
		}
	}

	// try to pop an empty queue, expect an error
	_, err := q.Pop()
	if err == nil {
		t.Error("Pop from an empty queue, expect an error, got nil.")
	}
}

func TestPeek(t *testing.T){
	q := queue.New()
	elements := []int{1, 2, 3}
	for _, elem := range elements {
		q.Push(elem)
	}

	if q.Peek() != 1 {
		t.Error("Peek value, expect 1, got %d", q.Peek())
	}

	if q.Length() != 3 {
		t.Error("Peek should not change queue size. Expect %d, got %d", 3, q.Length())
	}
}
