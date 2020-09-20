package main

import "fmt"

// detectDFS traverses a graph using dfs technique to detect cycle in a graph
func (g graph) detectDFS() ([]string, bool) {
	visited := make(map[string]bool)
	cycle := false
	var result []string
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil, false
	}
	for i := range g {
		if !visited[i] {
			result, cycle = g.dfsHelper(i, visited, result, "-1")
			if cycle {
				return nil, cycle
			}
		}
	}
	return result, cycle
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfsHelper(vtx string, visited map[string]bool, result []string, parent string) ([]string, bool) {
	visited[vtx] = true
	cycle := false
	result = append(result, vtx)
	for i := range g[vtx] {
		fmt.Println(parent, vtx, g[vtx][i])
		if parent != g[vtx][i] && parent != "-1" && visited[g[vtx][i]] { // cycle detection logic here
			return nil, true
		}
		if !visited[g[vtx][i]] {
			result, cycle = g.dfsHelper(g[vtx][i], visited, result, vtx)
			return result, cycle
		}
	}
	return result, cycle
}

// TODO:
// detectBFS traverses a graph using dfs technique to detect cycle in a graph
func (g graph) detectBFS() []string {

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
				queue = queue[1:]                   // remove element from the queue
				result = append(result, currentVtx) // push the current vertex in result

				for i := range g[currentVtx] {
					if !visited[g[currentVtx][i]] {
						visited[g[currentVtx][i]] = true        // mark as visited
						queue = append(queue, g[currentVtx][i]) // push neighbors of current vertex in the queue
					}
				}

			}

		}
	}

	return result
}
