package binarysearchtree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cizixs/go-algorithms/binarysearchtree"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	b := binarysearchtree.New()
	assert.Equal(b.Nodes(), 0, "New binary tree should have 0 nodes")
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	b := binarysearchtree.New()
	data := []int{56, 37, 89, 2, 10, 72}
	for _, element := range data {
		b.Insert(element)
	}

	assert.False(b.IsEmpty(), "Tree should not be empty after inserting elements.")
	assert.Equal(b.Nodes(), len(data), "Counting tree nodes, expected %d, got %d", len(data), b.Nodes())

	for _, element := range data {
		assert.True(b.Contains(element), "Element %d should be in tree.", element)
	}
}

func TestMin(t *testing.T) {
	assert := assert.New(t)

	b := binarysearchtree.New()

	_, err := b.FindMin()
	assert.Error(err, "Find min value in empty tree should return error")

	data := []int{56, 37, 89, 2, 10, 72}
	for _, element := range data {
		b.Insert(element)
	}
	min, _ := b.FindMin()
	assert.Equal(2, min)
}

func TestMax(t *testing.T) {
	assert := assert.New(t)

	b := binarysearchtree.New()

	_, err := b.FindMax()
	assert.Error(err, "Find max value in empty tree should return error")

	data := []int{56, 37, 89, 2, 10, 72}
	for _, element := range data {
		b.Insert(element)
	}
	max, _ := b.FindMax()
	assert.Equal(89, max)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	b := binarysearchtree.New()

	data := []int{56, 37, 89, 2, 10, 72}
	for _, element := range data {
		b.Insert(element)
	}

	// Delete a one-child node
	b.Delete(37)
	assert.Equal(b.Nodes(), 5, "After deleting, tree should have 5 nodes, got %d", b.Nodes())
	assert.False(b.Contains(37), "Delete 37 does not work.")

	// Delete a not-exist node
	b.Delete(100)
	assert.False(b.Contains(100), "Delete a node not exist should do nothing")
	assert.Equal(b.Nodes(), 5, "Delete a node not exist should not change tree's size")

	// Delete a lead node
	b.Delete(72)
	assert.Equal(b.Nodes(), 4, "After deleting, tree should have 5 nodes, got %d", b.Nodes())
	assert.False(b.Contains(72), "Delete 37 does not work.")
}

func TestDeleteOneNodeTree(t *testing.T) {
	assert := assert.New(t)
	b := binarysearchtree.New()
	b.Insert(34)
	b.Delete(34)

	assert.Equal(b.Nodes(), 0, "Tree should have 0 node, got %d", b.Nodes())
	assert.False(b.Contains(34), "Tree should not contain deleted node 34.")
	assert.True(b.IsEmpty(), "Tree should be empty.")
}

func TestDeleteTwoChildNode(t *testing.T) {
	assert := assert.New(t)
	b := binarysearchtree.New()

	data := []int{56, 37, 89, 2, 10, 72, 40}
	for _, element := range data {
		b.Insert(element)
	}

	b.Delete(37)
	assert.Equal(b.Nodes(), 6, "After deleting, tree should have 5 nodes, got %d", b.Nodes())
	assert.False(b.Contains(37), "Delete 37 should remove the element from tree.")
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	b := binarysearchtree.New()
	data := []int{56, 37, 89, 2, 10, 72}
	for _, element := range data {
		b.Insert(element)
	}

	for _, element := range data {
		assert.Equal(b.Find(element).Value, element, "Find should return a node contains the element")
	}
	assert.Nil(b.Find(9999), "Find a non-exist element should return nil")
}

func TestWalk(t *testing.T) {
	assert := assert.New(t)

	b := binarysearchtree.New()
	data := []int{56, 37, 89, 2, 10, 72}
	for _, element := range data {
		b.Insert(element)
	}

	var items = []int{}
	b.Walk(func(item int) {
		items = append(items, item)
	})

	assert.Equal([]int{2, 10, 37, 56, 72, 89}, items, "walk a tree shold give the exact elements as inserted.")
}
