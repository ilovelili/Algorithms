package binarysearchtree

// insert recursive
func (root *Node) insert(newNode *Node) {
	if newNode.Value <= root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			root.Left.insert(newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
		} else {
			root.Right.insert(newNode)
		}
	}
}

// Insert insert a node
func (bst *BST) Insert(value int) {
	if bst.root == nil {
		bst.root = &Node{Left: nil, Right: nil, Value: value}
		bst.size = 1
	} else {
		bst.root.insert(&Node{Left: nil, Right: nil, Value: value})
	}
}
