package main

import (
	"fmt"
	"math"
)

var id float64
var ids map[string]float64
var low map[string]float64

func init() {
	id = 0
	ids = make(map[string]float64)
	low = make(map[string]float64)
}

// DFS traverses a graph using dfs technique
func (g graph) findBridges() []string {
	visited := make(map[string]bool)
	bridges := make([]string, 0)
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil
	}
	for i := range g {
		if !visited[i] {
			bridges = g.dfs(i, visited, "-1", bridges)
		}
	}
	return bridges
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfs(at string, visited map[string]bool, parent string, bridges []string) []string {
	visited[at] = true
	id = id + 1
	low[at] = id
	ids[at] = id
	for i := range g[at] {
		to := g[at][i]
		if to == parent {
			continue
		}
		if !visited[to] {
			bridges = g.dfs(to, visited, at, bridges)
			low[at] = math.Min(low[at], low[to])
			if ids[at] < low[to] {
				bridges = append(bridges, at)
				bridges = append(bridges, to)
			}
		} else {
			low[at] = math.Min(low[at], ids[to])
		}
	}
	return bridges
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
