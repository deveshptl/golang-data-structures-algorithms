package main

import (
	"fmt"
	"math"
)

func knapsack() {
	op := make([][]int, size+1)
	for i := range op {
		op[i] = make([]int, capacity+1)
		for j := range op[i] {
			op[i][0] = 0
			op[0][j] = 0
		}
	}

	for i := range op {
		for j := range op[i] {
			if i == 0 || j == 0 {
				continue
			} else if j < weights[i] {
				op[i][j] = op[i-1][j]
			} else if j >= weights[i] {
				op[i][j] = int(math.Max(float64(op[i-1][j]), float64(op[i-1][j-weights[i]]+values[i])))
			}
		}
	}

	// finding the right weights and values from the matrix generated
	i := size
	j := capacity
	var finalWeights []int
	var finalValues []int

	for i != 0 {
		if op[i][j] == op[i-1][j] {
			i--
		} else {
			finalWeights = append(finalWeights, weights[i])
			finalValues = append(finalValues, values[i])
			j -= weights[i]
		}
	}

	// printing the answer
	fmt.Println("Matrix formed:")
	for i := range op {
		fmt.Println(op[i])
	}
	fmt.Println("\nWeights taken:")
	fmt.Println(finalWeights)
	fmt.Println("Values taken:")
	fmt.Println(finalValues)
}

var weights, values []int
var capacity, size int

func init() {
	weights = append(weights, 0)
	values = append(values, 0)
}

func main() {
	var temp int
	fmt.Println("\n-- 0-1 Knapsack Problem using Dynamic Programming --")
	fmt.Print("\nEnter the number of objects: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering weight and value of each object:")
	for i := 0; i < size; i++ {
		fmt.Print("\nWeight: ")
		fmt.Scanf("%d\n", &temp)
		weights = append(weights, temp)
		fmt.Print("Value: ")
		fmt.Scanf("%d\n", &temp)
		values = append(values, temp)
	}
	fmt.Print("\nEnter the capacity of knapsack: ")
	fmt.Scanf("%d\n", &capacity)
	knapsack()
}

// FORMULA
// stp1: make v[i][0] and v[0][j] = 0
// stp2: if j < weights[i] then v[i][j] = v[i-1][j]
// stp3: if j >= weights[i] then v[i][j] = max(v[i-1][j], v[i-1][j-weights[i]]+v[i])
