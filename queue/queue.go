/* Queue is a FIFO data strcuture

The basic operations on a queue are `Enqueue`, which inserts an
element at the end, and `Dequeue`, which deletes an element at
the start.

*/
package queue

import (
	"errors"
)

// Queue declaration
type Queue struct {
	data []interface{}    // use slice to store queue data
	len int               // queue size
}

// Create a new queue with nil value
func New() *Queue {
	return &Queue{}
}

// Return how many elements are in queue
func (q *Queue) Length() int{
	return q.len
}

// Check if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

// Push an element into queue to the end
func (q *Queue) Push(element interface{}){
	q.data = append(q.data, element)
	q.len += 1
}

// Remove an element from queue from the start and return it
func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty(){
		return nil, errors.New("Can not pop from an empty queue.")
	}

	item := q.data[0]
	q.data = q.data[1:]
	q.len -= 1
	return item, nil
}

// Peek return
func (q *Queue) Peek() interface{} {
	return q.data[0]
}
