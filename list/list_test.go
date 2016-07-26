package list_test

import (
	"fmt"
	"testing"

	"github.com/cizixs/go-algorithms/list"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	assert.True(l.IsEmpty(), "New initial list should be empty.")
	assert.Equal(l.Length(), 0, "New initial list should have length 0")
}

func TestAppend(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	l.Append(3)
	assert.Equal(l.Length(), 1, "After inserting, list should have one element.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")
	assert.True(l.Contains(3), "1 should be in list.")

	// Test append multi elements
	data := []interface{}{4, 5, 6}
	l.Append(data...)
	assert.Equal(l.Length(), 4, "After inserting, list should have 4 elements.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")

	for _, elem := range data {
		assert.True(l.Contains(elem), fmt.Sprintf("%d should be in list.", elem))
	}
}

func TestFind(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	data := []interface{}{3, 4, 5, 6}
	l.Append(data...)

	for index, elem := range data {
		assert.Equal(l.Find(elem), index, fmt.Sprintf("%d should be at %d.", elem, index))
	}
}
