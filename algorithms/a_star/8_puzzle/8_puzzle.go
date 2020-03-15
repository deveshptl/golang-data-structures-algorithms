package main

import "fmt"

type matrix [][]int
type sliceOfMatrix []matrix

// input scrambled puzzle
var inputPuzzle matrix

// the final required output puzzle
var solvedPuzzle matrix

// the steps taken to reach from scrambled puzzle to required output puzzle
var solvedSteps sliceOfMatrix

// heap stores the list of nodes and returns the one having the least cost
var heap []*Node

// Node represents value, path, level and move
// value is the cost of the last matrix in path
// path is the list of matrices which represents the traversal (steps) done from the initial matrix
// level represents the depth at which the current (last matrix from path) state is
// move represents the move done on the last state to transform into the current state
type Node struct {
	value int
	path  sliceOfMatrix
	level int
	move  []int
}

func init() {
	solvedPuzzle = [][]int{[]int{1, 2, 3}, []int{8, 0, 4}, []int{7, 6, 5}}
	solvedSteps = make(sliceOfMatrix, 0)
	heap = make([]*Node, 0)
	inputPuzzle = make(matrix, 3)
	for i := range inputPuzzle {
		inputPuzzle[i] = make([]int, 3)
		for j := range inputPuzzle[i] {
			inputPuzzle[i][j] = 0
		}
	}
}

func main() {
	// scan the input values
	fmt.Println("\n-- Note: For blank cell use '0' --")
	fmt.Println("\nStart entering puzzle values:")
	for i := range inputPuzzle {
		for j := range inputPuzzle[i] {
			value := -1
			fmt.Print("Enter for ", i+1, j+1, ": ")
			fmt.Scanf("%d", &value)
			inputPuzzle[i][j] = value
		}
		fmt.Println("")
	}

	// print the puzzle
	fmt.Println("\nYour scrambled puzzle is: ")
	inputPuzzle.displayMatrix()

	// validate the puzzle
	if inputPuzzle.isValid() {
		fmt.Println("\nA* Algorithm started...")
		inputPuzzle.aStar()
	} else {
		fmt.Println("-- Your input puzzle is invalid. --")
	}
}

// puzzle = [][]int{[]int{2, 8, 3}, []int{1, 6, 4}, []int{7, 0, 5}}
