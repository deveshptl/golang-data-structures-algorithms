package main

// Note: Adjacency list representation of graph was already implemented by me in previous graphs section. Hence I have
// modified few things which converts this adjacency list to adjacency matrix and then the bellman ford algorithm is
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
var distanceMatrix []float64
var vertices map[string]int

// BellmanFord calculates all points shortest path from adjacency matrix of a graph
func BellmanFord(start string) {
	flag := 0

	// start relaxing each vertex for |v|-1 times
	k := len(vertices) - 1
	for k != 0 {
		for i := range vertices {
			for j := range vertices {
				if distanceMatrix[vertices[i]]+adjacencyMatrix[vertices[i]][vertices[j]] < distanceMatrix[vertices[j]] {
					distanceMatrix[vertices[j]] = distanceMatrix[vertices[i]] + adjacencyMatrix[vertices[i]][vertices[j]]
				}
			}
		}
		k--
	}

	// display generated distance matrix
	fmt.Println("\nDistance Matrix: ")
	for i := range vertices {
		fmt.Println(i, "=>", distanceMatrix[vertices[i]])
	}
	fmt.Println("\nValidating the answer by relaxing each vertices one more time...")

	// validate the generated matrix by relaxing each vertex one more time
	for i := range vertices {
		for j := range vertices {
			if distanceMatrix[vertices[i]]+adjacencyMatrix[vertices[i]][vertices[j]] < distanceMatrix[vertices[j]] {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
	}

	if flag == 1 {
		fmt.Println("-- Distance matrix formed is WRONG. Found a loop with negative weight in graph. --")
	} else {
		fmt.Println("Distance matrix formed is CORRECT.")
	}

}

func init() {
	vertices = make(map[string]int)
}

func main() {
	fmt.Println("\n-- Bellman Ford Algorithm --")
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
			startBellmanFord()
		case 5:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

// generates an adjacency matrix from adjacency list
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

func startBellmanFord() {
	// scan for starting vertex name
	var vtx string
	fmt.Print("Enter the name of starting vertex: ")
	fmt.Scanf("%s\n", &vtx)

	// initialize distance matrix with necessary values
	distanceMatrix = make([]float64, len(vertices))
	count := len(vertices) - 1
	for count != -1 {
		if vertices[vtx] == count {
			distanceMatrix[count] = 0
		} else {
			distanceMatrix[count] = math.Inf(0)
		}
		count--
	}

	// validate few things before starting the algorithm
	if len(graph) == 0 {
		fmt.Println("\n-- Graph is null/empty. --")
		return
	} else if graph[vtx] == nil {
		fmt.Println("\n-- Start vertex not found in graph. --")
		return
	}

	// generate adjacency matrix and then start bellman ford algorithm
	constructMatrix()
	BellmanFord(vtx)
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

// Graph with no negative weight cycle
// addVertexToGraph("a")
// addVertexToGraph("b")
// addVertexToGraph("c")
// addVertexToGraph("d")
// addVertexToGraph("e")
// addVertexToGraph("f")
// addVertexToGraph("g")
// addEdgeToGraph("a", "b", 6)
// addEdgeToGraph("a", "c", 5)
// addEdgeToGraph("a", "d", 5)
// addEdgeToGraph("b", "e", -1)
// addEdgeToGraph("c", "b", -2)
// addEdgeToGraph("c", "e", 1)
// addEdgeToGraph("d", "c", -2)
// addEdgeToGraph("d", "f", -1)
// addEdgeToGraph("e", "g", 3)
// addEdgeToGraph("f", "g", 3)

// Graph with negative weight cycle
// addVertexToGraph("a")
// addVertexToGraph("b")
// addVertexToGraph("c")
// addVertexToGraph("d")
// addEdgeToGraph("a", "b", 4)
// addEdgeToGraph("a", "d", 5)
// addEdgeToGraph("b", "d", 5)
// addEdgeToGraph("c", "b", -10)
// addEdgeToGraph("d", "c", 3)
