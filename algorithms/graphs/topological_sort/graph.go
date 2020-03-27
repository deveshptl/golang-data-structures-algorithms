package main

import "fmt"

type graph map[string][]string

func (g graph) addVertexToGraph(vtx string) {
	if g[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	g[vtx] = make([]string, 0)
}

func (g graph) addEdgeToGraph(fromVtx, toVtx string) {
	if g[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range g[fromVtx] { // check if edge already exists
		if g[fromVtx][i] == toVtx {
			fmt.Println("\n-- Edge from " + fromVtx + " to " + toVtx + " already exists. --")
			return
		}
	}
	if g[toVtx] == nil { // create new destination vertext if it does not exists
		g[toVtx] = make([]string, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	g[fromVtx] = append(g[fromVtx], toVtx)
	return
}

// detectDFS traverses a graph using dfs technique to detect a cycle in a graph
func (g graph) detectDFS() ([]string, bool) {
	visited := make(map[string]bool)
	stack := make(map[string]string) // keeps track of stack
	var cycle bool
	var result []string
	var finalResult []string
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil, false
	}
	for i := range g {
		if !visited[i] {
			stack[i] = i
			result, cycle = g.dfsHelper(i, visited, result, stack)

			// appending the result into finalResult in reverse order making topological sort
			for i := len(result) - 1; i >= 0; i-- {
				finalResult = append(finalResult, result[i])
			}
			result = make([]string, 0)

			delete(stack, i)
			if cycle {
				return nil, cycle
			}
		}
	}
	return finalResult, cycle
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfsHelper(vtx string, visited map[string]bool, result []string, stack map[string]string) ([]string, bool) {
	visited[vtx] = true
	cycle := false
	result = append(result, vtx)
	for i := range g[vtx] {
		if stack[g[vtx][i]] == g[vtx][i] { // cycle detection logic is here
			return nil, true
		}
		if !visited[g[vtx][i]] {
			stack[g[vtx][i]] = g[vtx][i]
			result, cycle = g.dfsHelper(g[vtx][i], visited, result, stack)
			delete(stack, g[vtx][i])
			return result, cycle
		}
	}
	return result, cycle
}

func (g graph) addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	g.addVertexToGraph(vtxName)
}

func (g graph) addEdge() {
	var fromVtx, toVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	g.addEdgeToGraph(fromVtx, toVtx)
}
