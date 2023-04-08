package backend

type Route struct {
	buffer    []*Vertex
	nVertex   int
	accWeight int
}


func (r *Route) insertLastVertex(v Vertex) {
	r.buffer = append(r.buffer, &v)
	r.nVertex++
}

func (r *Route) deleteFirstVertex() Vertex {
	temp := *r.buffer[0]
	r.buffer = r.buffer[1:]
	r.nVertex--
	return temp
}

func (r *Route) isEmpty() bool {
	return r.nVertex == 0
} 

func (r *Route) isLessWeight( r2 Route) bool {
	return r.accWeight < r2.accWeight
}