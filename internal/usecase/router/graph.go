package router

import "math"

type (
	Weight   = int
)

type Vertex struct {
	name    string
	edges []Edge
	weight Weight
	from *Vertex
	visited bool
}

type Edge struct {
	to     Vertex
	weight Weight
}

type DirectedWeightedGraph struct {
	graph   []Vertex
	weights map[*Vertex]int
}

func NewDirectedWeightedGraph() *DirectedWeightedGraph {
	return &DirectedWeightedGraph{
		weights: make(map[*Vertex]int),
	}
}

func (d *DirectedWeightedGraph) AddVertex(vertex Vertex) {
	d.graph = append(d.graph, vertex)
	d.weights[&vertex] = vertex.weight
}

func (d *DirectedWeightedGraph) Init(startStop Vertex) {
	for _, item := range d.graph {
		item.weight = int(math.Inf(1))
		item.visited = false
	}
}
