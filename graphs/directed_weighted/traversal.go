package main

import "fmt"

// DFS traverses a graph using dfs technique
func DFS() ([]string, []int) {
	visited := make(map[string]bool)

	var result []string
	var weights []int

	if len(graph) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil, nil
	}
	for i := range graph {
		if !visited[i] {
			result, weights = dfsHelper(i, visited, result, weights)
		}
	}
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
func BFS() ([]string, []int) {

	if len(graph) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil, nil
	}

	visited := make(map[string]bool)
	var queue []string
	var result []string
	var currentVtx string
	var weights []int

	for i := range graph {
		if !visited[i] {
			visited[i] = true
			queue = append(queue, i)

			for len(queue) != 0 {
				currentVtx = queue[0]               // get the element from the queue
				queue = queue[1:len(queue)]         // remove element from the queue
				result = append(result, currentVtx) // push the current vertex in result

				for j := range graph[currentVtx] {
					if !visited[graph[currentVtx][j].name] {
						visited[graph[currentVtx][j].name] = true // mark as visited
						weights = append(weights, graph[currentVtx][j].value)
						queue = append(queue, graph[currentVtx][j].name) // push neighbors of current vertex in the queue
					}
				}

			}
		}
	}

	return result, weights
}
