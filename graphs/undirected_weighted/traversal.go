package main

import "fmt"

// DFS traverses a graph using dfs technique
func DFS(start string) ([]string, []int) {
	visited := make(map[string]bool)
	var result []string
	var weights []int
	if graph[start] == nil {
		fmt.Println("\n-- No vertex named " + start + " present in graph. --")
		return nil, nil
	}
	result, weights = dfsHelper(start, visited, result, weights)
	return result, weights
}

// dfsHelper recursively calls itself to solve dfs traversal
func dfsHelper(vtx string, visited map[string]bool, result []string, weights []int) ([]string, []int) {
	visited[vtx] = true
	result = append(result, vtx)
	for i := range graph[vtx] {
		if !visited[graph[vtx][i].name] {
			weights = append(weights, graph[vtx][i].value)
			result, weights = dfsHelper(graph[vtx][i].name, visited, result, weights)
		}
	}
	return result, weights
}

// BFS traverses a graph using dfs technique
func BFS(start string) ([]string, []int) {

	if graph[start] == nil {
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
		queue = queue[1:len(queue)]         // remove element from the queue
		result = append(result, currentVtx) // push the current vertex in result

		for i := range graph[currentVtx] {
			if !visited[graph[currentVtx][i].name] {
				visited[graph[currentVtx][i].name] = true // mark as visited
				weights = append(weights, graph[currentVtx][i].value)
				queue = append(queue, graph[currentVtx][i].name) // push neighbors of current vertex in the queue
			}
		}

	}

	return result, weights
}
