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
	root  *Node
	depth int
}

func (n TreeNode) GetValue() string {
	return strconv.Itoa(n.value)
}

func (t Tree) GetTreeValues() string {
	sb := strings.Builder{}
	t.InOrderTraversal(&sb)
	return sb.String()
}

func (t Tree) InOrderTraversal(sb *strings.Builder) {

}

func (t Tree) InOrderTraversalByNode(sb *strings.Builder, root *TreeNode) {

}
