package claudeds

import "cmp"

// ----------------------------------------------------------------------------
// BST — binary search tree over any cmp.Ordered type.
// ----------------------------------------------------------------------------

type BST[T cmp.Ordered] struct {
	root *bstNode[T]
	size int
}

type bstNode[T cmp.Ordered] struct {
	val         T
	left, right *bstNode[T]
}

func (t *BST[T]) Len() int { return t.size }

func (t *BST[T]) Insert(v T) {
	var inserted bool
	t.root, inserted = bstInsert(t.root, v)
	if inserted {
		t.size++
	}
}

func bstInsert[T cmp.Ordered](n *bstNode[T], v T) (*bstNode[T], bool) {
	if n == nil {
		return &bstNode[T]{val: v}, true
	}
	switch cmp.Compare(v, n.val) {
	case -1:
		var ok bool
		n.left, ok = bstInsert(n.left, v)
		return n, ok
	case 1:
		var ok bool
		n.right, ok = bstInsert(n.right, v)
		return n, ok
	}
	return n, false
}

func (t *BST[T]) Has(v T) bool {
	n := t.root
	for n != nil {
		switch cmp.Compare(v, n.val) {
		case -1:
			n = n.left
		case 1:
			n = n.right
		default:
			return true
		}
	}
	return false
}

// InOrder visits values in ascending order.
func (t *BST[T]) InOrder(fn func(T)) { bstInOrder(t.root, fn) }

func bstInOrder[T cmp.Ordered](n *bstNode[T], fn func(T)) {
	if n == nil {
		return
	}
	bstInOrder(n.left, fn)
	fn(n.val)
	bstInOrder(n.right, fn)
}
