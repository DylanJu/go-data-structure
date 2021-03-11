package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Heap(t *testing.T) {
	heap := &Heap{}
	heap.Insert(50)

	assert.True(t, heap.nodes[0] == 50)

	heap.Insert(16)
	assert.True(t, heap.nodes[1] == 16)

	heap.Insert(48)
	assert.True(t, heap.nodes[2] == 48)

	heap.Insert(14)
	assert.True(t, heap.nodes[3] == 14)
	heap.Insert(8)
	assert.True(t, heap.nodes[4] == 8)
	heap.Insert(34)
	assert.True(t, heap.nodes[5] == 34)
	heap.Insert(20)
	assert.True(t, heap.nodes[6] == 20)

	heap.Insert(9)
	assert.True(t, heap.nodes[7] == 9)
	heap.Insert(1)
	assert.True(t, heap.nodes[8] == 1)
	heap.Insert(5)
	assert.True(t, heap.nodes[9] == 5)
	heap.Insert(7)
	assert.True(t, heap.nodes[10] == 7)

	heap.Insert(63)
	assert.True(t, heap.nodes[0] == 63)
	fmt.Println(heap.nodes)

	max := heap.Pop()
	assert.True(t, max == 63)
	assert.True(t, heap.nodes[0] == 50)
	assert.True(t, len(heap.nodes) == 11)
	fmt.Println(heap.nodes)

	max = heap.Pop()
	assert.True(t, max == 50)
	assert.True(t, heap.nodes[0] == 48)
	fmt.Println(heap.nodes)

	max = heap.Pop()
	assert.True(t, max == 48)
	assert.True(t, heap.nodes[0] == 34)
	fmt.Println(heap.nodes)

	heap.Pop()
	fmt.Println(heap.nodes)

	max = heap.Pop()
	assert.True(t, max == 20)
	assert.True(t, heap.nodes[0] == 16)
	fmt.Println(heap.nodes)
}
