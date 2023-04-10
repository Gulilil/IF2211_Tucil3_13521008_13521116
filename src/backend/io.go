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
			vInfo := strings.Split(content.Text(), " ")
			g.AddVertex(vInfo[0])
			long,_ := strconv.ParseFloat(vInfo[1], 64)
			g.vertices[idx].longPos = float64(long)
			lat, _ := strconv.ParseFloat(vInfo[2],64)
			g.vertices[idx].latPos = float64(lat)
			idx++
		} else {
			line := strings.Split(content.Text(), " ")
			for i:= 0 ; i < n; i++{
				val, _ := strconv.ParseFloat(line[i], 64)
				if (val != 0){
					g.AddEdge(*g.vertices[tempIdx], *g.vertices[i], val)
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

func AskUserStartEndVertices () (string, string) {
	var startV string
	var endV string
	fmt.Print("What is the name of the start Vertex: ")
	fmt.Scanln(&startV)
	fmt.Print("What is the name of the end Vertex: ")
	fmt.Scanln(&endV)
	return startV, endV

}

