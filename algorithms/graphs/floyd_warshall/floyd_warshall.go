package main

// Note: Adjacency list representation of graph was already implemented by me in previous graphs section. Hence I have
// modified few things which converts this adjacency list to adjacency matrix and then the floyd warshall algorithm is
// implemented. However you can directly take user inputs and add it in adjacency matrix.

import (
	"fmt"
	"math"
)

type vertexNumber struct {
	value int
	name  string
}

var adjacencyMatrix [][]float64
var vertices map[string]int

// FloydWarshall calculates all points shortest path from adjacency matrix of a graph
func FloydWarshall() {
	for k := range adjacencyMatrix {
		for i := range adjacencyMatrix {
			for j := range adjacencyMatrix {
				adjacencyMatrix[i][j] = math.Min(adjacencyMatrix[i][j], adjacencyMatrix[i][k]+adjacencyMatrix[k][j])
			}
		}
	}
	fmt.Println("\n-- Final Adjacency Matrix --")
	for i := range adjacencyMatrix {
		fmt.Println(adjacencyMatrix[i])
	}
}

func init() {
	vertices = make(map[string]int)
}

func main() {
	fmt.Println("\n-- Floyd Warshall's Algorithm --")
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. SIMPLE DISPLAY")
		fmt.Println("4. RUN FLOYD WARSHALL")
		fmt.Println("5. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			addVertex()
		case 2:
			addEdge()
		case 3:
			simpleDisplay()
		case 4:
			constructMatrix()
			FloydWarshall()
		case 5:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

// generates a adjacency matrix from adjacency list
func constructMatrix() {
	count := 0
	adjacencyMatrix = make([][]float64, len(vertices))
	for i := range graph {
		adjacencyMatrix[vertices[i]] = make([]float64, len(vertices))
		for j := range graph[i] {
			adjacencyMatrix[vertices[i]][vertices[graph[i][j].name]] = float64(graph[i][j].value)
		}
		count++
	}

	for i := range adjacencyMatrix {
		for j := range adjacencyMatrix[i] {
			if i != j && adjacencyMatrix[i][j] == 0 {
				adjacencyMatrix[i][j] = math.Inf(0)
			}
		}
	}
	fmt.Println("\n-- Initial Adjacency Matrix --")
	for i := range adjacencyMatrix {
		fmt.Println(adjacencyMatrix[i])
	}
}

func simpleDisplay() {
	fmt.Println("")
	for i := range graph {
		fmt.Print(i, " => ")
		for j := range graph[i] {
			fmt.Print(graph[i][j])
		}
		fmt.Println("")
	}
}

// addVertexToGraph("a")
// addVertexToGraph("b")
// addVertexToGraph("c")
// addVertexToGraph("d")
// addVertexToGraph("e")
// addVertexToGraph("f")
// addEdgeToGraph("a", "b", 18)
// addEdgeToGraph("a", "c", 20)
// addEdgeToGraph("b", "d", 4)
// addEdgeToGraph("c", "e", 10)
// addEdgeToGraph("d", "e", 15)
// addEdgeToGraph("d", "f", 23)
// addEdgeToGraph("e", "f", 7)
// addEdgeToGraph("d", "a", 5)
// addEdgeToGraph("e", "a", 17)
// addEdgeToGraph("b", "e", 21)
