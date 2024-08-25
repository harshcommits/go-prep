package algorithms

import (
	"github.com/harshcommits/go-prep/ds"
	"golang.org/x/exp/slices"
)

func BFS(graph map[string][]string, start string) []string {

	stack := ds.Stack{}
	stack.Push(start)

	var visited []string
	var traversed []string

	for len(stack.Items) > 0 {
		current := stack.Pop()
		visited = append(visited, current)

		for _, neighbour := range graph[current] {
			if !slices.Contains(visited, neighbour) {
				stack.Push(neighbour)
				visited = append(visited, neighbour)
			}
		}
	}

	// remove duplicate values
	for _, value := range visited {
		if !slices.Contains(traversed, value) {
			traversed = append(traversed, value)
		}
	}

	return traversed

}
