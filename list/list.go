/* List is a simple yet often used data structure.

l := list.New()   // create a new empty list
l.PushBack(1, 2, 3) // append one or multiple value to list
l.Length()           // return the length of list
l.Index(index)    // get the value at index
l.Find(value)     // get the index of the first appearence of value
l.Lpop()          // remove and return the first value in a list
l.Rpop()          // remove and return the last value in a list
*/

package list

import (
	"errors"
	"fmt"
)

type Node struct {
	Value      interface{} // List can store any type of value
	next, prev *Node       // double linked list
}

// List represent a serial of values in sequence
type List struct {
	root   Node // sentinal node
	length int  // how many nodea are in the list
}

// Return a new empty list
func New() *List {
	l := &List{}
	l.length = 0

	// create a sentinal node, makes it easy to deal with empty list
	l.root.next = &l.root
	l.root.prev = &l.root

	return l
}

// Check whether the list is empty
func (l *List) IsEmpty() bool {
	return l.root.next == &l.root
}

// Return how many nodes are in the list
func (l *List) Length() int {
	return l.length
}

// Add values at the front
func (l *List) PushFront(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = l.root.next
		n.prev = &l.root
		l.root.next.prev = n
		l.root.next = n
		l.length++
	}
}

// Append values at the end
func (l *List) PushBack(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = &l.root     // since n is the last element, its next should be the head
		n.prev = l.root.prev // n's prev should be the tail
		l.root.prev.next = n // tail's next should be n
		l.root.prev = n      // head's prev should be n
		l.length++
	}
}

// Find the element in list, return the index if found, otherwise return -1
func (l *List) Find(element interface{}) int {
	index := 0
	p := l.root.next
	for p != &l.root && p.Value != element {
		p = p.next
		index++
	}

	if p != &l.root {
		return index
	}

	return -1
}

func (l *List) indexFrontwise(index int) *Node {
	pos := 0
	p := l.root.next

	for p != &l.root && pos < index {
		p = p.next
		pos++
	}

	if p == &l.root {
		return nil
	} else {
		return p
	}
}

func (l *List) indexBackwise(index int) *Node {
	pos := 1
	p := l.root.prev

	for p != &l.root && pos < index {
		p = p.prev
		pos++
	}

	if p == &l.root {
		return nil
	} else {
		return p
	}
}

func (l *List) index(index int) *Node {
	var n *Node
	if index >= 0 {
		n = l.indexFrontwise(index)
	} else {
		n = l.indexBackwise(-index)
	}
	return n
}

// Return the element at index, if element is not valid, return error
// Support negative index, like -1, -2 etc, it will count backwise though.
func (l *List) Index(index int) (interface{}, error) {
	n := l.index(index)
	if n == nil {
		return nil, errors.New(fmt.Sprintf("Index %d is not valid.", index))
	}
	return n.Value, nil
}

// remove element in list
func (l *List) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	l.length--
}

// Remove and get the first element in the list
// or nil if the
func (l *List) Lpop() interface{} {
	if l.length == 0 {
		return nil
	}

	n := l.root.next
	l.remove(n)
	return n.Value
}

// Remove and get the last element in the list
func (l *List) Rpop() interface{} {
	if l.length == 0 {
		return nil
	}

	n := l.root.prev
	l.remove(n)
	return n.Value
}

// Given a index of list, return the normal index between 0 and len-1
func (l *List) normalIndex(index int) int {
	if index > l.length-1 {
		index = l.length - 1
	}

	if index < -l.length {
		index = 0
	}

	index = (l.length + index) % l.length
	return index
}

// Return a slice containing elements in range
// end can be negative, like -1, -2 etc
// if *start* is large than *end*, return empty slice `[]'
// if the end is large than the actual end, it will be treated like the last element
func (l *List) Range(start, end int) []interface{} {
	// When start or end exceeds list length
	start = l.normalIndex(start)
	end = l.normalIndex(end)
	res := []interface{}{}
	if start > end {
		return res
	}

	sNode := l.index(start)
	eNode := l.index(end)
	for n := sNode; n != eNode; {
		res = append(res, n.Value)
		n = n.next
	}
	res = append(res, eNode.Value)
	return res
}
