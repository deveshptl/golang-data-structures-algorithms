package main

import (
	"fmt"
	"strconv"
)

// isValid function checks whether the given scrambled puzzle contains valid values or not
// returns false if invalid
// returns true if otherwise
func (puzzle matrix) isValid() bool {
	values := make(map[string]int)
	for i := 1; i <= 8; i++ {
		values[strconv.Itoa(i)] = 0
	}

	values[strconv.Itoa(0)] = 0

	for i := range puzzle {
		for j := range puzzle[i] {
			currValue := strconv.Itoa(puzzle[i][j])
			if _, ok := values[currValue]; ok {
				values[strconv.Itoa(puzzle[i][j])]++
			} else {
				return false
			}
		}
	}

	for i := range values {
		if values[i] != 1 {
			return false
		}
	}

	return true
}

// aStar function runs a* algorithm on the given scrambled puzzle
func (puzzle matrix) aStar() {
	states, parentBlankCell := puzzle.nextStates([]int{-1, -1})
	for i := range states {
		// f = h + g
		cost := states[i].calculateCost() + 1
		path := make(sliceOfMatrix, 0)
		path = append(path, copyMatrix(puzzle))
		path = append(path, copyMatrix(states[i]))
		Enqueue(&Node{value: cost, path: path, level: 1, parentBlankCell: parentBlankCell})
	}

	node := traverseFurther(Dequeue())
	if node.value == -1 {
		fmt.Println("To solve this particular input, it would take more memory and time. Hence, terminating the algorithm.")
		fmt.Println("Level reached:", node.level)
		return
	}
	node.path.displayPath()
}

// traverseFurther recursively generates new puzzles from current state and stores them into the heap
func traverseFurther(puzzle *Node) *Node {
	states, parentBlankCell := puzzle.path[len(puzzle.path)-1].nextStates(puzzle.parentBlankCell)
	for i := range states {
		// f = h + g
		cost := states[i].calculateCost() + puzzle.level + 1

		path := copySliceOfMatrix(puzzle.path)
		temp := copyMatrix(states[i])

		path = append(path, temp)

		node := &Node{value: cost, path: path, level: puzzle.level + 1, parentBlankCell: parentBlankCell}
		Enqueue(node)
	}

	node := Dequeue()
	cost := node.path[len(node.path)-1].calculateCost()
	if puzzle.level >= 15 {
		return &Node{value: -1, level: puzzle.level}
	} else if cost != 0 {
		return traverseFurther(node)
	}
	return node
}

// nextStates function generates new states from the possible moves
func (puzzle matrix) nextStates(move []int) (sliceOfMatrix, []int) {
	i, j := puzzle.findBlankCell()
	moves := findNextMoves(i, j)

	states := make(sliceOfMatrix, 0)

	for k := range moves {
		if moves[k][0] == move[0] && moves[k][1] == move[1] {
			moves = append(moves[:k], moves[k+1:]...)
			break
		}
	}

	for k := range moves {
		newState := copyMatrix(puzzle)
		newState[i][j] = newState[moves[k][0]][moves[k][1]]
		newState[moves[k][0]][moves[k][1]] = 0
		states = append(states, newState)
	}

	return states, []int{i, j}
}

// findNextMoves generates possible moves of blank (0) cell
func findNextMoves(i, j int) matrix {
	moves := make([][]int, 0)
	finalMoves := make([][]int, 0)

	// decrement i
	moves = append(moves, []int{i - 1, j})

	// increment j
	moves = append(moves, []int{i, j + 1})

	// increment i
	moves = append(moves, []int{i + 1, j})

	// decrement j
	moves = append(moves, []int{i, j - 1})

	// validate new moves
	for _, move := range moves {
		if move[0] >= 0 && move[0] <= 2 && move[1] >= 0 && move[1] <= 2 {
			finalMoves = append(finalMoves, move)
		}
	}
	return finalMoves
}

// calculateCost returns the number of misplaced cell
func (puzzle matrix) calculateCost() int {
	// calculating h
	h := 0
	for i := range puzzle {
		for j := range puzzle[i] {
			if puzzle[i][j] != solvedPuzzle[i][j] {
				h++
			}
		}
	}
	return h
}

// findBlankCell return the row and column in which the blank (0) cell resides
func (puzzle matrix) findBlankCell() (int, int) {
	for i := range puzzle {
		for j := range puzzle[i] {
			if puzzle[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// copyMatrix function copies the values of given matrix and returns a new matrix
func copyMatrix(mtx matrix) matrix {
	newMtx := make([][]int, 3)
	for i := range mtx {
		newMtx[i] = make([]int, 3)
		for j := range mtx[i] {
			newMtx[i][j] = mtx[i][j]
		}
	}
	return newMtx
}

// copySliceOfMatrix function copies the values of given sliceOfMatrix and returns a new sliceOfMatrix
func copySliceOfMatrix(mtx sliceOfMatrix) sliceOfMatrix {
	newSliceOfMtx := make(sliceOfMatrix, len(mtx))
	for i := range mtx {
		newSliceOfMtx[i] = make(matrix, 3)
		for j := range mtx[i] {
			newSliceOfMtx[i][j] = make([]int, 3)
			for k := range mtx[i][j] {
				newSliceOfMtx[i][j][k] = mtx[i][j][k]
			}
		}
	}
	return newSliceOfMtx
}

// displayPath function will display the moves taken in a given path
func (path sliceOfMatrix) displayPath() {
	fmt.Println("")
	for i := range path {
		fmt.Println("-- Level", i, "--")
		fmt.Println("Cost:", path[i].calculateCost())
		path[i].displayMatrix()
		fmt.Println("")
	}
}

// displayMatrix function will display a given matrix
func (puzzle matrix) displayMatrix() {
	for i := range puzzle {
		for j := range puzzle[i] {
			fmt.Print(puzzle[i][j], " ")
		}
		fmt.Println("")
	}
}
