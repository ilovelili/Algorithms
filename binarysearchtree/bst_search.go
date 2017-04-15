package binarysearchtree

// Search check if a node with given value is in the tree
func (bst *BST) Search(value int) bool {
	return bst.root.Search(value)
}

// Search check if a node with given value is in the tree
func (root *Node) Search(value int) bool {
	if root == nil {
		return false
	}

	if root.Value == value {
		return true
	}

	if root.Value < value {
		return root.Right.Search(value)
	}

	return root.Left.Search(value)
}
