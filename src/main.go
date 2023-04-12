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
	assetsPath := path.Join(getSrcPath(), "frontend", "assets")

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/result", processHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(assetsPath))))
	
	fmt.Println("Server started at => localhost:9000/")
	http.ListenAndServe(":9000", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var filePath = path.Join(getSrcPath(), "frontend", "index.html")	
		var tmpl= template.Must(template.New("form").ParseFiles(filePath))
	
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
	}


}

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var filePath = path.Join(getSrcPath(), "frontend", "result.html")
		var tmpl = template.Must(template.New("result").ParseFiles(filePath))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var algo = r.FormValue("algoritma")
		var file = r.FormValue("myfile")
		var start = r.FormValue("startV")
		var end = r.FormValue("endV")

		fmt.Println(algo)
		fmt.Println(file)
		fmt.Println(start)
		fmt.Println(end)

		var data = Data {
			Vertices : start,
			Solution : end,
			TotalNodes: 0,
			TotalCost: 0,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	http.Error(w, "", http.StatusBadRequest)
}


func getSrcPath() string{
	curDir, _ := os.Getwd()
	return path.Join(curDir, "src")
}

