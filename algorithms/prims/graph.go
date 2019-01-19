package main

import "fmt"

// GraphNode is a single node in a graph list
type GraphNode struct {
	name  string
	value float64
}

func addVertexToGraph(vtx string) {
	if graph[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	graph[vtx] = make([]GraphNode, 0)
}

func addEdgeToGraph(fromVtx, toVtx string, edgeValue float64) {
	if graph[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range graph[fromVtx] { // check if edge already exists
		if graph[fromVtx][i].name == toVtx {
			fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " already exists. --")
			return
		}
	}
	if graph[toVtx] == nil { // create new destination vertext if it does not exists
		graph[toVtx] = make([]GraphNode, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	graph[fromVtx] = append(graph[fromVtx], GraphNode{name: toVtx, value: edgeValue})
	graph[toVtx] = append(graph[toVtx], GraphNode{name: fromVtx, value: edgeValue})
	return
}

var graph map[string][]GraphNode

func init() {
	graph = make(map[string][]GraphNode)
}

func addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s", &vtxName)
	addVertexToGraph(vtxName)
}

func addEdge() {
	var fromVtx, toVtx string
	var edgeValue float64
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d", &edgeValue)
	addEdgeToGraph(fromVtx, toVtx, edgeValue)
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
