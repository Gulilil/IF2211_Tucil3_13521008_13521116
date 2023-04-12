package backend

import (
	"fmt"
	"math"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Vertex struct {
	key string
	longPos float64
	latPos float64
}

type Edge struct {
	startVertex Vertex
	endVertex   Vertex
	weight      float64
}

type Graph struct {
	nVertex  int
	vertices []*Vertex
	nEdge    int
	edges    []*Edge
}

func (v Vertex) calculateDistance(v2 Vertex) float64 {
	return math.Sqrt(math.Pow(v2.latPos - v.latPos, 2) + math.Pow(v2.longPos - v.longPos, 2))
}

func (v Vertex) copyVertex(v2 Vertex) {
	v.key = v2.key
	v.latPos = v2.latPos
	v.longPos = v2.longPos
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

func (g *Graph) AddEdge(v1 Vertex, v2 Vertex, w float64) {
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
		if IsSameVertex(*vertices[i], vertex) {
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
	return v1.key == v2.key && v1.latPos == v2.latPos && v1.longPos == v2.longPos
}

func (g *Graph) DisplayGraph() {
	fmt.Printf("Amount of Vertices : %d\n", g.nVertex)
	for i := 0; i < g.nVertex; i++ {
		fmt.Printf("%s  (%.2f, %.2f)\n",g.vertices[i].key, g.vertices[i].longPos, g.vertices[i].latPos)
	}
	fmt.Printf("Amount of Edges : %d\n", g.nEdge)
	for i := 0; i < g.nEdge; i++ {
		fmt.Printf("| \"%s\" -> (%.2f) -> \"%s\" |\n", g.edges[i].startVertex.key, g.edges[i].weight, g.edges[i].endVertex.key)
	}
}

func (g* Graph) GraphVerticesToString () string {
	var result string
	for i := 0 ; i < g.nVertex; i++ {
		if i == g.nVertex -1 {
			result = result + "(" + g.vertices[i].key +")"
		} else {
			result = result + "("+ g.vertices[i].key +"), "
		}
	}
	return result
}

func (g *Graph) VisualizeGraph(solution Route){
	gv := graph.New(graph.StringHash, graph.Weighted())

	for i := 0; i < g.nVertex; i++{
		if IsContainVertex(solution.buffer, *g.vertices[i]){
			gv.AddVertex(g.vertices[i].key, graph.VertexAttribute("colorscheme", "blues3"), graph.VertexAttribute("style", "filled"), graph.VertexAttribute("color", "2"), graph.VertexAttribute("fillcolor", "1"))
		} else {
			gv.AddVertex(g.vertices[i].key)
		}

	}
	for i := 0; i < g.nEdge; i++{
		weight := fmt.Sprintf("%.2f", g.edges[i].weight)
		if IsInSolution(solution, *g.edges[i]){
			gv.AddEdge(g.edges[i].startVertex.key, g.edges[i].endVertex.key, graph.EdgeAttribute("label", weight), graph.EdgeAttribute("color", "red"))
		} else {
			gv.AddEdge(g.edges[i].startVertex.key, g.edges[i].endVertex.key, graph.EdgeAttribute("label", weight), graph.EdgeAttribute("color", "black"))
		}
	}

	file, _ := os.Create("src/frontend/assets/solutionGraph.gv")
    _ = draw.DOT(gv, file)

	cmd := exec.Command("dot","-Tpng","-O","src/frontend/assets/solutionGraph.gv")

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not run command: ", err)
	}
}

func IsInSolution(r Route, e Edge) bool {
	for i:= 0; i < r.nVertex-1 ; i++{
		if (e.startVertex.key ==  r.buffer[i].key && e.endVertex.key ==  r.buffer[i+1].key){
			return true
		}
	}
	return false
}

