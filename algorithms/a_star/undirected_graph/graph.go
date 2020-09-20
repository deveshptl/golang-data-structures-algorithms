package main

import "fmt"

// Node is a single node in a graph list
type Node struct {
	name  string
	value int
}

type graph map[string][]Node

var vtxWeights map[string]int

func (g graph) addVertexToGraph(vtx string, vtxWeight int) {
	if g[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	vtxWeights[vtx] = vtxWeight
	g[vtx] = make([]Node, 0)
}

func (g graph) addEdgeToGraph(fromVtx, toVtx string, edgeValue int) {
	if g[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}

	if g[toVtx] == nil { // check if destination vertex exists
		fmt.Println("\n-- Destination vertex " + toVtx + " does not exist. --")
		return
	}

	for i := range g[fromVtx] { // check if edge already exists
		if g[fromVtx][i].name == toVtx {
			fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " already exists. --")
			return
		}
	}

	g[fromVtx] = append(g[fromVtx], Node{name: toVtx, value: edgeValue})
	g[toVtx] = append(g[toVtx], Node{name: fromVtx, value: edgeValue})
}

func (g graph) removeVertexFromGraph(vtx string) {
	length := len(g[vtx]) - 1
	for length != -1 {
		g.removeEdgeFromGraph(vtx, g[vtx][length].name)
		length--
	}
	delete(g, vtx)
	delete(vtxWeights, vtx)
}

func (g graph) removeEdgeFromGraph(fromVtx, toVtx string) {
	if g[fromVtx] == nil || g[toVtx] == nil {
		fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " does not exist. --")
		return
	}

	for i := range g[fromVtx] {
		if g[fromVtx][i].name == toVtx {
			if i == 0 {
				g[fromVtx] = g[fromVtx][1:len(g[fromVtx])]
			} else if i == (len(g[fromVtx]) - 1) {
				g[fromVtx] = g[fromVtx][0:(len(g[fromVtx]) - 1)]
			} else {
				initial := g[fromVtx][0:i]
				final := g[fromVtx][i+1 : len(g[fromVtx])]
				g[fromVtx] = append(initial, final...)
			}
			break
		}
	}

	for i := range g[toVtx] {
		if g[toVtx][i].name == fromVtx {
			if i == 0 {
				g[toVtx] = g[toVtx][1:len(g[toVtx])]
			} else if i == (len(g[toVtx]) - 1) {
				g[toVtx] = g[toVtx][0:(len(g[toVtx]) - 1)]
			} else {
				initial := g[toVtx][0:i]
				final := g[toVtx][i+1 : len(g[toVtx])]
				g[toVtx] = append(initial, final...)
			}
			break
		}
	}
}

func main() {
	i := 0
	g := make(graph)
	vtxWeights = make(map[string]int)
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. REMOVE VERTEX")
		fmt.Println("4. REMOVE AN EDGE")
		fmt.Println("5. RUN A STAR ALGO")
		fmt.Println("6. DISPLAY GRAPH")
		fmt.Println("7. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			g.addVertex()
		case 2:
			g.addEdge()
		case 3:
			g.removeVertex()
		case 4:
			g.removeEdge()
		case 5:
			g.initAStarAlgo()
		case 6:
			g.simpleDisplay()
		case 7:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (g graph) addVertex() {
	var vtxName string
	var vtxWeight int
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	fmt.Print("Enter the weight of vertex: ")
	fmt.Scanf("%d\n", &vtxWeight)
	g.addVertexToGraph(vtxName, vtxWeight)
}

func (g graph) addEdge() {
	var fromVtx, toVtx string
	var edgeValue int
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d\n", &edgeValue)
	g.addEdgeToGraph(fromVtx, toVtx, edgeValue)
}

func (g graph) removeVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	g.removeVertexFromGraph(vtxName)
}

func (g graph) removeEdge() {
	var fromVtx, toVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	g.removeEdgeFromGraph(fromVtx, toVtx)
}

func (g graph) initAStarAlgo() {
	var initialVtx, goalVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &initialVtx)
	fmt.Print("Enter the goal vertex name: ")
	fmt.Scanf("%s\n", &goalVtx)
	if g[initialVtx] == nil {
		fmt.Println("Initial vertex does not exist.")
		return
	} else if g[goalVtx] == nil {
		fmt.Println("Initial vertex does not exist.")
		return
	}
	g.runAStarAlgo(initialVtx, goalVtx)
}

func (g graph) simpleDisplay() {
	fmt.Println("")
	for i := range g {
		fmt.Print(i, " => ")
		for j := range g[i] {
			fmt.Print(g[i][j])
		}
		fmt.Println("")
	}
}

// g.addVertexToGraph("a", 10)
// g.addVertexToGraph("b", 8)
// g.addVertexToGraph("c", 5)
// g.addVertexToGraph("d", 7)
// g.addVertexToGraph("e", 3)
// g.addVertexToGraph("f", 6)
// g.addVertexToGraph("g", 5)
// g.addVertexToGraph("h", 3)
// g.addVertexToGraph("i", 1)
// g.addVertexToGraph("j", 0)
// g.addEdgeToGraph("a", "b", 6)
// g.addEdgeToGraph("a", "f", 3)
// g.addEdgeToGraph("b", "c", 3)
// g.addEdgeToGraph("b", "d", 2)
// g.addEdgeToGraph("c", "d", 1)
// g.addEdgeToGraph("c", "e", 5)
// g.addEdgeToGraph("d", "e", 8)
// g.addEdgeToGraph("e", "i", 5)
// g.addEdgeToGraph("e", "j", 5)
// g.addEdgeToGraph("f", "g", 1)
// g.addEdgeToGraph("f", "h", 7)
// g.addEdgeToGraph("g", "i", 3)
// g.addEdgeToGraph("h", "i", 2)
// g.addEdgeToGraph("i", "j", 3)
// g.runAStarAlgo("a", "j")
// a-f-g-i-j 10

// g.addVertexToGraph("a", 14)
// g.addVertexToGraph("b", 12)
// g.addVertexToGraph("c", 11)
// g.addVertexToGraph("d", 6)
// g.addVertexToGraph("e", 4)
// g.addVertexToGraph("f", 11)
// g.addVertexToGraph("z", 0)
// g.addEdgeToGraph("a", "b", 4)
// g.addEdgeToGraph("a", "c", 3)
// g.addEdgeToGraph("b", "e", 12)
// g.addEdgeToGraph("b", "f", 5)
// g.addEdgeToGraph("c", "d", 7)
// g.addEdgeToGraph("c", "e", 10)
// g.addEdgeToGraph("d", "e", 2)
// g.addEdgeToGraph("e", "z", 5)
// g.addEdgeToGraph("f", "z", 16)
// g.runAStarAlgo("a", "z")
// a-c-d-e-z 17
