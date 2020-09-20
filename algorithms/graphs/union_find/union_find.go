package main

import "fmt"

// UnionFind run union find algorithm and detects cycle in give graph
func UnionFind() {
	flag := 0
	for i := 0; i < len(uniqueEdges); i++ {
		x := find(uniqueEdges[i].src)
		y := find(uniqueEdges[i].dest)
		if x == y {
			fmt.Println("\n-- Cycle FOUND between:", uniqueEdges[i].src, uniqueEdges[i].dest, ". --")
			flag = 1
		} else {
			union(x, y)
		}
	}
	if flag == 0 {
		fmt.Println("\n-- NO cycle found in graph. --")
	}
}

func find(i string) string {
	if parent[i] == "-1" {
		return i
	}
	return find(parent[i])
}

func union(vtx1, vtx2 string) {
	xset := find(vtx1)
	yset := find(vtx2)
	parent[xset] = yset
}

var parent map[string]string
var vertices map[string]int

func init() {
	vertices = make(map[string]int)
	parent = make(map[string]string)
}

func main() {
	fmt.Println("\n-- Union Find --")
	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. RUN UNION FIND")
		fmt.Println("4. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			addVertex()
		case 2:
			addEdge()
		case 3:
			UnionFind()
		case 4:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func startUnionFind() {
	for i := range vertices {
		parent[i] = "-1"
	}
	UnionFind()
}

// addVertexToGraph("0")
// addVertexToGraph("1")
// addVertexToGraph("2")
// addVertexToGraph("3")
// addVertexToGraph("4")
// addEdgeToGraph("0", "1")
// addEdgeToGraph("0", "2")
// addEdgeToGraph("1", "3")
// addEdgeToGraph("1", "4")
// addEdgeToGraph("3", "4")

// addVertexToGraph("1")
// addVertexToGraph("2")
// addVertexToGraph("3")
// addVertexToGraph("4")
// addVertexToGraph("5")
// addVertexToGraph("6")
// addVertexToGraph("7")
// addVertexToGraph("8")
// addEdgeToGraph("1", "3")
// addEdgeToGraph("1", "2")
// addEdgeToGraph("2", "4")
// addEdgeToGraph("3", "4")
// addEdgeToGraph("2", "5")
// addEdgeToGraph("5", "7")
// addEdgeToGraph("5", "6")
// addEdgeToGraph("8", "6")
// addEdgeToGraph("8", "7")

// addVertexToGraph("0")
// addVertexToGraph("1")
// addVertexToGraph("2")
// addEdgeToGraph("0", "1")
// addEdgeToGraph("1", "2")
// addEdgeToGraph("0", "2")
