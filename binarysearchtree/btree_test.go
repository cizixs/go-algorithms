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
		assert.True(b.Exists(element), "Element %d should be in tree.", element)
	}
}
