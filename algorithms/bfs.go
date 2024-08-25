package algorithms

import (
	"github.com/harshcommits/go-prep/ds"
	"golang.org/x/exp/slices"
)

func BFS(graph map[string][]string, start string) []string {

	queue := ds.Queue[string]{}
	queue.Enqueue(start)

	visited := []string{}

	for len(queue.Entries) > 0 {
		current := queue.Dequeue()
		if !slices.Contains(visited, current) {
			visited = append(visited, current)
		}

		for _, neighbour := range graph[current] {
			if !slices.Contains(visited, neighbour) {
				queue.Enqueue(neighbour)
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
