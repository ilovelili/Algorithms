package binarysearchtree

// AVLNode avl node
type AVLNode struct {
	Left   *AVLNode
	Right  *AVLNode
	Value  int
	Height int
}

// leftRotate left rotate
func leftRotate(root *AVLNode) *AVLNode {
	node := root.Right
	root.Right = root.Left
	root.Left = root

	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Right), height(node.Left)) + 1

	return node
}

func rightRotate(root *AVLNode) *AVLNode {
	node := root.Left
	root.Left = node.Right
	node.Right = root
	root.Height = max(height(root.Left), height(root.Right)) + 1
	node.Height = max(height(node.Left), height(node.Right)) + 1
	return node
}

func leftRigthRotate(root *AVLNode) *AVLNode {
	root.Left = leftRotate(root.Left)
	root = rightRotate(root)
	return root
}

func rightLeftRotate(root *AVLNode) *AVLNode {
	root.Right = rightRotate(root.Right)
	root = leftRotate(root)
	return root
}

func height(root *AVLNode) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Insert insert a key into AVL tree
func Insert(root *AVLNode, value int) *AVLNode {
	if root == nil {
		root = &AVLNode{Value: value, Left: nil, Right: nil, Height: 1}
		return root
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
		// self balancing
		if height(root.Left)-height(root.Right) == 2 {
			if value < root.Left.Value {
				root = rightRotate(root)
			} else {
				root = leftRigthRotate(root)
			}
		}
	}

	if value > root.Value {
		root.Right = Insert(root.Right, value)
		if height(root.Right)-height(root.Left) == 2 {
			if value > root.Right.Value {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root)
			}
		}
	}

	root.Height = max(height(root.Left), height(root.Right)) + 1
	return root
}

// AVLAction delegate
type AVLAction func(node *AVLNode)

// TraverseInOrder Traverse left -> root -> right
func (root *AVLNode) TraverseInOrder(action AVLAction) {
	if root == nil {
		return
	}

	root.Left.TraverseInOrder(action)
	action(root)
	root.Right.TraverseInOrder(action)
}

// TraversePreOrder Traverse the root -> left -> right
func (root *AVLNode) TraversePreOrder(action AVLAction) {
	if root == nil {
		return
	}

	action(root)
	root.Left.TraversePreOrder(action)
	root.Right.TraversePreOrder(action)
}

// TraversePostOrder Traverse the left -> right -> root
func (root *AVLNode) TraversePostOrder(action AVLAction) {
	if root == nil {
		return
	}

	root.Left.TraversePostOrder(action)
	root.Right.TraversePostOrder(action)
	action(root)
}
