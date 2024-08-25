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

	// check_operator := math.Floor(float64((10 + 3) / 2))
	// fmt.Println(check_operator)

	values := []int{1, 2, 3, 4}
	valuesChar := []string{"A", "B", "C", "D"}

	fmt.Println(algorithms.BinarySearch(values, 3))
	fmt.Println(algorithms.BinarySearch(valuesChar, "E"))

	// graph traversal
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D", "E"}
	graph["C"] = []string{"A", "F"}
	graph["D"] = []string{"B"}
	graph["E"] = []string{"B", "F"}
	graph["F"] = []string{"C", "E"}

	// DFS traversal
	DFSValues := algorithms.DFS(graph, "F")
	fmt.Println("DFS traversal: ", algorithms.Traversed(DFSValues))

	// DFS recursively
	DFSRecursiveValues := algorithms.DFS(graph, "F")
	fmt.Println("DFS traversal done recursively: ", algorithms.Traversed(DFSRecursiveValues))

	// BFS traversal
	BFSValues := algorithms.BFS(graph, "F")
	fmt.Println("BFS traversal: ", algorithms.Traversed(BFSValues))

}
