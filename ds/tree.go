package ds

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root  *TreeNode
	depth int
}

// NewTree creates a new tree with a root value
func NewTree(value int) *Tree {
	return &Tree{
		root:  &TreeNode{value: value},
		depth: 1,
	}
}

// NewTreeNode creates a new tree node with the given value
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

// Insert adds a new node to the tree (BST insertion)
func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &TreeNode{value: value}
		t.depth = 1
		return
	}
	t.insertNode(t.root, value, 1)
}

// insertNode is a helper function for BST insertion
func (t *Tree) insertNode(node *TreeNode, value int, currentDepth int) {
	if currentDepth > t.depth {
		t.depth = currentDepth
	}

	if value < node.value {
		if node.left == nil {
			node.left = &TreeNode{value: value}
		} else {
			t.insertNode(node.left, value, currentDepth+1)
		}
	} else {
		if node.right == nil {
			node.right = &TreeNode{value: value}
		} else {
			t.insertNode(node.right, value, currentDepth+1)
		}
	}
}

// // GetValue returns the value of a node as a string
// func (n TreeNode) GetValue() string {
// 	return strconv.Itoa(n.value)
// }

// GetValue method for pointer receiver
func (n *TreeNode) GetValue() string {
	if n == nil {
		return ""
	}
	return strconv.Itoa(n.value)
}

// GetTreeValues returns in-order traversal as a string
func (t *Tree) GetTreeValues() string {
	sb := strings.Builder{}
	t.InOrderTraversal(&sb)
	return sb.String()
}

// InOrderTraversal performs in-order traversal (Left-Root-Right)
// Result: sorted order for BST
func (t *Tree) InOrderTraversal(sb *strings.Builder) {
	t.InOrderTraversalByNode(sb, t.root)
}

// InOrderTraversalByNode performs in-order traversal starting from a given node
func (t *Tree) InOrderTraversalByNode(sb *strings.Builder, root *TreeNode) {
	if root == nil {
		return
	}

	// Traverse left subtree
	t.InOrderTraversalByNode(sb, root.left)

	// Visit root
	sb.WriteString(root.GetValue())
	sb.WriteString(" ")

	// Traverse right subtree
	t.InOrderTraversalByNode(sb, root.right)
}

// PreOrderTraversal performs pre-order traversal (Root-Left-Right)
// Result: useful for creating copy of tree, prefix notation
func (t *Tree) PreOrderTraversal(sb *strings.Builder) {
	t.preOrderTraversalByNode(sb, t.root)
}

// preOrderTraversalByNode performs pre-order traversal starting from a given node
func (t *Tree) preOrderTraversalByNode(sb *strings.Builder, root *TreeNode) {
	if root == nil {
		return
	}

	// Visit root
	sb.WriteString(root.GetValue())
	sb.WriteString(" ")

	// Traverse left subtree
	t.preOrderTraversalByNode(sb, root.left)

	// Traverse right subtree
	t.preOrderTraversalByNode(sb, root.right)
}

// PostOrderTraversal performs post-order traversal (Left-Right-Root)
// Result: useful for deleting tree, postfix notation
func (t *Tree) PostOrderTraversal(sb *strings.Builder) {
	t.postOrderTraversalByNode(sb, t.root)
}

// postOrderTraversalByNode performs post-order traversal starting from a given node
func (t *Tree) postOrderTraversalByNode(sb *strings.Builder, root *TreeNode) {
	if root == nil {
		return
	}

	// Traverse left subtree
	t.postOrderTraversalByNode(sb, root.left)

	// Traverse right subtree
	t.postOrderTraversalByNode(sb, root.right)

	// Visit root
	sb.WriteString(root.GetValue())
	sb.WriteString(" ")
}

// LevelOrderTraversal performs level-order traversal (BFS)
// Result: breadth-first, level by level
func (t *Tree) LevelOrderTraversal(sb *strings.Builder) {
	if t.root == nil {
		return
	}

	queue := []*TreeNode{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		sb.WriteString(node.GetValue())
		sb.WriteString(" ")

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

// Search checks if a value exists in the tree (BST search)
func (t *Tree) Search(value int) bool {
	return t.searchNode(t.root, value)
}

// searchNode is a helper function for BST search
func (t *Tree) searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if value == node.value {
		return true
	} else if value < node.value {
		return t.searchNode(node.left, value)
	} else {
		return t.searchNode(node.right, value)
	}
}

// GetHeight returns the height of the tree
func (t *Tree) GetHeight() int {
	return t.getHeightNode(t.root)
}

// getHeightNode is a helper function to get height from a node
func (t *Tree) getHeightNode(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := t.getHeightNode(node.left)
	rightHeight := t.getHeightNode(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// GetRoot returns the root node of the tree
func (t *Tree) GetRoot() *TreeNode {
	return t.root
}

// SetRoot sets the root node of the tree
func (t *Tree) SetRoot(node *TreeNode) {
	t.root = node
}
