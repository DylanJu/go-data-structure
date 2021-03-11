package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakedLinkedList(t *testing.T) {
	linkedList := MakeLinkedList(1)
	assert.True(t, linkedList.head.value == 1)
	assert.True(t, linkedList.head.next == nil)

	linkedList.addNode(2)
	linkedList.addNode(3)
	linkedList.addNode(4)
	assert.True(t, linkedList.head.value == 1)
	assert.True(t, linkedList.head.next.value == 2)
	assert.True(t, linkedList.head.next.next.value == 3)
	assert.True(t, linkedList.head.next.next.next.value == 4)
	assert.True(t, linkedList.tail.value == 4)

	linkedList.prependNode(5)
	assert.True(t, linkedList.head.value == 5)
	assert.True(t, linkedList.head.next.value == 1)

	linkedList.insertNode(6, 2)
	assert.True(t, linkedList.head.next.next.next.value == 6)
	assert.True(t, linkedList.head.next.next.value == 2)
	linkedList.insertNode(7, 4)
	assert.True(t, linkedList.head.next.next.next.next.next.next.value == 7)
	assert.True(t, linkedList.tail.value == 7)

	linkedList.removeNode(6)
	assert.True(t, linkedList.head.next.next.value == 2)
	assert.True(t, linkedList.head.next.next.next.value == 3)
	linkedList.removeNode(5)
	assert.True(t, linkedList.head.value == 1)
	linkedList.removeNode(7)
	assert.True(t, linkedList.tail.value == 4)

	linkedList.printNodes()
}
