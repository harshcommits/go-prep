package main

import (
	"fmt"

	"github.com/harshcommits/go-prep/algorithms"
)

func main() {
	fmt.Println("This is for running the data structures")

	// stack := ds.Stack{}
	// items := []string{'A', 'B', 'C', 'D', 'E'}

	// for _, value := range items {
	// 	stack.Push(value)
	// }

	// for len(stack.Items) > 0 {
	// 	fmt.Printf("%c ", stack.Pop())
	// }

	// linkedlist := ds.LinkedList{}
	// linkedlist.Add(10)
	// linkedlist.Add(20)
	// linkedlist.Add(30)

	// fmt.Println(linkedlist.Length())
	// fmt.Println(linkedlist.GetValues())

	// concurrency.RunWaitGroup()
	// concurrency.SelectFunc()

	// graph traversal
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D", "E"}
	graph["C"] = []string{"A", "F"}
	graph["D"] = []string{"B"}
	graph["E"] = []string{"B", "F"}
	graph["F"] = []string{"C", "E"}

	// values := algorithms.BFS(graph, 'F')
	// for _, char := range values {
	// 	fmt.Printf("%c ", char)
	// }

	moreValues := algorithms.DFS(graph, "F", nil)
	fmt.Println(algorithms.Traversed(moreValues))

}
