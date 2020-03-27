package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var distanceMatrix [][]float64
var numToVtx map[int]string
var vtxToNum map[string]int
var n uint32

func main() {
	fmt.Println("-- Travelling Salesman Problem using Dynamic Programming --")
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
			fmt.Print("Enter the starting vertex: ")
			var initVtx string
			fmt.Scanf("%s\n", &initVtx)
			tours, cost := TSP(g, initVtx)
			fmt.Println("\nTours:")
			for i := 0; i < len(tours); i++ {
				fmt.Println(strings.Join(tours[i], ""))
			}
			fmt.Println("\nCost: ", cost)
		case 7:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

// TSP performs travelling saleman algorithm using naive approach on a given undirected weighted graph.
func TSP(g graph, initVtx string) ([][]string, float64) {
	constructMatrix(g)

	// some other tsp related initializations
	n = uint32(len(distanceMatrix))
	start := uint32(vtxToNum[initVtx])
	minTourCost := math.Inf(1)
	tour := make([]string, 0)
	endState := uint32((1 << n) - 1)
	memo := make([][]float64, n)

	// assign all elements in memo table to infinity
	for i := range memo {
		memo[i] = make([]float64, 1<<n)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = math.Inf(1)
		}
	}

	// initialize memo table by adding all outgoing edges starting from start vertex
	for end := uint32(0); end < n; end++ {
		if end == start {
			continue
		}
		memo[end][(1<<start)|(1<<end)] = distanceMatrix[start][end]
	}

	for r := uint32(3); r <= n; r++ {
		subsets := new([]uint32)
		var subset uint32
		subsets = combinations(0, 0, r, n, subsets)
		for i := 0; i < len(*subsets); i++ {
			subset = (*subsets)[i]
			if notIn(start, subset) {
				continue
			}
			for next := uint32(0); next < n; next++ {
				if next == start || notIn(next, subset) {
					continue
				}
				subsetWithoutNext := subset ^ (1 << next)
				minDist := math.Inf(1)
				for end := uint32(0); end < n; end++ {
					if end == start || end == next || notIn(end, subset) {
						continue
					}
					newDist := memo[end][subsetWithoutNext] + distanceMatrix[end][next]
					if newDist < minDist {
						minDist = newDist
					}
				}
				memo[next][subset] = minDist
			}
		}
	}

	// Finding minimum cost
	for i := uint32(0); i < n; i++ {
		if i == start {
			continue
		}
		tourCost := memo[i][endState] + distanceMatrix[i][start]
		if tourCost < minTourCost {
			minTourCost = tourCost
		}
	}

	lastIndex := start
	state := endState
	flag := 1

	// construct the final path from memo table
	for flag != 0 {

		flag = 0
		tour = append(tour, strconv.Itoa(int(start)))

		for i := uint32(1); i < n; i++ {
			index := -1
			for j := uint32(0); j < n; j++ {
				if j == start || notIn(j, state) {
					continue
				}
				if index == -1 {
					index = int(j)
				}
				prevDist := memo[index][state] + distanceMatrix[index][lastIndex]
				newDist := memo[j][state] + distanceMatrix[j][lastIndex]
				if newDist < prevDist {
					index = int(j)
				} else if newDist == prevDist && index != int(j) && newDist != math.Inf(1) {
					flag = 1
				}
			}

			memo[index][state] = math.Inf(1)

			tour = append(tour, ", ", strconv.Itoa(index))
			state = state ^ (1 << uint32(index))
			lastIndex = uint32(index)
		}
		tour = append(tour, ", ", strconv.Itoa(int(start)), ", ", fmt.Sprintf("%f\n", math.Inf(1)))
		state = endState
		lastIndex = start
	}

	// Some Reformating of strings
	//									--------------	---------------
	// Format of type: [2, 0, 1, 3, 2,  2, 3, 1, 0, 2,  ]
	tours := strings.Split(strings.Join(tour, ""), fmt.Sprintf("%f\n", math.Inf(1)))
	// eclude the trailing empty []string
	tours = tours[0 : len(tours)-1]

	reverseTour := make([][]string, 0)

	for i := 0; i < len(tours); i++ {
		reverseTour = append(reverseTour, make([]string, 0))
		arr := strings.Split(tours[i], ", ")
		// exclude the trailing empty string
		arr = arr[0 : len(arr)-1]
		for j := len(arr) - 1; j >= 0; j-- {
			num, _ := strconv.Atoi(arr[j])
			reverseTour[i] = append(reverseTour[i], numToVtx[num], ", ")
		}
		reverseTour[i] = reverseTour[i][0 : len(reverseTour[i])-1]
	}

	return reverseTour, minTourCost
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
// -- minimal --
// 12431 - 35

// Reference: https://github.com/williamfiset
// Couldn't have done without it. Huge thanks to William. :-)
