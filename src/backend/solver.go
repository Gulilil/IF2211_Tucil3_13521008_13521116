package backend

import (
	"fmt"
)

type Solver struct {
	solRoute Route
}

func (s *Solver) SolveUCS(g Graph, startVKey string, endVKey string) {

	s.ResetSolver()
	startVertex := g.GetVertex(startVKey)

	if startVertex != nil {
		// Declaring Variables
		q := &QueueRoute{}
		curRoute := &Route{}

		// Preparing Variables Setup
		curRoute.InsertLastVertex(*startVertex)
		q.Enqueue(*curRoute)

		count := 0
		check := false

		// Looping for finding solution
		for !check {

			if q.nRoute == 0 {
				break
			}

			curRoute = q.Dequeue()
			curVertex := curRoute.GetLastVertex()

			if IsSolution(*curRoute, endVKey) {
				check = true
				break
			}

			availableEdges := AvailableEdges(*curRoute, g.GetEdgeWithStartV(curVertex))
			if len(availableEdges) != 0 {
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
			count++
		}

		// Checking if solution found
		if !check {
			fmt.Println("No Solution Found")
		} else {
			s.solRoute.CopyConstructorRoute(curRoute)
		}
	}
}

func (s *Solver) SolveAStar(g Graph, startVKey string, endVKey string) {

	s.ResetSolver()
	startVertex := g.GetVertex(startVKey)

	if startVertex != nil {
		// Declaring Variables
		q := &QueueRoute{}
		curRoute := &Route{}

		// Preparing Variables Setup
		curRoute.InsertLastVertex(*g.GetVertex(startVKey))
		q.Enqueue(*curRoute)

		goalVertex := g.GetVertex(endVKey)

		count := 0
		check := false

		// Looping for finding solution
		for !check {

			if q.nRoute == 0 {
				break
			}

			curRoute = q.Dequeue()
			curVertex := curRoute.GetLastVertex()

			if IsSolution(*curRoute, endVKey) {
				check = true
				break
			}

			availableEdges := AvailableEdges(*curRoute, g.GetEdgeWithStartV(curVertex))
			if len(availableEdges) != 0 {
				temp := &Route{}
				temp.CopyConstructorRoute(curRoute)
				for i, e := range availableEdges {
					if i > 0 {
						temp.CopyRoute(curRoute)
					}
					temp.InsertLastVertex(e.endVertex)
					temp.accWeight += e.weight
					temp.aStarDistance = e.endVertex.calculateDistance(*goalVertex)
					q.Enqueue(*temp)
				}
			}
			q.SortAStarAscending()
			count++
		}

		// Checking if solution found
		if !check {
			fmt.Println("No Solution Found")
		} else {
			s.solRoute.CopyConstructorRoute(curRoute)
		}
	}
}

func AvailableEdges(r Route, edges []*Edge) []*Edge {
	result := []*Edge{}
	for _, e := range edges {
		if !IsContainVertex(r.buffer, e.endVertex) {
			result = append(result, e)
		}
	}
	return result
}

func (s *Solver) DisplaySolutionRoute() {
	s.solRoute.DisplayRoute()
}

func IsSolution(r Route, endVKey string) bool {
	return r.GetLastVertex().key == endVKey
}

func (s* Solver) SolutionRouteToString () string {
	var result string
	for i := 0 ; i < s.solRoute.nVertex; i++ {
		if i == s.solRoute.nVertex -1 {
			result = result + "(" + s.solRoute.buffer[i].key +")"
		} else {
			result = result + "("+ s.solRoute.buffer[i].key +") -> "
		}
	}
	return result
}

func (s *Solver) GetSolutionNodes() int {
	return s.solRoute.nVertex
}

func (s *Solver) GetSolutionCost() float64 {
	return (s.solRoute.accWeight)
}

func (s *Solver) ResetSolver () {
	s.solRoute = Route{}
}

func (s *Solver) GetSolutionRoute() Route {
	return s.solRoute
}