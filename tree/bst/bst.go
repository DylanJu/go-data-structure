package bst

type BinarySearchTree struct {
	root *treeNode
}

func (bst *BinarySearchTree) Insert(val int) {
	bst.root.addNode(val)
}

func (bst *BinarySearchTree) Search(val int) bool {
	node := bst.root

	for node != nil {
		if val < node.value {
			if node.left == nil {
				return false
			}
			node = node.left
		} else if val > node.value {
			if node.right == nil {
				return false
			}
			node = node.right
		} else {
			return true
		}
	}

	return false
}

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (node *treeNode) addNode(val int) {
	if val < node.value {
		if node.left == nil {
			node.left = &treeNode{value: val}
		} else {
			node.left.addNode(val)
		}
	} else if val > node.value {
		if node.right == nil {
			node.right = &treeNode{value: val}
		} else {
			node.right.addNode(val)
		}
	}
}

func MakeBinarySearchTree(value int) *BinarySearchTree {
	bst := &BinarySearchTree{root: &treeNode{value: value}}
	return bst
}
