package main

import "fmt"

// Kruskals will run kruskal's algorithm on a given graph
func Kruskals(graph map[string][]GraphNode, start string) {
	count := 0
	mst := make(map[string][]GraphNode)
	for count != len(graph)-1 {
		node := Dequeue()
		addVertexToGraph(mst, string(node.name[0]))
		addVertexToGraph(mst, string(node.name[1]))
		addEdgeToGraph(mst, string(node.name[0]), string(node.name[1]), node.priority)
		isCyclic := DFS(mst, string(node.name[0]))
		if isCyclic {
			removeEdgeFromGraph(mst, string(node.name[0]), string(node.name[1]))
			count--
		}
		count++
	}
	printKruskals(mst, start)
}

var uniqueEdges []Edge

func init() {
	uniqueEdges = make([]Edge, 0)
}

func main() {
	graph := make(map[string][]GraphNode)
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
			addVertex(graph)
			break
		case 2:
			addEdge(graph)
			break
		case 3:
			startKruskals(graph)
			break
		case 4:
			simpleDisplay(graph)
			break
		case 5:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func startKruskals(graph map[string][]GraphNode) {
	var start string
	fmt.Print("Enter the source vertex: ")
	fmt.Scanf("%s", &start)
	for _, edge := range uniqueEdges {
		insertNode(edge.src+edge.dest, edge.value)
	}
	Kruskals(graph, start)
}

func printKruskals(mst map[string][]GraphNode, start string) {
	visited := make(map[string]bool)
	for i := range mst {
		for _, node := range mst[i] {
			if !visited[node.name+i] {
				fmt.Println(i, "-", node.name, "=>", node.value)
			}
			visited[node.name+i] = true
			visited[i+node.name] = true
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
