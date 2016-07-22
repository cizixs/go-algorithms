package binarysearchtree

// TODO: add support for any type, not just int
type Node struct {
	value       int   // Value the node stores
	left, right *Node // pointers to the left and right children of the key
}

func (n *Node) Insert(newNode *Node) {
	if n.value < newNode.value {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.Insert(newNode)
		}
	}

	if n.value > newNode.value {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.Insert(newNode)
		}
	}
}

func (n *Node) Walk(f func(int)) {
	if n == nil {
		return
	}

	n.left.Walk(f)
	f(n.value)
	n.right.Walk(f)
}

type BinarySearchTree struct {
	root      *Node
	nodes_num int // nodes number: how many nodes are in the tree
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (b *BinarySearchTree) Nodes() int {
	return b.nodes_num
}

func (b *BinarySearchTree) Insert(value int) {
	node := Node{value: value}

	if b.root == nil {
		b.root = &node
	} else {
		b.root.Insert(&node)
	}
	b.nodes_num += 1
}

func (b *BinarySearchTree) Walk(f func(int)) {
	b.root.Walk(f)
}

func (b *BinarySearchTree) Exists(element int) bool {
	n := b.root
	for n != nil {
		if n.value == element {
			return true
		}

		if n.value < element {
			n = n.right
		} else {
			n = n.left
		}
	}

	return false
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.nodes_num == 0
}
