package main

import "fmt"

// DFS traverses a graph using dfs technique
func DFS() []string {
	visited := make(map[string]bool)
	var result []string
	if len(graph) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil
	}
	for i := range graph {
		if !visited[i] {
			result = dfsHelper(i, visited, result)
		}
	}
	return result
}

// dfsHelper recursively calls itself to solve dfs traversal
func dfsHelper(vtx string, visited map[string]bool, result []string) []string {
	visited[vtx] = true
	result = append(result, vtx)
	for i := range graph[vtx] {
		if !visited[graph[vtx][i]] {
			result = dfsHelper(graph[vtx][i], visited, result)
		}
	}
	return result
}

// BFS traverses a graph using dfs technique
func BFS() []string {

	if len(graph) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil
	}

	visited := make(map[string]bool)

	var queue []string
	var result []string
	var currentVtx string

	for i := range graph {
		if !visited[i] {
			visited[i] = true
			queue = append(queue, i)

			for len(queue) != 0 {
				currentVtx = queue[0]               // get the element from the queue
				queue = queue[1:len(queue)]         // remove element from the queue
				result = append(result, currentVtx) // push the current vertex in result

				for j := range graph[currentVtx] {
					if !visited[graph[currentVtx][j]] {
						visited[graph[currentVtx][j]] = true        // mark as visited
						queue = append(queue, graph[currentVtx][j]) // push neighbors of current vertex in the queue
					}
				}

			}

		}
	}

	return result
}
