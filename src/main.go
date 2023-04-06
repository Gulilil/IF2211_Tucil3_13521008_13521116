package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/", homePageHandler)
	fmt.Println("Server started at => localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	curDir, _ := os.Getwd()
	var filepath = path.Join(curDir, "src", "frontend", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Route Planner",
		"name":  "xixixixi",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
