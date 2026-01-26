package main

import (
	"fmt"

	"github.com/harshcommits/go-prep/ds"
	"github.com/harshcommits/go-prep/misc"
)

func main() {

	result := misc.TwoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)

	merged := misc.MergeStrings("ab", "pqrs")
	fmt.Println(merged)

	graph := ds.NewGraph()
	graph.AddUndirectedEdge(1, 2, 10)
	graph.AddUndirectedEdge(1, 3, 15)
	graph.AddUndirectedEdge(2, 4, 12)
	graph.AddUndirectedEdge(3, 4, 5)

	for v, vertex := range graph.Vertices {
		fmt.Printf("Vertex %d:\n", v)
		for n, edge := range vertex.Neighbours {
			fmt.Printf("  connects to %d with weight %d\n", n, edge.Weight)
		}
	}

	tree := ds.NewTree(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)

	fmt.Printf("Tree Depth: %d\n", tree.GetHeight())

}
