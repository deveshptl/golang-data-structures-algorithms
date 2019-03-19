package main

import "fmt"

// DFS traverses a graph using dfs technique
func (g graph) DFS() []string {
	visited := make(map[string]bool)
	var result []string
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil
	}
	for i := range g {
		if !visited[i] {
			result = g.dfsHelper(i, visited, result)
		}
	}
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
func (g graph) BFS() []string {

	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil
	}

	visited := make(map[string]bool)

	var queue []string
	var result []string
	var currentVtx string

	for i := range g {
		if !visited[i] {
			visited[i] = true
			queue = append(queue, i)

			for len(queue) != 0 {
				currentVtx = queue[0]               // get the element from the queue
				queue = queue[1:len(queue)]         // remove element from the queue
				result = append(result, currentVtx) // push the current vertex in result

				for j := range g[currentVtx] {
					if !visited[g[currentVtx][j]] {
						visited[g[currentVtx][j]] = true        // mark as visited
						queue = append(queue, g[currentVtx][j]) // push neighbors of current vertex in the queue
					}
				}

			}

		}
	}

	return result
}
