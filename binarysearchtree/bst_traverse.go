package binarysearchtree

import (
	"fmt"
)

// TraversePreOrder bst in pre order Traverse the root, left sub-tree, right sub-tree
func (root *Node) TraversePreOrder() {
	if root != nil {
		fmt.Println(root.Value)
		root.Left.TraverseInOrder()
		root.Right.TraverseInOrder()
	}
}

// TraversePostOrder bst in post order Traverse left sub-tree, right sub-tree, root
func (root *Node) TraversePostOrder() {
	if root != nil {
		root.Left.TraverseInOrder()
		fmt.Println(root.Value)
		root.Right.TraverseInOrder()
	}
}

// TraverseInOrder Traverse the left sub-tree, root, right sub-tree
func (root *Node) TraverseInOrder() {
	if root != nil {
		root.Left.TraverseInOrder()
		root.Right.TraverseInOrder()
		fmt.Println(root.Value)
	}
}
