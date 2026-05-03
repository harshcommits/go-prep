package claudeds

// ----------------------------------------------------------------------------
// Graph — directed adjacency list with BFS traversal.
// ----------------------------------------------------------------------------

type Graph[T comparable] struct {
	adj map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{adj: make(map[T][]T)}
}

func (g *Graph[T]) AddEdge(from, to T) {
	g.adj[from] = append(g.adj[from], to)
}

func (g *Graph[T]) BFS(start T, visit func(T)) {
	visited := Set[T]{}
	var q Queue[T]
	q.Enqueue(start)
	visited.Add(start)
	for q.Len() > 0 {
		v, _ := q.Dequeue()
		visit(v)
		for _, next := range g.adj[v] {
			if !visited.Has(next) {
				visited.Add(next)
				q.Enqueue(next)
			}
		}
	}
}
