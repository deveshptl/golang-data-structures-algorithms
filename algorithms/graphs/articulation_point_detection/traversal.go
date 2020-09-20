package main

import (
	"fmt"
	"math"
)

var id float64
var outEdgeCount float64
var ids map[string]float64
var low map[string]float64
var isArt map[string]bool

func init() {
	id = 0
	outEdgeCount = 0
	ids = make(map[string]float64)
	low = make(map[string]float64)
	isArt = make(map[string]bool)
}

// DFS traverses a graph using dfs technique
func (g graph) findArticulationPoints() map[string]bool {
	visited := make(map[string]bool)
	if len(g) == 0 {
		fmt.Println("\n-- Graph is empty. --")
		return nil
	}
	for i := range g {
		if !visited[i] {
			outEdgeCount = 0
			g.dfs(i, visited, i, "-1")
			isArt[i] = outEdgeCount > 1
		}
	}
	return isArt
}

// dfsHelper recursively calls itself to solve dfs traversal
func (g graph) dfs(root string, visited map[string]bool, at string, parent string) {
	visited[at] = true
	id++
	low[at] = id
	ids[at] = id
	for i := range g[at] {
		to := g[at][i]
		if to == parent {
			continue
		}
		if !visited[to] {
			g.dfs(root, visited, to, at)
			low[at] = math.Min(low[at], low[to])
			if ids[at] < low[to] {
				isArt[at] = true
			}
			if ids[at] == low[to] {
				isArt[at] = true
			}
		} else {
			low[at] = math.Min(low[at], ids[to])
		}
	}
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
