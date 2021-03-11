package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeStack(t *testing.T) {
	s := MakeStack()

	assert.Equal(t, len(s.items), 0, "Initial length is 0")
	assert.Equal(t, cap(s.items), 0, "Initial capacity is 0")

	s.Push(3)
	assert.True(t, len(s.items) == 1)
	assert.Equal(t, s.items[0], 3, "It can push a new value")

	s.Push(4)
	s.Push(5)
	s.Push(6)
	assert.True(t, len(s.items) == 4)
	assert.Equal(t, s.items[3], 6, "It can push a new value in the last index")

	last := s.Pop()
	assert.True(t, last == 6)
	assert.True(t, len(s.items) == 3)

	s.Pop()
	s.Pop()
	assert.True(t, len(s.items) == 1)
}
