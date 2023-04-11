package main

import (
	"RoutePlanner/src/backend"
	"fmt"
)

func main() {
	fullPath := backend.AskingUserInput()
	g := backend.ReadFileToGraph(fullPath)
	// g.DisplayGraph()

	start, end := backend.AskUserStartEndVertices()

	s := &backend.Solver{}
	s.SolveAStar(*g, start, end)
	fmt.Print("Solution Route : ")
	s.DisplaySolutionRoute()
	fmt.Println(s.GetExecutionTime())
}