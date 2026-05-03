package ds

import (
	"fmt"
	"strings"
)

type Graph struct {
	vertices map[int]*Vertex
	directed bool
}

type Vertex struct {
	key      int
	adjacent map[int]*Edge
}

type Edge struct {
	to     *Vertex
	weight float64
}

// NewGraph creates a new graph. If directed is true, edges are one-way; otherwise, edges are bidirectional.
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertices: make(map[int]*Vertex),
		directed: directed,
	}
}

// AddVertex adds a vertex with the given key to the graph.
func (g *Graph) AddVertex(k int) error {
	if _, exists := g.vertices[k]; exists {
		return fmt.Errorf("vertex %d already exists", k)
	}
	g.vertices[k] = &Vertex{
		key:      k,
		adjacent: make(map[int]*Edge),
	}
	return nil
}

// AddEdge adds an unweighted edge (weight = 1.0) from vertex 'from' to vertex 'to'.
func (g *Graph) AddEdge(from, to int) error {
	return g.AddWeightedEdge(from, to, 1.0)
}

// AddWeightedEdge adds a weighted edge from vertex 'from' to vertex 'to' with the specified weight.
func (g *Graph) AddWeightedEdge(from, to int, weight float64) error {
	fromV, ok := g.vertices[from]
	if !ok {
		return fmt.Errorf("vertex %d not found", from)
	}
	toV, ok := g.vertices[to]
	if !ok {
		return fmt.Errorf("vertex %d not found", to)
	}

	fromV.adjacent[to] = &Edge{to: toV, weight: weight}

	// For undirected graphs, add the reverse edge
	if !g.directed {
		toV.adjacent[from] = &Edge{to: fromV, weight: weight}
	}

	return nil
}

// GetVertex returns the vertex with the given key, or nil if it doesn't exist.
func (g *Graph) GetVertex(k int) *Vertex {
	return g.vertices[k]
}

// HasVertex checks if a vertex exists in the graph.
func (g *Graph) HasVertex(k int) bool {
	_, exists := g.vertices[k]
	return exists
}

// GetNeighbors returns all neighboring vertices of the given vertex.
func (g *Graph) GetNeighbors(k int) ([]*Vertex, error) {
	v, ok := g.vertices[k]
	if !ok {
		return nil, fmt.Errorf("vertex %d not found", k)
	}

	neighbors := make([]*Vertex, 0, len(v.adjacent))
	for _, edge := range v.adjacent {
		neighbors = append(neighbors, edge.to)
	}
	return neighbors, nil
}

// Print outputs the graph structure to the console.
func (g *Graph) Print() {
	fmt.Println("\nGraph structure:")
	for key := range g.vertices {
		neighbors := g.vertices[key].adjacent
		if len(neighbors) == 0 {
			fmt.Printf("Vertex %d: (no edges)\n", key)
			continue
		}

		neighborKeys := make([]string, 0, len(neighbors))
		for nKey, edge := range neighbors {
			neighborKeys = append(neighborKeys, fmt.Sprintf("%d(w:%.1f)", nKey, edge.weight))
		}
		fmt.Printf("Vertex %d -> [%s]\n", key, strings.Join(neighborKeys, ", "))
	}
	fmt.Println()
}

// VertexCount returns the total number of vertices in the graph.
func (g *Graph) VertexCount() int {
	return len(g.vertices)
}

// EdgeCount returns the total number of edges in the graph.
func (g *Graph) EdgeCount() int {
	count := 0
	for _, v := range g.vertices {
		count += len(v.adjacent)
	}
	// For undirected graphs, each edge is counted twice, so divide by 2
	if !g.directed {
		count /= 2
	}
	return count
}
