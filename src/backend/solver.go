package backend

import (
	"fmt"
	"time"
)

type Solver struct {
	solRoute Route
	start time.Time
	duration time.Duration
}

func (s Solver) SolveUCS(g Graph, startVKey string, endVKey string) {
	// Declaring Variables
	q := &QueueRoute{}
	curRoute := &Route{}

	// Preparing Variables Setup
	curRoute.InsertLastVertex(Vertex{key : startVKey})
	q.Enqueue(*curRoute)

	check := false
	for (!check) {
		curRoute = q.Dequeue()
		

		if (IsSolution(*curRoute, endVKey)){
			check = true
			break
		}
	}
}

func (s Solver) SolveAStar(g Graph, startVKey string, endVKey string) {

	s.StartTime()
	// Declaring Variables
	q := &QueueRoute{}
	curRoute := &Route{}
	// Preparing Variables Setup
	curRoute.InsertLastVertex(*g.GetVertex(startVKey))
	q.Enqueue(*curRoute)

	count := 0
	check := false
	for (!check) {

		curRoute = q.Dequeue()
		curVertex := curRoute.GetLastVertex()

		fmt.Println(curVertex)

		if (IsSolution(*curRoute, endVKey)){
			check = true
			break
		}		

		availableEdges := availableEdges(*curRoute, g.GetEdgeWithStartV(curVertex))
		if (len(availableEdges) != 0){ 
			temp := &Route{}
			temp.CopyConstructorRoute(curRoute)
			for i, e := range availableEdges {
				if i > 0 {
					temp.CopyRoute(curRoute)
				}
				temp.InsertLastVertex(e.endVertex)
				temp.accWeight += e.weight
				q.Enqueue(*temp)
			}
		}
		q.SortAscending()

		fmt.Print("curRoute : ")
		curRoute.DisplayRoute()
		q.DisplayQueue()

		if (count == 3){
			check = true
		}
		count++
	}
	s.StopTime()
}

func availableEdges(r Route, edges []*Edge) [] *Edge {
	result := []*Edge{}
	for _, e := range edges {
		if (!IsContainVertex(r.buffer, e.endVertex)){
			result = append(result, e)
		}
	}
	return result
}

func getOptimalEdgeAStar(g Graph, edges []*Edge, endVKey string) Edge {
	goalVertex := g.GetVertex(endVKey)

	minDistance := float64(0)
	optimalE := Edge{}
	for i, e := range edges{
		distance := e.endVertex.calculateDistance(*goalVertex) + e.weight
		if i == 0{
			minDistance = distance
			optimalE = *e
		} else {
			if (distance < minDistance){
				minDistance = distance
				optimalE = *e
			}
		}
	}
	return optimalE
}

func (s Solver) StartTime() {
	s.start = time.Now()
}

func (s Solver) StopTime() {
	s.duration = time.Since(s.start)
}

func IsSolution(r Route, endVKey string) bool {
	return r.GetLastVertex().key == endVKey
}

