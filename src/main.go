package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

type Data struct {
	Vertices string
	Solution string
	TotalNodes int
	TotalCost int
}


func main() {

	http.HandleFunc("/", homePageHandler)
	fmt.Println("Server started at => localhost:9000/")
	assetsPath := path.Join(getSrcPath(), "frontend", "assets")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(assetsPath))))
	http.ListenAndServe(":9000", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	// s := &backend.Solver{}
	// g := &backend.Graph {}


	// var data = Data{
	// 	Vertices: "(Node1) (Node2) (Node3) (Node4) (Node5)",
	// 	Solution: "(Node1) (Node3) (Node5)",
	// 	TotalNodes : 4,
	// 	TotalCost : 240,
	// }

	var data = Data {}

	var filePath = path.Join(getSrcPath(), "frontend", "*.html")
	var tmpl= template.Must(template.ParseGlob(filePath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getSrcPath() string{
	curDir, _ := os.Getwd()
	return path.Join(curDir, "src")
}
