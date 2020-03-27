package main

import "fmt"

// Node is a single node in a graph list
type Node struct {
	name  string
	value int
}

var graph map[string][]Node

func init() {
	graph = make(map[string][]Node)
}

func addVertexToGraph(vtx string) {
	if graph[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	vertices[vtx] = len(vertices)
	graph[vtx] = make([]Node, 0)
}

func addEdgeToGraph(fromVtx, toVtx string, edgeValue int) {
	if graph[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range graph[fromVtx] { // check if edge already exists
		if graph[fromVtx][i].name == toVtx {
			fmt.Println("\n-- Edge from " + fromVtx + " to " + toVtx + " already exists. --")
			return
		}
	}
	if graph[toVtx] == nil { // create new destination vertex if it does not exists
		vertices[toVtx] = len(vertices)
		graph[toVtx] = make([]Node, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	graph[fromVtx] = append(graph[fromVtx], Node{name: toVtx, value: edgeValue})
	return
}

func addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	addVertexToGraph(vtxName)
}

func addEdge() {
	var fromVtx, toVtx string
	var edgeValue int
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d\n", &edgeValue)
	addEdgeToGraph(fromVtx, toVtx, edgeValue)
}
