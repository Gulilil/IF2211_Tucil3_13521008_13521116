package main

import (
	"RoutePlanner/src/backend"
)

func main() {
	fullPath := backend.AskingUserInput()
	g := backend.ReadFileToGraph(fullPath)
	// g.DisplayGraph()

	start, end := backend.AskUserStartEndVertices()

	s := backend.Solver{}
	s.SolveUCS(*g, start, end)
	// fmt.Print("Solution Route : ")
	// s.DisplaySolutionRoute()
}