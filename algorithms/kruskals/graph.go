package main

import "fmt"

// GraphNode is a single node in a graph list
type GraphNode struct {
	name  string
	value float64
}

// Edge represents an edge between two vertices
type Edge struct {
	src   string
	dest  string
	value float64
}

func addVertexToGraph(graph map[string][]GraphNode, vtx string) {
	if graph[vtx] != nil {
		return
	}
	graph[vtx] = make([]GraphNode, 0)
}

func addEdgeToGraph(graph map[string][]GraphNode, fromVtx, toVtx string, edgeValue float64) {
	if graph[fromVtx] == nil { // check if initial vertex exists
		return
	}
	for i := range graph[fromVtx] { // check if edge already exists
		if graph[fromVtx][i].name == toVtx {
			return
		}
	}
	if graph[toVtx] == nil { // create new destination vertext if it does not exists
		graph[toVtx] = make([]GraphNode, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	graph[fromVtx] = append(graph[fromVtx], GraphNode{name: toVtx, value: edgeValue})
	graph[toVtx] = append(graph[toVtx], GraphNode{name: fromVtx, value: edgeValue})
	uniqueEdges = append(uniqueEdges, Edge{src: fromVtx, dest: toVtx, value: edgeValue})
	return
}

func addVertex(graph map[string][]GraphNode) {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s", &vtxName)
	addVertexToGraph(graph, vtxName)
}

func addEdge(graph map[string][]GraphNode) {
	var fromVtx, toVtx string
	var edgeValue float64
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d", &edgeValue)
	addEdgeToGraph(graph, fromVtx, toVtx, edgeValue)
}

func removeEdgeFromGraph(graph map[string][]GraphNode, fromVtx, toVtx string) {
	if graph[fromVtx] == nil || graph[toVtx] == nil {
		fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " does not exist. --")
		return
	}

	for i := range graph[fromVtx] {
		if graph[fromVtx][i].name == toVtx {
			if i == 0 {
				graph[fromVtx] = graph[fromVtx][1:len(graph[fromVtx])]
			} else if i == (len(graph[fromVtx]) - 1) {
				graph[fromVtx] = graph[fromVtx][0:(len(graph[fromVtx]) - 1)]
			} else {
				initial := graph[fromVtx][0:i]
				final := graph[fromVtx][i+1 : len(graph[fromVtx])]
				graph[fromVtx] = append(initial, final...)
			}
			break
		}
	}

	for i := range graph[toVtx] {
		if graph[toVtx][i].name == fromVtx {
			if i == 0 {
				graph[toVtx] = graph[toVtx][1:len(graph[toVtx])]
			} else if i == (len(graph[toVtx]) - 1) {
				graph[toVtx] = graph[toVtx][0:(len(graph[toVtx]) - 1)]
			} else {
				initial := graph[toVtx][0:i]
				final := graph[toVtx][i+1 : len(graph[toVtx])]
				graph[toVtx] = append(initial, final...)
			}
			break
		}
	}
}

// DFS traverses a graph using dfs technique
func DFS(graph map[string][]GraphNode, start string) bool {
	visited := make(map[string]bool)
	parent := "-1"
	if graph[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return false
	}
	check := dfsHelper(graph, start, visited, parent)
	return check
}

// dfsHelper recursively calls itself to solve dfs traversal
func dfsHelper(graph map[string][]GraphNode, vtx string, visited map[string]bool, parent string) bool {
	visited[vtx] = true
	for i := range graph[vtx] {
		if visited[graph[vtx][i].name] && parent != graph[vtx][i].name && parent != vtx {
			return true
		}
		if !visited[graph[vtx][i].name] {
			parent = vtx
			check := dfsHelper(graph, graph[vtx][i].name, visited, parent)
			if check {
				return check
			}
		}
	}
	return false
}

func simpleDisplay(graph map[string][]GraphNode) {
	fmt.Println("")
	for i := range graph {
		fmt.Print(i, " => ")
		for j := range graph[i] {
			fmt.Print(graph[i][j])
		}
		fmt.Println("")
	}
}
