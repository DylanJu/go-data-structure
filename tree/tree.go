package tree

import "fmt"

type Tree struct {
	root *TreeNode
}

func (t *Tree) DFS() {
	t.root.dfs()

	fmt.Println()
}

func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.root)

	for len(s) != 0 {
		lastNode := s[len(s)-1]
		// print
		fmt.Print(s[len(s)-1].value, " -> ")
		// pop
		s = s[:len(s)-1]
		// push children
		for i := len(lastNode.children) - 1; i >= 0; i-- {
			s = append(s, lastNode.children[i])
		}
	}
}

func (t *Tree) BFS() {
	q := []*TreeNode{}
	q = append(q, t.root)

	for len(q) != 0 {
		firstNode := q[0]
		// print
		fmt.Print(q[0].value, " -> ")
		// pop
		q = q[1:]
		// push children
		for i := 0; i < len(firstNode.children); i++ {
			q = append(q, firstNode.children[i])
		}
	}
}

type TreeNode struct {
	value    int
	children []*TreeNode
}

func (t *TreeNode) AppendNode(value int) {
	t.children = append(t.children, &TreeNode{value: value})
}

func (t *TreeNode) dfs() {
	fmt.Print(t.value, " -> ")

	for i := 0; i < len(t.children); i++ {
		t.children[i].dfs()
	}
}

func MakeTree(value int) Tree {
	tree := Tree{root: &TreeNode{value: value}}
	return tree
}
