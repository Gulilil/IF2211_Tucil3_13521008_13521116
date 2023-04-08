package backend

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func readFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}

func main() {
	curDir, _ := os.Getwd()
	test := path.Join("test", "graph1.txt")
	fullPath := path.Join(curDir, test)
	readFile(fullPath)
}
