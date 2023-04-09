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
	fmt.Println("huhuhu UCS")
}

func (s Solver) SolveAStar(g Graph, startVKey string, endVKey string) {
	q := &QueueRoute{}
	curRoute := &Route{}
	

	check := false
	for (!check) {

		if (IsSolution(*curRoute, endVKey)){
			check = true
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
	return r.buffer[r.nVertex-1].key == endVKey
}

