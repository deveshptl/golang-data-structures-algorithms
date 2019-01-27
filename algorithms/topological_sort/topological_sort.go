package main

import "fmt"

func main() {
	g := make(graph)
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. RUN TOPOLOGICAL SORT")
		fmt.Println("4. EXIT")
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
			g.runTopological()
			break
		case 4:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (g graph) runTopological() {
	result, cycle := g.detectDFS()
	if cycle {
		fmt.Println("\n-- Can't run Topological Sort. The graph is cyclic. --")
	} else {
		fmt.Println("\nSuccessfully ran Topological Sort:")
		for i := range result {
			fmt.Print(" ", result[i])
		}
		fmt.Println("")
	}
}

// cyclic
// g.addVertexToGraph("0")
// g.addVertexToGraph("1")
// g.addVertexToGraph("2")
// g.addVertexToGraph("3")
// g.addVertexToGraph("4")
// g.addEdgeToGraph("0", "1")
// g.addEdgeToGraph("1", "2")
// g.addEdgeToGraph("2", "4")
// g.addEdgeToGraph("2", "3")
// g.addEdgeToGraph("3", "0")
// g.addEdgeToGraph("4", "2")

// acyclic
// g.addVertexToGraph("0")
// g.addVertexToGraph("1")
// g.addVertexToGraph("2")
// g.addVertexToGraph("3")
// g.addVertexToGraph("4")
// g.addEdgeToGraph("0", "1")
// g.addEdgeToGraph("1", "2")
// g.addEdgeToGraph("2", "3")
// g.addEdgeToGraph("4", "2")
