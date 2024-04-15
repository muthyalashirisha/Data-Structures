package problems

import (
	"fmt"
)

// Node represents a node in the AVL Tree
type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

// AVLTree represents the AVL Tree
type AVLTree struct {
	root *Node
}

func (node *Node) init(data int, height int) {
	node.key = data
	node.height = height
	node.left = nil
	node.right = nil
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// height returns the height of the tree rooted at node
func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// rightRotate performs a right rotation on the subtree rooted at y
func (tree *AVLTree) rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right

	// Perform rotation
	x.right = y
	y.left = T2

	// Update heights
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	// Return new root
	return x
}

// leftRotate performs a left rotation on the subtree rooted at x
func (tree *AVLTree) leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left

	// Perform rotation
	y.left = x
	x.right = T2

	// Update heights
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	// Return new root
	return y
}

// getBalance returns the balance factor of node N
func (tree *AVLTree) getBalance(N *Node) int {
	if N == nil {
		return 0
	}
	return height(N.left) - height(N.right)
}

func AVLTreeMain() {
	tree := AVLTree{}

	// Constructing tree given in the example
	tree.root = tree.Insert(tree.root, 10)
	tree.root = tree.Insert(tree.root, 20)
	tree.root = tree.Insert(tree.root, 30)
	tree.root = tree.Insert(tree.root, 40)
	tree.root = tree.Insert(tree.root, 50)
	tree.root = tree.Insert(tree.root, 25)

	/* The constructed AVL Tree would be
	     30
	    /  \
	  20   40
	 /  \     \
	10  25    50
	*/
	fmt.Println("Preorder traversal of constructed tree is : ")
	PreOrder(tree.root)
}

// insert inserts a key into the AVL Tree
func (tree *AVLTree) Insert(node *Node, key int) *Node {
	// 1. Perform the normal BST insertion
	if node == nil {
		return &Node{key: key, height: 1}
	}

	if key < node.key {
		node.left = tree.Insert(node.left, key)
	} else if key > node.key {
		node.right = tree.Insert(node.right, key)
	} else {
		// Duplicate keys not allowed
		return node
	}

	// 2. Update height of this ancestor node
	node.height = 1 + max(height(node.left), height(node.right))
	fmt.Println("node.key:", node.key, "node.height:", node.height)

	// 3. Get the balance factor of this ancestor node to check whether this node became unbalanced
	balance := tree.getBalance(node)
	fmt.Println("balance:", balance)

	// If this node becomes unbalanced, then there are 4 cases

	// Left Left Case
	if balance > 1 && key < node.left.key {
		return tree.rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && key > node.right.key {
		return tree.leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && key > node.left.key {
		node.left = tree.leftRotate(node.left)
		return tree.rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && key < node.right.key {
		node.right = tree.rightRotate(node.right)
		return tree.leftRotate(node)
	}

	// Return the (unchanged) node pointer
	return node
}

// preOrder prints the preorder traversal of the tree
func PreOrder(node *Node) {
	if node != nil {
		fmt.Print(node.key, " ", "height", node.height, "\n")
		PreOrder(node.left)
		PreOrder(node.right)
	}
}
