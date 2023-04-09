package backend

import (
	"time"
)

type Solver struct {
	solRoute Route
	curVertex Vertex
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
		*curRoute = q.Dequeue()
		

		if (IsSolution(*curRoute, endVKey)){
			check = true
			break
		}
	}
}

func (s Solver) SolveAStar(g Graph, startVKey string, endVKey string) {

	// Declaring Variables
	curRoute := &Route{}

	// Preparing Variables Setup
	curRoute.InsertLastVertex(Vertex{key : startVKey})

	check := false
	for (!check) {


		if (IsSolution(*curRoute, endVKey)){
			check = true
			break
		}
	}
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

