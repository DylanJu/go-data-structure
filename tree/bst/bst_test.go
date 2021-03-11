package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BST(t *testing.T) {
	bst := MakeBinarySearchTree(9)

	assert.True(t, bst.root.value == 9)

	bst.Insert(5)
	bst.Insert(15)
	assert.True(t, bst.root.left.value == 5)
	assert.True(t, bst.root.right.value == 15)

	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(1)
	assert.True(t, bst.root.left.left.value == 3)
	assert.True(t, bst.root.left.left.right.value == 4)
	assert.True(t, bst.root.left.left.left.value == 2)
	assert.True(t, bst.root.left.left.left.left.value == 1)

	bst.Insert(10)
	bst.Insert(19)
	bst.Insert(12)
	bst.Insert(17)
	bst.Insert(22)
	assert.True(t, bst.root.right.left.value == 10)
	assert.True(t, bst.root.right.left.right.value == 12)
	assert.True(t, bst.root.right.right.value == 19)
	assert.True(t, bst.root.right.right.left.value == 17)
	assert.True(t, bst.root.right.right.right.value == 22)

	assert.True(t, bst.Search(9))
	assert.True(t, bst.Search(17))
	assert.True(t, bst.Search(3))
	assert.False(t, bst.Search(26))
	assert.False(t, bst.Search(0))
}
