package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeTree(t *testing.T) {
	tree := MakeTree(1)
	val := 2

	for i := 0; i < 3; i++ {
		tree.root.AppendNode(val)
		val++
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			tree.root.children[i].AppendNode(val)
			val++
		}
	}

	assert.True(t, tree.root.children[0].value == 2)
	assert.True(t, tree.root.children[1].value == 3)
	assert.True(t, tree.root.children[2].value == 4)

	assert.True(t, tree.root.children[0].children[0].value == 5)
	assert.True(t, tree.root.children[0].children[1].value == 6)

	assert.True(t, tree.root.children[1].children[0].value == 7)
	assert.True(t, tree.root.children[1].children[1].value == 8)

	assert.True(t, tree.root.children[2].children[0].value == 9)
	assert.True(t, tree.root.children[2].children[1].value == 10)

	tree.DFS()
	fmt.Println()
	tree.DFS2()
	fmt.Println()
	tree.BFS()
	fmt.Println()
}
