package ds

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Val        int
	Neighbours map[int]*Edge
}

type Edge struct {
	Value  int
	Weight int
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
	}
}

// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(val int) *Vertex {
	if _, exists := g.Vertices[val]; exists {
		return g.Vertices[val]
	}
	vertex := &Vertex{
		Val:        val,
		Neighbours: make(map[int]*Edge),
	}
	g.Vertices[val] = vertex
	return vertex
}

// AddEdge adds a weighted edge from vertex u to vertex v
func (g *Graph) AddEdge(u, v, weight int) {
	// Ensure both vertices exist
	if _, exists := g.Vertices[u]; !exists {
		g.AddVertex(u)
	}
	if _, exists := g.Vertices[v]; !exists {
		g.AddVertex(v)
	}

	// Add edge from u to v
	g.Vertices[u].Neighbours[v] = &Edge{
		Value:  v,
		Weight: weight,
	}
}

// AddUndirectedEdge adds an undirected weighted edge between u and v
func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// BuildFromAdjacencyList builds the graph from an adjacency list representation
// adjacencyList: map[vertexID] -> slice of neighbor vertex IDs (unweighted)
func (g *Graph) BuildFromAdjacencyList(adjacencyList map[int][]int) {
	for vertex, neighbours := range adjacencyList {
		g.AddVertex(vertex)
		for _, neighbour := range neighbours {
			g.AddEdge(vertex, neighbour, 1) // Default weight of 1 for unweighted graphs
		}
	}
}

// GetNeighbours returns all neighbors of a given vertex
func (g *Graph) GetNeighbours(val int) []*Vertex {
	if vertex, exists := g.Vertices[val]; exists {
		neighbours := make([]*Vertex, 0, len(vertex.Neighbours))
		for _, edge := range vertex.Neighbours {
			neighbours = append(neighbours, g.Vertices[edge.Value])
		}
		return neighbours
	}
	return nil
}

// GetEdge returns the edge between two vertices if it exists
func (g *Graph) GetEdge(u, v int) *Edge {
	if vertex, exists := g.Vertices[u]; exists {
		return vertex.Neighbours[v]
	}
	return nil
}

// HasVertex checks if a vertex exists in the graph
func (g *Graph) HasVertex(val int) bool {
	_, exists := g.Vertices[val]
	return exists
}

// Size returns the number of vertices in the graph
func (g *Graph) Size() int {
	return len(g.Vertices)
}
