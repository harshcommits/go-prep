package ds

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (list *LinkedList) Add(value int) {

	node := &Node{data: value, next: nil}

	if list.head == nil {
		list.head = node
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = node
	list.len++

}

func (list *LinkedList) Length() int {
	return list.len + 1
}

func (list *LinkedList) GetValues() []int {
	values := []int{}

	for current := list.head; current != nil; current = current.next {
		values = append(values, current.data)
	}

	return values
}
