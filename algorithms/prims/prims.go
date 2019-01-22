package main

import (
	"fmt"
	"math"
)

// Prims will run prim's algorithm on a given graph
func (g graph) Prims(start string) {
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return
	}
	for len(heap) != 0 {
		node := Dequeue()
		mstSet[node.name] = true
		for _, graphNode := range g[node.name] {
			if (graphNode.value < vertexWeight[graphNode.name]) && !mstSet[graphNode.name] {
				DecreaseKey(hashMap[graphNode.name], graphNode.value)
				vertexWeight[graphNode.name] = graphNode.value
			}
		}
	}
	g.printPrims(start)
}

var mstSet map[string]bool
var vertexWeight map[string]float64

func init() {
	mstSet = make(map[string]bool)
	vertexWeight = make(map[string]float64)
}

func main() {
	g := make(graph)
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
			g.addVertex()
			break
		case 2:
			g.addEdge()
			break
		case 3:
			g.startPrims()
			break
		case 4:
			g.simpleDisplay()
			break
		case 5:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (g graph) startPrims() {
	var start string
	fmt.Print("Enter the source vertex: ")
	fmt.Scanf("%s", &start)
	for i := range g {
		if i == start {
			insertNode(i, 0)
			vertexWeight[i] = 0
		} else {
			insertNode(i, math.Inf(0))
			vertexWeight[i] = math.Inf(0)
		}
		mstSet[i] = false
	}
	g.Prims(start)
}

func (g graph) printPrims(start string) {
	for i := range g {
		for _, node := range g[i] {
			if node.value == vertexWeight[node.name] {
				fmt.Println(i, "-", node.name, "=>", vertexWeight[node.name])
			}
		}
	}
}

// g.addVertexToGraph("0")
// g.addVertexToGraph("1")
// g.addVertexToGraph("2")
// g.addVertexToGraph("3")
// g.addVertexToGraph("4")
// g.addVertexToGraph("5")
// g.addVertexToGraph("6")
// g.addVertexToGraph("7")
// g.addVertexToGraph("8")
// g.addEdgeToGraph("0", "1", 4)
// g.addEdgeToGraph("0", "7", 8)
// g.addEdgeToGraph("1", "2", 8)
// g.addEdgeToGraph("1", "7", 11)
// g.addEdgeToGraph("2", "3", 7)
// g.addEdgeToGraph("2", "5", 4)
// g.addEdgeToGraph("2", "8", 2)
// g.addEdgeToGraph("3", "4", 9)
// g.addEdgeToGraph("3", "5", 14)
// g.addEdgeToGraph("4", "5", 10)
// g.addEdgeToGraph("5", "6", 2)
// g.addEdgeToGraph("6", "8", 6)
// g.addEdgeToGraph("6", "7", 1)
// g.addEdgeToGraph("7", "8", 7)
