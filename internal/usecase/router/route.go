package router

type Router struct {
	dwg *DirectedWeightedGraph
}

func (r *Router) BuildRouter(start, finish Vertex) {
	proceed := Queue{vertex: []Vertex{*r.dwg.graph[start.name]}} 
	for !proceed.IsEmpty() {
		current := proceed.Pop()
		for _, item := range current.edges {
			if r.dwg.graph[item.to.name].visited {
				continue
			}
			proceed.Push(item.to)

			i := 0
			to := r.dwg.graph[current.name].edges[i].to
			weight := r.dwg.graph[current.name].edges[i].weight

			if !r.dwg.graph[to.name].visited {
				if current.weight + weight < r.dwg.graph[to.name].weight {
					r.dwg.graph[to.name].weight = current.weight + weight
					r.dwg.graph[to.name].from = &current

				}
			}

			i++
		}
	}
}

type Queue struct {
	vertex []Vertex
}

func (q *Queue) Len() int           { return len(q.vertex) }
func (q *Queue) Less(i, j int) bool { return q.vertex[i].weight < q.vertex[j].weight }
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