package binarysearchtree

// Node node
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// BST binary search tree
type BST struct {
	root *Node
	size int
}

// Size size of bst
func (bst *BST) Size() int {
	return bst.size
}

// Root root of bst
func (bst *BST) Root() *Node {
	return bst.root
}

// NewBST new bst
func NewBST() *BST {
	return &BST{root: nil, size: 0}
}
