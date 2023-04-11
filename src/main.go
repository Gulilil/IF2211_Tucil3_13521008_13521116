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
	fmt.Println("Server started at => localhost:9000/")
	assetsPath := path.Join(getSrcPath(), "frontend", "assets")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(assetsPath))))
	http.ListenAndServe(":9000", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	var filePath = path.Join(getSrcPath(), "frontend", "*.html")
	var tmpl= template.Must(template.ParseGlob(filePath))


	var data = map[string]interface{}{
		"title": "Route Planner",
		"name":  "xixixixi",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getSrcPath() string{
	curDir, _ := os.Getwd()
	return path.Join(curDir, "src")
}
