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

func TestPushBack(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	l.PushBack(3)
	assert.Equal(l.Length(), 1, "After inserting, list should have one element.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")
	assert.True(l.Find(3) != -1, "1 should be in list.")

	// Test append multi elements
	data := []interface{}{4, 5, 6}
	l.PushBack(data...)
	assert.Equal(l.Length(), 4, "After inserting, list should have 4 elements.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")

	for _, elem := range data {
		assert.True(l.Find(elem) != -1, fmt.Sprintf("%d should be in list.", elem))
	}
}

func TestPushFront(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	l.PushFront(3)
	assert.Equal(l.Length(), 1, "After inserting, list should have one element.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")
	assert.True(l.Find(3) != -1, "1 should be in list.")

	// Test append multi elements
	data := []interface{}{4, 5, 6}
	l.PushFront(data...)
	assert.Equal(l.Length(), 4, "After inserting, list should have 4 elements.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")

	for _, elem := range data {
		assert.True(l.Find(elem) != -1, fmt.Sprintf("%d should be in list.", elem))
	}
}

func TestFind(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	data := []interface{}{3, 4, 5, 6}
	l.PushBack(data...)

	for index, elem := range data {
		assert.Equal(l.Find(elem), index, fmt.Sprintf("%d should be at %d.", elem, index))
	}

	assert.Equal(l.Find(100), -1, "Find a non-exist element in list should return -1")
}

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	data := []interface{}{3, 4, 5, 6}
	l.PushBack(data...)
	assert.Equal(l.Length(), 4, "After inserting, list should have 4 elements.")
	assert.False(l.IsEmpty(), "After inserting, list should not be empty.")

	for i, elem := range data {
		res, err := l.Index(i)
		assert.NoError(err, "List should return value at index %d", i)
		assert.Equal(res, elem, fmt.Sprintf("List should have %v at %d.", elem, i))

		res, _ = l.Index(i - 4)
		assert.Equal(res, elem, fmt.Sprintf("List should have %v at %d.", elem, i-4))
	}

	res, err := l.Index(10)
	assert.Error(err, "Index exceeds list length should return err")
	assert.Nil(res, "Index exceeds list length should return nil")

	res, err = l.Index(-10)
	assert.Error(err, "Index exceeds list length should return err")
	assert.Nil(res, "Index exceeds list length should return nil")
}

func TestLpop(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	// Test append multi elements
	data := []interface{}{3, 4, 5, 6}
	l.PushBack(data...)

	for _, elem := range data {
		res := l.Lpop()
		assert.Equal(res, elem, "List lpop should return front value")
	}

	assert.Nil(l.Lpop(), "Lpop from en empty list should return nil")
}

func TestRpop(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	// Test append multi elements
	data := []interface{}{3, 4, 5, 6}
	l.PushFront(data...)

	for _, elem := range data {
		res := l.Rpop()
		assert.Equal(res, elem, "List rpop should return front value")
	}

	assert.Nil(l.Rpop(), "Rpop from en empty list should return nil")
}

func TestRange(t *testing.T) {
	assert := assert.New(t)
	l := list.New()

	// Test append multi elements
	data := []interface{}{3, 4, 5, 6}
	l.PushBack(data...)
	assert.Equal([]interface{}{3}, l.Range(0, 0), "In range [0,0] value should be 3")
	assert.Equal([]interface{}{3, 4, 5, 6}, l.Range(0, -1), "Values in range [0, -1] should return all values")
	assert.Equal([]interface{}{3, 4, 5, 6}, l.Range(0, 10), "Values in range [0, 10] should return all values")
	assert.Equal([]interface{}{3}, l.Range(0, -10), "Values in range [0, -10] should return the fisrt element")
	assert.Equal([]interface{}{3, 4}, l.Range(0, 1), "Values in range [0, 1] should return the first two values")
	assert.Equal([]interface{}{}, l.Range(3, 1), "Values in range [3, 1] should return empty slice")
	assert.Equal([]interface{}{}, l.Range(3, -2), "Values in range [3, -2] should return empty slice")
	assert.Equal([]interface{}{6}, l.Range(3, 3), "Values in range [3, 3] should return [6]")
	assert.Equal([]interface{}{4, 5, 6}, l.Range(-3, -1), "Values in range [-3, -3] should return [4, 5, 6]")
}
