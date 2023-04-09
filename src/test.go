package main

import (
	"RoutePlanner/src/backend"
)

func main() {
	fullPath := backend.AskingUserInput()
	g := backend.ReadFileToGraph(fullPath)
	g.DisplayGraph()

	s := Solver{}
	s.SolveAStar(g, "London", "Berlin")
}