package main

import (
	"fmt"
	"math"
)

// TODO: replace 999 with infinity

// Dijkstra starts dijkstra algorithm on a given graph
func Dijkstra(start, finish string) {
	var smallest string
	for vertex := range graph {
		if vertex == start {
			distances[vertex] = 0
			insertNodeInQueue(vertex, 0)
		} else {
			distances[vertex] = math.Inf(1)
			insertNodeInQueue(vertex, math.Inf(1))
		}
		previous[vertex] = ""
	}

	for len(pqueue) != 0 {
		smallest = Dequeue().name
		if smallest == finish {
			// Build path here
			for i := range previous {
				j := i
				fmt.Print("\nCost: ", distances[i], " units")
				fmt.Print(", Path: ")
				for j != "" {
					if previous[j] == "" {
						fmt.Print(j, "\n")
					} else {
						fmt.Print(j, " <= ")
					}
					j = previous[j]
				}
			}
			return
		}
		if smallest != "" || distances[smallest] != math.Inf(1) {
			for neighbor := range graph[smallest] {
				nextNode := graph[smallest][neighbor]
				least := distances[smallest] + nextNode.weight
				if least < distances[nextNode.node] {
					distances[nextNode.node] = least
					previous[nextNode.node] = smallest
					insertNodeInQueue(nextNode.node, least)
				}
			}
		}
	}

}

var graph map[string][]Node
var pqueue []*QueueNode
var distances map[string]float64
var previous map[string]string

func init() {
	graph = make(map[string][]Node)
	distances = make(map[string]float64)
	previous = make(map[string]string)
}

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. RUN DIJKSTRA ALGORITHM")
		fmt.Println("4. EXIT")
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
			DijkstraInit()
			break
		case 4:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

// DijkstraInit will take necessary input values from user to run
// dijkstra algorithm
func DijkstraInit() {
	var start, finish string
	fmt.Print("Enter the initial node: ")
	fmt.Scanf("%s", &start)
	fmt.Print("Enter the destination node: ")
	fmt.Scanf("%s", &finish)
	Dijkstra(start, finish)
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
// addEdgeToGraph("a", "b", float64(4))
// addEdgeToGraph("a", "c", float64(2))
// addEdgeToGraph("b", "e", float64(3))
// addEdgeToGraph("c", "d", float64(2))
// addEdgeToGraph("c", "f", float64(4))
// addEdgeToGraph("d", "e", float64(3))
// addEdgeToGraph("d", "f", float64(1))
// addEdgeToGraph("e", "f", float64(1))
// Dijkstra("a", "e")
