package main

import (
	"fmt"
	"strings"
)

func (g graph) runAStarAlgo(initialVtx, goalVtx string) {
	for vtx := range g[initialVtx] {
		f := g[initialVtx][vtx].value + vtxWeights[g[initialVtx][vtx].name]
		Enqueue(&Node{name: initialVtx + "-" + g[initialVtx][vtx].name, value: f})
	}
	path := g.traverseFurther(Dequeue(), goalVtx)
	fmt.Println("Best Path is:", path.name)
	fmt.Println("The cost is:", path.value)
}

func (g graph) traverseFurther(bestPath *Node, goalVtx string) *Node {
	vertices := strings.Split(bestPath.name, "-")
	currVtx := vertices[len(vertices)-1]
	for vtx := range g[currVtx] {
		visitingVtx := g[currVtx][vtx]
		f := bestPath.value - vtxWeights[currVtx] + g[currVtx][vtx].value + vtxWeights[visitingVtx.name]
		Enqueue(&Node{name: bestPath.name + "-" + visitingVtx.name, value: f})
	}

	nextBestPath := Dequeue()
	vertices = strings.Split(nextBestPath.name, "-")
	if vertices[len(vertices)-1] != goalVtx {
		return g.traverseFurther(nextBestPath, goalVtx)
	}

	return nextBestPath
}
