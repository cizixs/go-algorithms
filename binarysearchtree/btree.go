package binarysearchtree

import (
	"errors"
)

// Node represents a tree node
// TODO: add support for any type, not just int
type Node struct {
	Value       int   // Value the node stores
	left, right *Node // pointers to the left and right children of the key
}

// Insert will add a new node to the tree whose root is Node
func (n *Node) Insert(newNode *Node) {
	// New node goes to the right part
	if n.Value < newNode.Value {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.Insert(newNode)
		}
	}

	// New node goes to the left part
	if n.Value > newNode.Value {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.Insert(newNode)
		}
	}
}

// FindMin return the smallest element in tree
func (n *Node) FindMin() (int, error) {
	if n == nil {
		return 0, errors.New("Empty tree has no Min element.")
	}
	for n.left != nil {
		n = n.left
	}
	return n.Value, nil
}

// FindMax returns the biggest element in tree
func (n *Node) FindMax() (int, error) {
	if n == nil {
		return 0, errors.New("Empty tree has no Max element.")
	}
	for n.right != nil {
		n = n.right
	}
	return n.Value, nil
}

// Delete removes a given element
func (n *Node) Delete(element int) (*Node, bool) {
	deleted := false

	if n == nil {
		return n, false
	}

	if n.Value < element {
		n.right, deleted = n.right.Delete(element)
	} else if n.Value > element {
		n.left, deleted = n.left.Delete(element)
	} else if n.left != nil && n.right != nil {
		candidate, _ := n.right.FindMin()
		n.Value = candidate
		n.right, deleted = n.right.Delete(candidate)
		deleted = true
	} else {
		if n.left == nil {
			n = n.right
		} else {
			n = n.left
		}
		deleted = true
	}
	return n, deleted
}

// Walk function calls `f` on every node value
// FIXME: Which is a better as parameter of `f`, `Node` or `node.Value`?
func (n *Node) Walk(f func(int)) {
	if n == nil {
		return
	}

	n.left.Walk(f)
	f(n.Value)
	n.right.Walk(f)
}

// BinarySearchTree store a root and the tree node numbers.
type BinarySearchTree struct {
	root    *Node // pointer to root of the tree
	nodeNum int   // nodes number: how many nodes are in the tree
}

// New returns an initial new tree
func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Nodes returns how many nodes are in the tree
func (b *BinarySearchTree) Nodes() int {
	return b.nodeNum
}

// Insert an element into the tree
func (b *BinarySearchTree) Insert(element int) {
	node := Node{Value: element}

	if b.root == nil {
		b.root = &node
	} else {
		b.root.Insert(&node)
	}
	b.nodeNum++
}

// Walk calls `f` on every node
func (b *BinarySearchTree) Walk(f func(int)) {
	b.root.Walk(f)
}

// Contains checks if an element exists in a tree
func (b *BinarySearchTree) Contains(element int) bool {
	n := b.root
	for n != nil {
		if n.Value == element {
			return true
		}

		if n.Value < element {
			n = n.right
		} else {
			n = n.left
		}
	}

	return false
}

// IsEmpty returns if the tree is an empty tree
func (b *BinarySearchTree) IsEmpty() bool {
	return b.nodeNum == 0
}

// Find and return a pointer to the node whose value equals to `element`
func (b *BinarySearchTree) Find(element int) *Node {
	n := b.root

	for n != nil {
		if n.Value == element {
			return n
		}

		if n.Value < element {
			n = n.right
		} else {
			n = n.left
		}
	}

	return nil
}

// FindMin returns the smallest element in tree
func (b *BinarySearchTree) FindMin() (int, error) {
	return b.root.FindMin()
}

// FindMax returns the biggest element in tree
func (b *BinarySearchTree) FindMax() (int, error) {
	return b.root.FindMax()
}

// Delete the first appearence of the element
// There are several cases:
// - Node is a leaf, simply delete it
// - Node has one child, point parent to the child, and delete the node
// - Node has two children, replace the leftmost node of the right part with current node,
//   then delete the leftmost node
func (b *BinarySearchTree) Delete(element int) {
	n, deleted := b.root.Delete(element)
	b.root = n
	if deleted {
		b.nodeNum--
	}
}
