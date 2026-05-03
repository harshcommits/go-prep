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

	graph := ds.NewGraph(false) // undirected graph
	graph.AddVertex(1)
	graph.AddVertex(10)
	graph.AddVertex(20)

	// added edges
	graph.AddEdge(1, 10)
	graph.AddEdge(10, 20)
	graph.AddEdge(20, 1)
	fmt.Println("This is the graph being printed")
	graph.Print()

	tree := ds.NewTree(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)

	fmt.Printf("Tree Depth: %d\n", tree.GetHeight())

	binarySearch := misc.BinarySearch(3, []int{1, 3, 4, 5})
	fmt.Printf("The value is %d\n", binarySearch)
	if misc.ValidPalindrome("naman") {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}

	currentVersion := misc.ReleaseType{Web: misc.Version{Major: 1, Minor: 0, Patch: 0}, Desktop: misc.Version{Major: 1, Minor: 5, Patch: 0}, Agent: misc.Version{Major: 1, Minor: 0, Patch: 4}}
	commits := []string{
		"feat(web): Added fixes",
		"feat(core): Modified base libraries",
	}

	misc.CalculateNextVersions(&currentVersion, commits)
	fmt.Println(currentVersion)

}
