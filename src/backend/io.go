package backend

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
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

func main() {
	curDir, _ := os.Getwd()
	test := path.Join("test", "graph1.txt")
	fullPath := path.Join(curDir, test)
	fmt.Println(fullPath)
	fmt.Println(IdentifyAmountVertex(fullPath))
}
