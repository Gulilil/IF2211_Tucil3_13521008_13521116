package main

import (
	"RoutePlanner/src/backend"
	"fmt"
)

func main() {
	backend.OpenSplashScreen()
	println("")
	fullPath := backend.AskingUserInput()
	g := backend.ReadFileToGraph(fullPath)
	println("")
	println("Show information of graph : ")
	g.DisplayGraph()

	println("")
	start, end := backend.AskUserStartEndVertices()

	s := &backend.Solver{}
	var algorithm string
	println("")
	fmt.Println("What algorithm do you want to use: ")
	fmt.Println("1. Uniform Cost Search")
	fmt.Println("2. A* Search")
	fmt.Print("Input: ")
	fmt.Scanln(&algorithm)
	for algorithm != "1" && algorithm != "2" {
		println("")
		fmt.Println("What algorithm do you want to use: ")
		fmt.Println("1. Uniform Cost Search")
		fmt.Println("2. A* Search")
		fmt.Print("Input: ")
		fmt.Scanln(&algorithm)
	}

	if algorithm == "1" {
		s.SolveUCS(*g, start, end)
	} else if algorithm == "2" {
		s.SolveAStar(*g, start, end)
	}

	println("")
	fmt.Print("Solution Route: ")
	s.DisplaySolutionRoute()
	println("")
	backend.CloseSplashScreen()
}
