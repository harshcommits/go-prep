# Pointers in Go: Use Cases and Examples

## Common Use Cases

### 1. **Modifying Struct Fields**
Pointers allow methods to modify the receiver's fields. Without pointers, you'd be working on a copy.

```go
// With pointer receiver - modifies the original
func (t *Tree) Insert(value int) {
    t.root = &TreeNode{value: value}  // Changes the actual tree
}

// Without pointer - changes only the copy
func (t Tree) Insert(value int) {
    t.root = &TreeNode{value: value}  // Changes a copy, original unchanged
}
```

### 2. **Avoiding Large Memory Copies**
Passing pointers is cheaper than copying large structs.

```go
// Efficient - just passes a reference
func traverseTree(node *TreeNode) { }

// Inefficient - copies entire struct
func traverseTree(node TreeNode) { }
```

### 3. **Building Linked Data Structures**
Trees, linked lists, and graphs need pointers to reference child/sibling nodes.

```go
type TreeNode struct {
    value int
    left  *TreeNode  // Points to left child
    right *TreeNode  // Points to right child
}
```

### 4. **Nil Checks (Base Cases)**
Pointers can be `nil`, making it easy to represent "no node" in recursion.

```go
func depthOfNode(node *TreeNode) int {
    if node == nil {  // Base case
        return 0
    }
    // ... recurse
}
```

### 5. **Sharing References**
Multiple variables can reference the same object without duplication.

```go
node := NewTreeNode(5)
tree.root = node
otherVar := node  // Same reference, not a copy
```

---

## LeetCode Problem Examples

### Example 1: Binary Tree Maximum Path Sum
```go
// Need pointers to traverse and access child nodes
func maxPathSum(root *TreeNode) int {
    var result int
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {  // Pointer allows nil check
            return 0
        }
        left := dfs(node.left)   // Access child via pointer
        right := dfs(node.right)
        result = max(result, left + right + node.Val)
        return node.Val + max(0, max(left, right))
    }
    dfs(root)
    return result
}
```

### Example 2: Linked List Reversal
```go
// Pointers essential for modifying node connections
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next      // Save next before modifying
        head.Next = prev       // Reverse the pointer
        prev = head
        head = next
    }
    return prev
}
```

---

## Summary Table

| Use Case | Pointer? | Why |
|----------|----------|-----|
| Modifying original struct | ✅ | Value receivers get copies |
| Large structs | ✅ | Avoid expensive copies |
| Tree/Graph nodes | ✅ | Need references to children |
| Nil as "empty" | ✅ | Values can't be nil |
| Reading small data | ❌ | More efficient without indirection |

---

## Real-World Example: Your Tree Implementation

Your `tree.go` file demonstrates several pointer patterns:

### 1. Pointer Receivers for Modification
```go
// Modifies the actual tree structure
func (t *Tree) Insert(value int) {
    t.insertNode(t.root, value, 1)
}
```

### 2. Recursive Tree Traversal with Nil Checks
```go
// Nil check handles base case
func (t *Tree) InOrderTraversalByNode(sb *strings.Builder, root *TreeNode) {
    if root == nil {
        return
    }
    // ... recursively process children
}
```

### 3. Getter Methods with Pointer Receivers
```go
// Safe to call on nil nodes
func (n *TreeNode) GetLeft() *TreeNode {
    return n.left
}
```

### 4. Building Linked Structure
```go
type TreeNode struct {
    value int
    left  *TreeNode  // Child references
    right *TreeNode
}
```

All these patterns are essential for working with tree data structures efficiently in Go.