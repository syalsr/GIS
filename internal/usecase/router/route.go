package router

type Interface interface {
	BuildRouter(start, finish string) *Vertex
	FillData(vertex *Vertex, emptyRoadDistance bool)
}

type Router struct {
	Dwg *DirectedWeightedGraph
}

func NewRouter() Interface {
	return &Router{
		Dwg: NewDirectedWeightedGraph(),
	}
}

func (r *Router) BuildRouter(start, finish string) *Vertex {
	proceed := Queue{vertex: []Vertex{*r.Dwg.graph[start]}}
	for !proceed.IsEmpty() {
		current := proceed.Pop()
		for _, item := range current.Edges {
			if r.Dwg.graph[item.To.Name].Visited {
				continue
			}
			proceed.Push(item.To)

			i := 0
			to := r.Dwg.graph[current.Name].Edges[i].To
			Weight := r.Dwg.graph[current.Name].Edges[i].Weight

			if !r.Dwg.graph[to.Name].Visited {
				if current.Weight+Weight < r.Dwg.graph[to.Name].Weight {
					r.Dwg.graph[to.Name].Weight = current.Weight + Weight
					r.Dwg.graph[to.Name].From = &current

				}
			}

			i++
		}
	}
	return r.Dwg.graph[finish]
}

func (r *Router) FillData(vertex *Vertex, emptyRoadDistance bool) {
	if emptyRoadDistance {
		return
	}
	r.Dwg.graph[vertex.Name] = vertex
}

type Queue struct {
	vertex []Vertex
}

func (q *Queue) Len() int           { return len(q.vertex) }
func (q *Queue) Less(i, j int) bool { return q.vertex[i].Weight < q.vertex[j].Weight }
func (q *Queue) Swap(i, j int)      { q.vertex[i], q.vertex[j] = q.vertex[j], q.vertex[i] }
func (q *Queue) IsEmpty() bool      { return len(q.vertex) == 0 }

func (q *Queue) Push(vertex Vertex) {
	q.vertex = append(q.vertex, vertex)
}

func (q *Queue) Pop() Vertex {
	v := q.vertex[0]
	q.vertex = q.vertex[1:]
	return v
}
