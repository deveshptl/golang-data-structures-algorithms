package main

import "fmt"

// DFS traverses a graph using dfs technique
func (g graph) DFS(start string) []string {
	visited := make(map[string]bool)
	var result []string
	if g[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return nil
	}
	result = g.dfsHelper(start, visited, result)
	return result
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfsHelper(vtx string, visited map[string]bool, result []string) []string {
	visited[vtx] = true
	result = append(result, vtx)
	for i := range g[vtx] {
		if !visited[g[vtx][i]] {
			result = g.dfsHelper(g[vtx][i], visited, result)
		}
	}
	return result
}

// BFS traverses a graph using dfs technique
func (g graph) BFS(start string) []string {

	if g[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return nil
	}

	visited := make(map[string]bool)

	var queue []string
	queue = append(queue, start)

	var result []string
	var currentVtx string
	visited[start] = true

	for len(queue) != 0 {
		currentVtx = queue[0]               // get the element from the queue
		queue = queue[1:]                   // remove element from the queue
		result = append(result, currentVtx) // push the current vertex in result

		for i := range g[currentVtx] {
			if !visited[g[currentVtx][i]] {
				visited[g[currentVtx][i]] = true        // mark as visited
				queue = append(queue, g[currentVtx][i]) // push neighbors of current vertex in the queue
			}
		}

	}

	return result
}
