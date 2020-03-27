package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-- Travelling Salesman Problem using Naïve Approach --")
	g := make(graph)

	i := 0
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. REMOVE VERTEX")
		fmt.Println("4. REMOVE AN EDGE")
		fmt.Println("5. DISPLAY GRAPH")
		fmt.Println("6. PERFORM TSP")
		fmt.Println("7. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			g.addVertex()
		case 2:
			g.addEdge()
		case 3:
			g.removeVertex()
		case 4:
			g.removeEdge()
		case 5:
			g.simpleDisplay()
		case 6:
			var initVtx string
			fmt.Print("Enter the intial vertex: ")
			fmt.Scanf("%s\n", &initVtx)
			allRoutes, minCostRoutes := TSP(g, initVtx)

			fmt.Println("\n-- Cost of all possible Routes --")
			for i := range allRoutes {
				fmt.Println(i, "-", allRoutes[i])
			}

			fmt.Println("\n-- Routes with minimal Cost --")
			for i := range minCostRoutes {
				fmt.Println(i, "-", minCostRoutes[i])
			}
		case 7:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

// TSP performs travelling saleman algorithm using naive approach on a given undirected weighted graph.
func TSP(g graph, initVtx string) (map[string]int, map[string]int) {
	// costs to all possible routes will be stored
	result := make(map[string]int, 0)
	// maps numbers to vertices to generate all possible permutations later which will make it easy to map from numbers to vertices
	numToVtx := make(map[string]string, 0)
	// vtx number which represents the number count of the node being visited
	vtx := 0
	// string on which all possible permutations will performed
	strForPermutation := ""
	// holds the vertex's number count from where the tsp algorithm begins
	temp := 0

	for i := range g {
		strNum := strconv.Itoa(vtx)
		numToVtx[strNum] = i
		if i != initVtx {
			strForPermutation += strNum
		} else {
			temp = vtx
		}
		vtx++
	}

	// generate permutations
	arr := LOP(strings.Split(strForPermutation, ""))

	// append temp string at begining and ending of each permutation
	for i := 0; i < len(arr); i++ {
		arr[i] = strconv.Itoa(temp) + arr[i] + strconv.Itoa(temp)
	}

	// find cost of each routes
	for i := 0; i < len(arr); i++ {

		fromVtx := ""
		toVtx := ""
		currCost := 0
		flag := 0

		for j := 0; j < len(arr[i])-1; j++ {
			fromVtx = numToVtx[string(arr[i][j])]
			toVtx = numToVtx[string(arr[i][j+1])]

			for k := 0; k < len(g[fromVtx]); k++ {
				if g[fromVtx][k].name == toVtx {
					flag = 1
					currCost += g[fromVtx][k].value
					break
				}
			}

			if flag == 0 {
				break
			}

		}

		if flag != 0 {
			flag = 0
			result[arr[i]] = currCost
		}

	}

	flag := 0
	cost := 0

	// computing the least possible cost and replacing route names having numbers with route names having original vertex names
	for i := range result {

		// set initial cost
		if flag == 0 {
			flag = 1
			cost = result[i]
		}

		// holds the route according to graph's vertex name
		route := ""

		for j := 0; j < len(i); j++ {
			route += numToVtx[string(i[j])]
		}

		if cost > result[i] {
			cost = result[i]
		}

		val := result[i]
		delete(result, i)
		result[route] = val
	}

	routesWithMinCost := make(map[string]int, 0)
	// add routes that costs the least
	for i := range result {
		if result[i] == cost {
			routesWithMinCost[i] = cost
		}
	}

	return result, routesWithMinCost
}

// g.addVertexToGraph("1")
// g.addVertexToGraph("2")
// g.addVertexToGraph("3")
// g.addVertexToGraph("4")
// g.addEdgeToGraph("1", "2", 10)
// g.addEdgeToGraph("1", "3", 15)
// g.addEdgeToGraph("1", "4", 20)
// g.addEdgeToGraph("2", "1", 5)
// g.addEdgeToGraph("2", "3", 9)
// g.addEdgeToGraph("2", "4", 10)
// g.addEdgeToGraph("3", "1", 6)
// g.addEdgeToGraph("3", "2", 13)
// g.addEdgeToGraph("3", "4", 12)
// g.addEdgeToGraph("4", "1", 8)
// g.addEdgeToGraph("4", "2", 8)
// g.addEdgeToGraph("4", "3", 9)
// -- all --
// 14231 - 43
// 14321 - 47
// 12341 - 39
// 12431 - 35
// 13241 - 46
// 13421 - 40
// -- minimal --
// 12431 - 35
