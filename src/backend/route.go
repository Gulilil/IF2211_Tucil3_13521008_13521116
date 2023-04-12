package backend

import "fmt"

type Route struct {
	buffer        []*Vertex
	nVertex       int
	accWeight     float64
	aStarDistance float64
}

func (r *Route) InsertLastVertex(v Vertex) {
	r.buffer = append(r.buffer, &v)
	r.nVertex++
}

func (r *Route) DeleteFirstVertex() Vertex {
	temp := *r.buffer[0]
	r.buffer = r.buffer[1:]
	r.nVertex--
	return temp
}

func (r *Route) GetLastVertex() Vertex {
	return *r.buffer[r.nVertex-1]
}

func (r *Route) IsEmpty() bool {
	return r.nVertex == 0
}

func (r *Route) IsLessWeight(r2 Route) bool {
	return r.accWeight < r2.accWeight
}

func (r *Route) IsAStarLess(r2 Route) bool {
	return r.accWeight+r.aStarDistance < r2.accWeight+r2.aStarDistance
}

func (r *Route) CopyConstructorRoute(r2 *Route) {
	// Make sure that r is empty
	r.nVertex = r2.nVertex
	r.accWeight = r2.accWeight
	r.aStarDistance = r2.aStarDistance
	for i := range r2.buffer {
		r.buffer = append(r.buffer, r2.buffer[i])
	}
}

func (r *Route) CopyRoute(r2 *Route) {
	r.nVertex = r2.nVertex
	r.accWeight = r2.accWeight
	r.aStarDistance = r2.aStarDistance
	r.buffer = r2.buffer[:r2.nVertex]
}

func (r *Route) DisplayRoute() {
	for i := 0; i < r.nVertex; i++ {
		fmt.Print(r.buffer[i].key, " ")
	}
	println("")
	fmt.Println("Total of Nodes: ", r.nVertex)
	fmt.Println("Total of Cost: ", r.accWeight)
}
