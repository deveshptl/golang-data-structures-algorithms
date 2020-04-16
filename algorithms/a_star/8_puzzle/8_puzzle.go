package main

import "fmt"

type matrix [][]int
type sliceOfMatrix []matrix

// input scrambled puzzle
var inputPuzzle matrix

// the final required output puzzle
var outputPuzzle matrix

// the steps taken to reach from scrambled puzzle to required output puzzle
var solvedSteps sliceOfMatrix

// heap stores the list of nodes and returns the one having the least cost
var heap []*Node

// Node represents value, path, level and move
// value is the cost of the last matrix in path
// path is the list of matrices which represents the traversal (steps) done from the initial matrix
// level represents the depth at which the current (last matrix from path) state is
// parentBlankCell represents the row and column in which the blank cell was in the parent state
type Node struct {
	value           int
	path            sliceOfMatrix
	level           int
	parentBlankCell []int
}

func init() {
	outputPuzzle = make(matrix, 3)
	solvedSteps = make(sliceOfMatrix, 0)
	heap = make([]*Node, 0)
	inputPuzzle = make(matrix, 3)

	// initialize outputPuzzle
	for i := range outputPuzzle {
		outputPuzzle[i] = make([]int, 3)
		for j := range outputPuzzle[i] {
			outputPuzzle[i][j] = 0
		}
	}

	// initialize inputPuzzle
	for i := range inputPuzzle {
		inputPuzzle[i] = make([]int, 3)
		for j := range inputPuzzle[i] {
			inputPuzzle[i][j] = 0
		}
	}
}

func main() {
	// scan the input values
	fmt.Println("\n-- 8 Puzzle problem using A* Algorithm --")
	fmt.Println("\n-- Note: For blank cell use '0' --")

	fmt.Println("\nStart entering desired 3x3 output puzzle:")
	for i := range outputPuzzle {
		for j := range outputPuzzle[i] {
			value := -1
			fmt.Print("Enter for ", i+1, j+1, ": ")
			fmt.Scanf("%d\n", &value)
			outputPuzzle[i][j] = value
		}
		fmt.Println("")
	}

	if !outputPuzzle.isValid() {
		fmt.Println("\n-- Invalid output puzzle. --")
		return
	}

	fmt.Println("\nStart entering 3x3 scrambled input puzzle:")
	for i := range inputPuzzle {
		for j := range inputPuzzle[i] {
			value := -1
			fmt.Print("Enter for ", i+1, j+1, ": ")
			fmt.Scanf("%d\n", &value)
			inputPuzzle[i][j] = value
		}
		fmt.Println("")
	}

	// print the desired output puzzle
	fmt.Println("\nYour desired output puzzle is:")
	outputPuzzle.displayMatrix()

	// print the input scrambled puzzle
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

// puzzle = input - output
// puzzle = [][]int{[]int{2, 8, 3}, []int{1, 6, 4}, []int{7, 0, 5}} - [][]int{[]int{1, 2, 3}, []int{8, 0, 4}, []int{7, 6, 5}}
// puzzle = matrix{[]int{1, 8, 2}, []int{0, 4, 3}, []int{7, 6, 5}} - [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 0}}
