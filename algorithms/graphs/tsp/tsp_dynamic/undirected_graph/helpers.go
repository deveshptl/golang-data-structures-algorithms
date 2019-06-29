package main

import (
	"math"
)

func notIn(elem, subset uint32) bool {
	return ((1 << elem) & subset) == 0
}

func combinations(set, at, r, n uint32, subsets *[]uint32) *[]uint32 {
	elementsLeftToPick := n - at
	if elementsLeftToPick < r {
		return subsets
	}

	if r == 0 {
		*subsets = append(*subsets, set)
	} else {
		for i := at; i < n; i++ {
			set ^= (1 << i)
			combinations(set, i+1, r-1, n, subsets)
			set ^= (1 << i)
		}
	}
	return subsets
}

func constructMatrix(g graph) {
	numToVtx = make(map[int]string)
	vtxToNum = make(map[string]int)
	num := 0
	for i := range g {
		numToVtx[num] = i
		vtxToNum[i] = num
		num++
	}

	distanceMatrix = make([][]float64, len(numToVtx))

	for i := range g {
		distanceMatrix[vtxToNum[i]] = make([]float64, len(vtxToNum))
		for j := range distanceMatrix[vtxToNum[i]] {
			distanceMatrix[vtxToNum[i]][j] = math.Inf(1)
		}
		for j := range g[i] {
			distanceMatrix[vtxToNum[i]][vtxToNum[g[i][j].name]] = float64(g[i][j].value)
		}
	}
}
