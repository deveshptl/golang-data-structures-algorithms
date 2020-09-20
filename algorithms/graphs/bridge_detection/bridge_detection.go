package main

import "fmt"

func main() {
	g := make(graph)
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. DETECT BRIDGES")
		fmt.Println("4. SIMPLE DISPLAY")
		fmt.Println("5. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			g.addVertex()
		case 2:
			g.addEdge()
		case 3:
			result := g.findBridges()
			prev := ""
			if len(result) != 0 {
				fmt.Println("\n-- Bridge/s exists in a given graph. --")
				for i := range result {
					if prev != "" {
						fmt.Println("-- Bridge exists between:", prev, "-", result[i])
						prev = ""
					} else {
						prev = result[i]
					}
				}
			} else {
				fmt.Println("\n -- No bridge exists in a given graph. --")
			}
		case 4:
			g.simpleDisplay()
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

// c-d, c-h, h-i
// g.addVertexToGraph("a")
// g.addVertexToGraph("b")
// g.addVertexToGraph("c")
// g.addVertexToGraph("d")
// g.addVertexToGraph("e")
// g.addVertexToGraph("f")
// g.addVertexToGraph("g")
// g.addVertexToGraph("h")
// g.addVertexToGraph("i")
// g.addEdgeToGraph("a", "b")
// g.addEdgeToGraph("b", "c")
// g.addEdgeToGraph("c", "a")
// g.addEdgeToGraph("c", "d")
// g.addEdgeToGraph("d", "e")
// g.addEdgeToGraph("e", "f")
// g.addEdgeToGraph("f", "g")
// g.addEdgeToGraph("g", "d")
// g.addEdgeToGraph("c", "h")
// g.addEdgeToGraph("h", "i")
