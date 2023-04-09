package main

import (
	"RoutePlanner/src/backend"
)

func main() {
	fullPath := backend.AskingUserInput()
	g := backend.ReadFileToGraph(fullPath)
	g.DisplayGraph()

	// s := backend.Solver{}
	// s.SolveAStar(g, "London", "Berlin")
}