package deque_test

import (
	"testing"

	"github.com/cizixs/go-algorithms/deque"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	d := deque.New()

	assert.True(d.Length() == 0, "New initial deque should be empty")
}

func TestPush(t *testing.T) {
	assert := assert.New(t)
	d := deque.New()

	data := []interface{}{2, "hello", true}
	for _, item := range data {
		d.Push(item)
	}

	assert.Equal(d.Length(), uint(len(data)))
}

func TestPushPop(t *testing.T) {
	assert := assert.New(t)
	d := deque.New()

	data := []interface{}{2, "hello", true}
	for _, item := range data {
		d.Push(item)
	}

	assert.Equal(d.Length(), uint(len(data)))

	for i := 0; i < len(data); i++ {
		item, _ := d.Pop()
		assert.Equal(item, data[len(data)-1-i], "The popped element is not as expected.")
	}
}

func TestInjectEject(t *testing.T) {
	assert := assert.New(t)
	d := deque.New()

	data := []interface{}{2, "hello", true}
	for _, item := range data {
		d.Inject(item)
	}

	assert.Equal(d.Length(), uint(len(data)))

	for i := 0; i < len(data); i++ {
		item, _ := d.Eject()
		assert.Equal(item, data[len(data)-1-i], "The ejected element is not as expected.")
	}
}
