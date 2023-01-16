package router

import "math"

type (
	Weight = float64
)

type Vertex struct {
	Name    string
	Edges   []Edge
	Weight  Weight
	From    *Vertex
	Visited bool
}

type Edge struct {
	To     Vertex
	Weight Weight
}

type DirectedWeightedGraph struct {
	graph   map[string]*Vertex
}

func NewDirectedWeightedGraph() *DirectedWeightedGraph {
	return &DirectedWeightedGraph{
		graph:   make(map[string]*Vertex),
	}
}

func (d *DirectedWeightedGraph) AddVertex(vertex Vertex) {
	d.graph[vertex.Name] = &vertex
}

func (d *DirectedWeightedGraph) Init(startStop Vertex) {
	for _, item := range d.graph {
		item.Weight = math.Inf(1)
		item.Visited = false
	}
	d.graph[startStop.Name].Weight = 0
}
