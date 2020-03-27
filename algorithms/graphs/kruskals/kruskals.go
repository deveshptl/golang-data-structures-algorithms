package main

import "fmt"

// Kruskals will run kruskal's algorithm on a given graph
func (g graph) Kruskals(start string) {
	count := 0
	mst := make(graph)
	for count != len(g)-1 {
		node := Dequeue()
		mst.addVertexToGraph(string(node.name[0]))
		mst.addVertexToGraph(string(node.name[1]))
		mst.addEdgeToGraph(string(node.name[0]), string(node.name[1]), node.priority)
		isCyclic := mst.DFS(string(node.name[0]))
		if isCyclic {
			mst.removeEdgeFromGraph(string(node.name[0]), string(node.name[1]))
			count--
		}
		count++
	}
	mst.printKruskals(start)
}

var uniqueEdges []Edge

func init() {
	uniqueEdges = make([]Edge, 0)
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
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			g.addVertex()
			break
		case 2:
			g.addEdge()
			break
		case 3:
			g.startKruskals()
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

func (g graph) startKruskals() {
	var start string
	fmt.Print("Enter the source vertex: ")
	fmt.Scanf("%s\n", &start)
	for _, edge := range uniqueEdges {
		insertNode(edge.src+edge.dest, edge.value)
	}
	g.Kruskals(start)
}

func (g graph) printKruskals(start string) {
	visited := make(map[string]bool)
	for i := range g {
		for _, node := range g[i] {
			if !visited[node.name+i] {
				fmt.Println(i, "-", node.name, "=>", node.value)
			}
			visited[node.name+i] = true
			visited[i+node.name] = true
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
// g.startKruskals()
