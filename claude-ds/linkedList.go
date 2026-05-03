package claudeds

// ----------------------------------------------------------------------------
// LinkedList — singly linked, with head/tail pointers for O(1) append.
// For most real workloads prefer a slice; reach for a linked list only when
// you need O(1) splice or pointer-stable nodes. stdlib `container/list`
// provides a doubly-linked version.
// ----------------------------------------------------------------------------

type LinkedList[T any] struct {
	head, tail *listNode[T]
	size       int
}

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

func (l *LinkedList[T]) Len() int { return l.size }

func (l *LinkedList[T]) PushFront(v T) {
	l.head = &listNode[T]{val: v, next: l.head}
	if l.tail == nil {
		l.tail = l.head
	}
	l.size++
}

func (l *LinkedList[T]) PushBack(v T) {
	n := &listNode[T]{val: v}
	if l.tail == nil {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

func (l *LinkedList[T]) Each(fn func(T)) {
	for n := l.head; n != nil; n = n.next {
		fn(n.val)
	}
}
