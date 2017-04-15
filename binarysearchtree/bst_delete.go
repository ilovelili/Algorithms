package binarysearchtree

// minValue left most node
func (root *Node) minValue() int {
	if root.Left == nil {
		return root.Value
	}

	return root.Left.minValue()
}

// Exists node exists or not
func (root *Node) Exists(value int) bool {
	if root == nil {
		return false
	}

	if root.Value == value {
		return true
	}

	if root.Value < value {
		return root.Left.Exists(value)
	}

	return root.Right.Exists(value)
}

// link link nodes
func link(parent, node *Node) {
	if parent.Left == node {
		if node.Left != nil {
			parent.Left = node.Left
		} else {
			parent.Left = node.Right
		}
	} else if parent.Right == node {
		if node.Left != nil {
			parent.Right = node.Left
		} else {
			parent.Right = node.Right
		}
	}
}

// delete delete node
func delete(node, parent *Node, value int) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		if node.Left != nil && node.Right != nil {
			node.Value = node.Right.minValue()
			return delete(node.Right, node, value)
		}

		// else replace the parent node with left or right
		link(parent, node)
		return true
	}

	if node.Value < value {
		if node.Right == nil {
			return false
		}
		return delete(node.Right, node, value)
	}

	if node.Value > value {
		if node.Left == nil {
			return false
		}
		return delete(node.Left, node, value)
	}

	return false
}

// Delete delete a node
func (bst *BST) Delete(value int) bool {
	if !bst.root.Exists(value) {
		return false
	}

	if bst.root.Value == value {
		temproot := &Node{Left: bst.root}
		r := delete(bst.root, temproot, value)
		bst.root = temproot.Left
		return r
	}

	// delete right
	if bst.root.Value < value {
		return delete(bst.root.Right, bst.root, value)
	}

	// delete left
	return delete(bst.root.Left, bst.root, value)
}
