package main

import (
	"fmt"
	"math"
)

// Prims will run prim's algorithm on a given graph
func Prims(start string) {
	if len(graph) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return
	}
	for len(heap) != 0 {
		node := Dequeue()
		mstSet[node.name] = true
		for _, graphNode := range graph[node.name] {
			if (float64(graphNode.value) < vertexWeight[graphNode.name]) && !mstSet[graphNode.name] {
				DecreaseKey(hashMap[graphNode.name], float64(graphNode.value))
				vertexWeight[graphNode.name] = float64(graphNode.value)
			}
		}
	}
	printPrims(start)
}

var mstSet map[string]bool
var vertexWeight map[string]float64

func init() {
	mstSet = make(map[string]bool)
	vertexWeight = make(map[string]float64)
}

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. RUN PRIM's ALGORITHM")
		fmt.Println("4. SIMPLE DISPLAY")
		fmt.Println("5. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			addVertex()
			break
		case 2:
			addEdge()
			break
		case 3:
			startPrims()
			break
		case 4:
			simpleDisplay()
			break
		case 5:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func startPrims() {
	var start string
	fmt.Print("Enter the source vertex: ")
	fmt.Scanf("%s", &start)
	for i := range graph {
		if i == start {
			insertNode(i, 0)
			vertexWeight[i] = 0
		} else {
			insertNode(i, math.Inf(0))
			vertexWeight[i] = math.Inf(0)
		}
		mstSet[i] = false
	}
	Prims(start)
}

func printPrims(start string) {
	for i := range graph {
		for _, node := range graph[i] {
			if float64(node.value) == vertexWeight[node.name] {
				fmt.Println(i, "-", node.name, "=>", vertexWeight[node.name])
			}
		}
	}
}

// addVertexToGraph("0")
// addVertexToGraph("1")
// addVertexToGraph("2")
// addVertexToGraph("3")
// addVertexToGraph("4")
// addVertexToGraph("5")
// addVertexToGraph("6")
// addVertexToGraph("7")
// addVertexToGraph("8")
// addEdgeToGraph("0", "1", 4)
// addEdgeToGraph("0", "7", 8)
// addEdgeToGraph("1", "2", 8)
// addEdgeToGraph("1", "7", 11)
// addEdgeToGraph("2", "3", 7)
// addEdgeToGraph("2", "5", 4)
// addEdgeToGraph("2", "8", 2)
// addEdgeToGraph("3", "4", 9)
// addEdgeToGraph("3", "5", 14)
// addEdgeToGraph("4", "5", 10)
// addEdgeToGraph("5", "6", 2)
// addEdgeToGraph("6", "8", 6)
// addEdgeToGraph("6", "7", 1)
// addEdgeToGraph("7", "8", 7)
