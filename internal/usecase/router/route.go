package router

type Router struct {
	dwg *DirectedWeightedGraph
}

func (r *Router) BuildRouter(start, finish VertexId) {
	weights := make(map[string]int)
	weights[start] = 0
	r.dwg.graph[start].weight = 0

	proceed := r.dwg.graph[start].edges
	currentProceed := r.dwg.graph[start]
	for _, value := range proceed {
		if value.to.visited {
			continue
		}
		value.to.visited = true
		proceed = append(proceed, value.to)

		if currentProceed.weight+value.weight < value.to.weight {
			
		}
	}
}
