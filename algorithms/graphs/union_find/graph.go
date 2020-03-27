package main

import "fmt"

// Edge represents an edge between two vertices
type Edge struct {
	src  string
	dest string
}

var graph map[string][]string
var uniqueEdges []Edge

func init() {
	graph = make(map[string][]string)
	uniqueEdges = make([]Edge, 0)
}

func addVertexToGraph(vtx string) {
	if graph[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	vertices[vtx] = len(vertices)
	graph[vtx] = make([]string, 0)
}

func addEdgeToGraph(fromVtx, toVtx string) {
	if graph[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range graph[fromVtx] { // check if edge already exists
		if graph[fromVtx][i] == toVtx {
			fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " already exists. --")
			return
		}
	}
	if graph[toVtx] == nil { // create new destination vertext if it does not exists
		graph[toVtx] = make([]string, 0)
		vertices[toVtx] = len(vertices)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	graph[fromVtx] = append(graph[fromVtx], toVtx)
	graph[toVtx] = append(graph[toVtx], fromVtx)
	uniqueEdges = append(uniqueEdges, Edge{src: fromVtx, dest: toVtx})
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
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	addEdgeToGraph(fromVtx, toVtx)
}
