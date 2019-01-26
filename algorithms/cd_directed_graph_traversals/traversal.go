package main

import "fmt"

// detectDFS traverses a graph using dfs technique to detect a cycle in a graph
func (g graph) detectDFS() ([]string, bool) {
	visited := make(map[string]bool)
	stack := make(map[string]string) // keeps track of stack
	var cycle bool
	var result []string
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil, false
	}
	for i := range g {
		if !visited[i] {
			stack[i] = i
			result, cycle = g.dfsHelper(i, visited, result, stack)
			delete(stack, i)
			if cycle {
				return nil, cycle
			}
		}
	}
	return result, cycle
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

// TODO:
// detectBFS traverses a graph using dfs technique to detect cycle using BFS
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

func (g graph) simpleDisplay() {
	fmt.Println("")
	for i := range g {
		fmt.Print(i, " => ")
		for j := range g[i] {
			fmt.Print(g[i][j] + " ")
		}
		fmt.Println("")
	}
}
