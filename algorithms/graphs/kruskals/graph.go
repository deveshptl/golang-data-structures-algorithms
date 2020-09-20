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

// graph is a data structure which will be holding a graph
type graph map[string][]GraphNode

// addVertexToGraph adds a vertex to graph
func (g graph) addVertexToGraph(vtx string) {
	if g[vtx] != nil {
		return
	}
	g[vtx] = make([]GraphNode, 0)
}

// addEdgeToGraph adds an edge to a graph
func (g graph) addEdgeToGraph(fromVtx, toVtx string, edgeValue float64) {
	if g[fromVtx] == nil { // check if initial vertex exists
		return
	}
	for i := range g[fromVtx] { // check if edge already exists
		if g[fromVtx][i].name == toVtx {
			return
		}
	}
	if g[toVtx] == nil { // create new destination vertext if it does not exists
		g[toVtx] = make([]GraphNode, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	g[fromVtx] = append(g[fromVtx], GraphNode{name: toVtx, value: edgeValue})
	g[toVtx] = append(g[toVtx], GraphNode{name: fromVtx, value: edgeValue})
	uniqueEdges = append(uniqueEdges, Edge{src: fromVtx, dest: toVtx, value: edgeValue})
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

// removeEdgeFromGraph removes an edge from graph
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

// DFS traverses a graph using dfs technique
func (g graph) DFS(start string) bool {
	visited := make(map[string]bool)
	parent := "-1"
	if g[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return false
	}
	check := g.dfsHelper(start, visited, parent)
	return check
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfsHelper(vtx string, visited map[string]bool, parent string) bool {
	visited[vtx] = true
	for i := range g[vtx] {
		if visited[g[vtx][i].name] && parent != g[vtx][i].name && parent != vtx {
			return true
		}
		if !visited[g[vtx][i].name] {
			parent = vtx
			check := g.dfsHelper(g[vtx][i].name, visited, parent)
			if check {
				return check
			}
		}
	}
	return false
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
