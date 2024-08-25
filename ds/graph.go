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

func (v *Vertex) NewVertex(map[int][]int) {

}
