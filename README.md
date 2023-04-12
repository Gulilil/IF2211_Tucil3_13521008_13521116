# IF2211_Strategi Algoritma

## *Tucil3_13521008_13521116*

## **Table of Contents**
* [Program Description](#program-description)
* [Required Program](#required-program)
* [How to Run The Program](#how-to-run-the-program)
* [Implementation Screenshots](#implementation-screenshots)
* [Progress Report](#progress-report)
* [Folders and Files Description](#folders-and-files-description)
* [Author](#author)

## **Program Description**
This program is a shortest-path finder for traversal graphs by implementing the uniform cost search (UCS) and A-Star (A*) algorithm for solving. The graph used by the program will be inputted from a .txt file. Users can customize the graph in the .txt file. After inputting the file, users can choose which algorithm they want to use to find the path. Users also can determine the desired start and end of the vertex they want to search. The user will get the visualization of the graph, the path from start to end of the vertex, the total of nodes, and the total of cost.

## **Required Program**
| Required Program | Reference Link|
|-------------------|-------------|
| Go (Golang)       | [Go (Golang)](https://go.dev/doc/install) |
| Tailwind CSS   | [Tailwind CSS](https://tailwindcss.com/docs/installation) |
| Graphviz | [Graphviz]() |

## **How to Run The Program**
1. Clone this repo using this command

```
git clone https://github.com/Gulilil/Tucil3_13521008_13521116.git
```

2. Change the current directory into the src folder
```
cd Tucil3_13521008_13521116
```


## **Implementation Screenshots**


## **Progress Report**

| Point | Yes | No |
|-----|-----|------|
|The program is able to calculate shortest path using UCS method| &check; |    |
|The program is able to calculate shortest path using A* method | &check; |  |
|The program is able to display the calculated solution, both the shortest path and the distance| &check; |  |
| Implemented  Bonus Aspect using Google Map API |  |  |


## **Folders and Files Description**
```bash                             
├── doc
    ├── Spesifikasi Tucil 3 Stima 2023.pdf
    └── Tucil_13521008_13521116.pdf
├── node_modules
├── src
    ├── frontend
        ├── assets
            ├── graph.png
            ├── output.css
            └── style.css
        └── main.go
    ├── backend
        ├── graph.go
        ├── io.go
        ├── queue.go
        ├── route.go
        └── solver.go
    └── main.go
├── test
    ├── graph1.txt
    ├── graph2.txt
    ├── graph3.txt
    ├── graph4.txt
    └── graph5.txt
├── go.mod
├── go.sum
├── package-lock.json
├── package.json
├── tailwind.config.js
└── README.md
```

## **Author**
| Name | Student ID |
|-------|------------|
| Jason Rivalino | 13521008 |
| Juan Christopher Santoso | 13521116|
