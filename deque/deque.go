/*Package deque provides a useful interface of a list
that can be add/remove elements from both end.

    d := deque.New()
    d.Push(3)
    d.Length()
    d.Push(4)
    data := d.pop()
    d.Inject("world")
    d.Eject()
*/
package deque

import (
	"fmt"
)

type node struct {
	value      interface{}
	prev, next *node
}

// Deque is the main struct
type Deque struct {
	head, tail *node
	length     uint
}

// New creates and returns a deque
func New() *Deque {
	d := &Deque{}
	// create sentinal nodes
	hn := &node{}
	d.head = hn
	d.tail = hn

	return d
}

// Length returns how many elements are in deque
func (d *Deque) Length() uint {
	return d.length
}

// Push insert an item at the front of deque
func (d *Deque) Push(value interface{}) error {
	n := &node{}
	n.value = value
	n.next = d.head.next
	n.prev = d.head
	d.head.next = n

	if d.length == 0 {
		d.tail = n
	}
	d.length++

	return nil
}

// Pop removes and returns the element from the front of deque
func (d *Deque) Pop() (interface{}, error) {
	if d.length == 0 {
		return nil, fmt.Errorf("Can not pop from an empty deque.")
	}

	item := d.head.next
	d.head.next = d.head.next.next

	if d.head.next != nil {
		d.head.next.prev = d.head
	}

	d.length--

	return item.value, nil
}

// Inject adds an item at the end of deque
func (d *Deque) Inject(value interface{}) error {
	n := &node{}
	n.value = value

	d.tail.next = n
	n.prev = d.tail
	d.tail = n

	d.length++
	return nil
}

// Eject removes and returns the last item from deque
func (d *Deque) Eject() (interface{}, error) {
	if d.length == 0 {
		return nil, fmt.Errorf("Can not eject from an empty deque.")
	}

	item := d.tail
	d.tail.prev.next = d.tail.next
	d.tail = d.tail.prev

	return item.value, nil
}
