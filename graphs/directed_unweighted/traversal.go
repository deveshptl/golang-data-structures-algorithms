package main

import "fmt"

// DFS traverses a graph using dfs technique
func DFS(start string) []string {
	visited := make(map[string]bool)
	var result []string
	if graph[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return nil
	}
	result = dfsHelper(start, visited, result)
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
func BFS(start string) []string {

	if graph[start] == nil {
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
		queue = queue[1:len(queue)]         // remove element from the queue
		result = append(result, currentVtx) // push the current vertex in result

		for i := range graph[currentVtx] {
			if !visited[graph[currentVtx][i]] {
				visited[graph[currentVtx][i]] = true        // mark as visited
				queue = append(queue, graph[currentVtx][i]) // push neighbors of current vertex in the queue
			}
		}

	}

	return result
}
