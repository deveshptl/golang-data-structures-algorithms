package main

import "fmt"

// DFS traverses a graph using dfs technique
func (g graph) DFS(start string) ([]string, []int) {
	visited := make(map[string]bool)
	var result []string
	var weights []int
	if g[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return nil, nil
	}
	result, weights = g.dfsHelper(start, visited, result, weights)
	return result, weights
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfsHelper(vtx string, visited map[string]bool, result []string, weights []int) ([]string, []int) {
	visited[vtx] = true
	result = append(result, vtx)
	for i := range g[vtx] {
		if !visited[g[vtx][i].name] {
			weights = append(weights, g[vtx][i].value)
			result, weights = g.dfsHelper(g[vtx][i].name, visited, result, weights)
		}
	}
	return result, weights
}

// BFS traverses a graph using dfs technique
func (g graph) BFS(start string) ([]string, []int) {

	if g[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return nil, nil
	}

	visited := make(map[string]bool)

	var queue []string
	queue = append(queue, start)

	var result []string
	var weights []int
	var currentVtx string
	visited[start] = true

	for len(queue) != 0 {
		currentVtx = queue[0]               // get the element from the queue
		queue = queue[1:]                   // remove element from the queue
		result = append(result, currentVtx) // push the current vertex in result

		for i := range g[currentVtx] {
			if !visited[g[currentVtx][i].name] {
				visited[g[currentVtx][i].name] = true // mark as visited
				weights = append(weights, g[currentVtx][i].value)
				queue = append(queue, g[currentVtx][i].name) // push neighbors of current vertex in the queue
			}
		}

	}

	return result, weights
}
