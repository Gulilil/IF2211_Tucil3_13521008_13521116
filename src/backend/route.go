package backend

type Route struct {
	buffer    []*Vertex
	nVertex   int
	accWeight int
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

func (r *Route) IsEmpty() bool {
	return r.nVertex == 0
}

func (r *Route) IsLessWeight(r2 Route) bool {
	return r.accWeight < r2.accWeight
}