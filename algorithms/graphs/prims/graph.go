package main

import "fmt"

// GraphNode is a node in a graph list
type GraphNode struct {
	name  string
	value float64
}

// graph is a data structure which will be holding a graph
type graph map[string][]GraphNode

// addVertexToGraph adds a vertex to graph
func (g graph) addVertexToGraph(vtx string) {
	if g[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	g[vtx] = make([]GraphNode, 0)
}

// addEdgeToGraph adds an edge to a graph
func (g graph) addEdgeToGraph(fromVtx, toVtx string, edgeValue float64) {
	if g[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range g[fromVtx] { // check if edge already exists
		if g[fromVtx][i].name == toVtx {
			fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " already exists. --")
			return
		}
	}
	if g[toVtx] == nil { // create new destination vertext if it does not exists
		g[toVtx] = make([]GraphNode, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	g[fromVtx] = append(g[fromVtx], GraphNode{name: toVtx, value: edgeValue})
	g[toVtx] = append(g[toVtx], GraphNode{name: fromVtx, value: edgeValue})
	}

// addVertex asks user to enter vertex name
func (g graph) addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	g.addVertexToGraph(vtxName)
}

// addEdge asks user to enter necessary values before adding an edge to graph
func (g graph) addEdge() {
	var fromVtx, toVtx string
	var edgeValue float64
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d\n", &edgeValue)
	g.addEdgeToGraph(fromVtx, toVtx, edgeValue)
}

// simpleDisplay simply displays a graph
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
