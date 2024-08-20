package main

import (
	"fmt"

	"github.com/harshcommits/go-ds/concurrency"
	"github.com/harshcommits/go-ds/ds"
)

func main() {
	fmt.Println("This is for running the data structures")

	stack := ds.Stack{}
	stack_values := []int{1, 2, 4, 5, 6, 9}

	for _, value := range stack_values {
		stack.Push(value)
	}

	linkedlist := ds.LinkedList{}
	linkedlist.Add(10)
	linkedlist.Add(20)
	linkedlist.Add(30)

	fmt.Println(linkedlist.Length())
	fmt.Println(linkedlist.GetValues())

	concurrency.RunWaitGroup()

}
