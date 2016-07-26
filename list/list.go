/* List is a simple yet often used data structure.

l := list.New()   // create a new empty list
l.Append(1, 2, 3) // append one or multiple value to list
l.Length()           // return the length of list
l.Index(index)    // get the value at index
l.Find(value)     // get the index of the first appearence of value
l.Remove(value)   // remove the first appearence of value
l.InsertBefore(value, mark)  // insert a value after mark
l.InsertAfter(value, mark)  // insert a value before mark
l.pop()          // return the value in a list
*/

package list

type Node struct {
	Value      interface{} // List can store any type of value
	next, prev *Node       // double linked list
}

// List represent a serial of values in sequence
type List struct {
	head   *Node // head point to the the first node.
	tail   *Node // tail points to the last node
	length int   // how many nodea are in the list
}

// Return a new empty list
func New() *List {
	l := &List{}
	l.length = 0

	// create a sentinal node, makes it easy to deal with empty list
	n := &Node{}
	l.head = n
	l.tail = n
	return l
}

// Check whether the list is empty
func (l *List) IsEmpty() bool {
	return l.head.next == nil
}

// Return how many nodes are in the list
func (l *List) Length() int {
	return l.length
}

// FIXME(cizixs): What is element is not comparable
func (l *List) Contains(element interface{}) bool {
	p := l.head.next
	for p != nil {
		if p.Value == element {
			return true
		} else {
			p = p.next
		}
	}
	return false
}

// Append a value at the end
func (l *List) Append(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = nil
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

// Find the element in list, return the index if found, otherwise return -1
func (l *List) Find(element interface{}) int {
	index := 0
	p := l.head.next
	for p != nil && p.Value != element {
		p = p.next
		index++
	}

	if p != nil {
		return index
	}

	return -1
}
