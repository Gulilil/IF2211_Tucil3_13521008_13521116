package backend

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"strconv"
)

func IdentifyAmountVertex(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	content := bufio.NewScanner(file)
	content.Split(bufio.ScanLines)
	for content.Scan() {
		count++
	}
	return count/2
}

func ReadFileToGraph(path string) *Graph {
	g := &Graph{}
	idx := 0
	tempIdx := 0

	n := IdentifyAmountVertex(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	content := bufio.NewScanner(file)
	content.Split(bufio.ScanLines)
	for content.Scan() {
		if idx < n {
			g.AddVertex(content.Text())
			idx++
		} else {
			line := strings.Split(content.Text(), " ")
			for i:= 0 ; i < n; i++{
				val, _ := strconv.Atoi(line[i])
				if (val != 0){
					g.AddEdge(g.vertices[tempIdx].key, g.vertices[i].key, val)
				}
			}
			tempIdx++
		}
	}
	return g
}

func AskingUserInput() string {
	curDir, _ := os.Getwd()
	var name string
	fmt.Print("Which file do you want to use: ")
	fmt.Scanln(&name)
	name = name + ".txt"
	fullPath := path.Join(curDir, "test", name)

	return fullPath
}

