package algorithms

import (
	"github.com/harshcommits/go-prep/ds"
	"golang.org/x/exp/slices"
)

func DFS(graph map[string][]string, start string, visited []string) []string {

	stack := ds.Stack{}
	stack.Push(start)

	if visited != nil {
		visited = []string{}
	}

	for len(stack.Items) > 0 {
		current := stack.Pop()
		if !slices.Contains(visited, current) {
			visited = append(visited, current)
		}

		for _, neighbour := range graph[current] {
			if !slices.Contains(visited, neighbour) {
				stack.Push(neighbour)
				visited = append(visited, current)
			}
		}
	}

	// if visited == nil {
	// 	visited = []string{}
	// }

	// visited = append(visited, start)

	// for _, neighbour := range graph[start] {
	// 	if !slices.Contains(visited, neighbour) {
	// 		DFS(graph, neighbour, visited)
	// 	}
	// }

	return visited

}

func Traversed(values []string) []string {

	var traversed []string

	// remove duplicate values
	for _, value := range values {
		if !slices.Contains(traversed, value) {
			traversed = append(traversed, value)
		}
	}

	return traversed
}
