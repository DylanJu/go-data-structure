package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeQueue(t *testing.T) {
	q := MakeQueue()

	assert.Equal(t, len(q.items), 0, "Initial length is 0")
	assert.Equal(t, cap(q.items), 0, "Initial capacity is 0")

	q.Enqueue(3)
	assert.True(t, len(q.items) == 1)
	assert.Equal(t, q.items[0], 3, "It can Enqueue a new value")

	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	assert.True(t, len(q.items) == 4)
	assert.Equal(t, q.items[0], 3, "It can Enqueue a new value in the last index")

	first := q.Dequeue()
	assert.True(t, first == 3)
	assert.True(t, len(q.items) == 3)

	q.Dequeue()
	q.Dequeue()
	assert.True(t, len(q.items) == 1)
}
