package backend

import "fmt"

type Vertex struct {
	key string
}

type Edge struct {
	startVertex Vertex
	endVertex   Vertex
	weight      int
}

type Graph struct {
	nVertex  int
	vertices []*Vertex
	nEdge    int
	edges    []*Edge
}

func (g *Graph) AddVertex(k string) {
	v := Vertex{
		key: k,
	}
	if !IsContainVertex(g.vertices, v) {
		g.vertices = append(g.vertices, &v)
		g.nVertex++
	}
}

func (g *Graph) AddEdge(k1 string, k2 string, w int) {
	v1 := Vertex{
		key: k1,
	}
	v2 := Vertex{
		key: k2,
	}
	e := Edge{
		startVertex: v1,
		endVertex:   v2,
		weight:      w,
	}
	if !IsContainEdge(g.edges, e) {
		g.edges = append(g.edges, &e)
		g.nEdge++
	}
}

func (g *Graph) GetVertex(k string) *Vertex {
	for i := 0; i < g.nVertex; i++ {
		if g.vertices[i].key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) GetEdgeWithStartV(startV Vertex) []*Edge {
	result := []*Edge{}
	for i := 0; i < g.nEdge; i++ {
		if IsSameVertex(g.edges[i].startVertex, startV) {
			result = append(result, g.edges[i])
		}
	}
	return result
}

func (g *Graph) GetEdgeWithEndV(endV Vertex) []*Edge {
	result := []*Edge{}
	for i := 0; i < g.nEdge; i++ {
		if IsSameVertex(g.edges[i].endVertex, endV) {
			result = append(result, g.edges[i])
		}
	}
	return result
}

func IsContainVertex(vertices []*Vertex, vertex Vertex) bool {
	for i := 0; i < len(vertices); i++ {
		if vertices[i].key == vertex.key {
			return true
		}
	}
	return false
}

func IsContainEdge(edges []*Edge, edge Edge) bool {
	for i := 0; i < len(edges); i++ {
		if edges[i].startVertex == edge.startVertex && edges[i].endVertex == edge.endVertex {
			return true
		}
	}
	return false
}

func IsSameVertex(v1 Vertex, v2 Vertex) bool {
	return v1.key == v2.key
}

func (g *Graph) DisplayGraph() {
	fmt.Printf("Amount of Vertices : %d\n", g.nVertex)
	for i := 0; i < g.nVertex; i++ {
		fmt.Println(g.vertices[i].key)
	}
	fmt.Printf("Amount of Edges : %d\n", g.nEdge)
	for i := 0; i < g.nEdge; i++ {
		fmt.Printf("| \"%s\" -> (%d) -> \"%s\" |\n", g.edges[i].startVertex.key, g.edges[i].weight, g.edges[i].endVertex.key)
	}
}

// func main() {
// 	g := &Graph{}
// 	g.addVertex("haha")
// 	g.addVertex("huhu")
// 	g.addVertex("hihi")
// 	g.addVertex("hehe")
// 	g.addEdge("haha", "hehe", 4)
// 	g.addEdge("hehe", "hihi", 7)
// 	g.displayGraph()
// }
