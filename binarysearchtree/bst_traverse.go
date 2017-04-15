package binarysearchtree

type Action func(node *Node)

// TraversePreOrder bst in pre order Traverse the root, left sub-tree, right sub-tree
func (root *Node) TraversePreOrder(action Action) {
	if root != nil {
		action(root)
		root.Left.TraversePreOrder(action)
		root.Right.TraversePreOrder(action)
	}
}

// TraversePostOrder bst in post order Traverse left sub-tree, right sub-tree, root
func (root *Node) TraversePostOrder(action Action) {
	if root != nil {
		root.Left.TraversePostOrder(action)
		action(root)
		root.Right.TraversePostOrder(action)
	}
}

// TraverseInOrder Traverse the left sub-tree, root, right sub-tree
func (root *Node) TraverseInOrder(action Action) {
	if root != nil {
		root.Left.TraverseInOrder(action)
		root.Right.TraverseInOrder(action)
		action(root)
	}
}
