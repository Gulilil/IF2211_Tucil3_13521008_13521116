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

func (s Solver) solveUCS(g Graph, startVKey string, endVKey string) {
	fmt.Println("huhuhu UCS")
}

func (s Solver) solveAStar(g Graph, startVKey string, endVKey string) {
	fmt.Println("aaaaa AStar")
}

func (s Solver) startTime() {
	s.start = time.Now()
}

func (s Solver) stopTime() {
	s.duration = time.Since(s.start)
}

func isSolution(r Route, endVKey string) bool {
	return r.buffer[r.nVertex-1].key == endVKey
}