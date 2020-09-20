package main

import (
	"fmt"
	"math"
)

// MakeChange will generate the necessary matrix which will then be evaluated to know
// the number of coins and which coins will be required to have a given change
func MakeChange() {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 {
				continue
			} else if i == 1 {
				matrix[i][j] = 1 + matrix[i][j-int(coins[i-1])]
			} else if j < int(coins[i-1]) {
				matrix[i][j] = matrix[i-1][j]
			} else {
				matrix[i][j] = math.Min(matrix[i-1][j], 1+matrix[i][j-int(coins[i-1])])
			}
		}
	}
	fmt.Println("\nMatrix formed:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	i := len(matrix) - 1
	j := len(matrix[0]) - 1
	result := make([]float64, 0)
	count := matrix[i][j]
	for count != 0 {
		if matrix[i-1][j] == count {
			i--
		} else {
			j -= int(coins[i-1])
			result = append(result, coins[i-1])
			count--
		}
	}
	fmt.Println("You will require", len(result), "coins of values", result, "to make a change for", change, ".")
}

var change float64
var numCoins float64
var coins []float64
var matrix [][]float64

func main() {
	fmt.Println("\n-- Making Change Problem using Dynamic Programming. --")
	fmt.Print("Enter the value for the amount of change you want: ")
	fmt.Scanf("%f\n", &change)
	fmt.Print("Enter the number of coins you have: ")
	fmt.Scanf("%f\n", &numCoins)
	var coin float64
	fmt.Println("Start enter the values of these coins in increasing order:")
	for i := 0; i < int(numCoins); i++ {
		fmt.Scanf("%f\n", &coin)
		coins = append(coins, coin)
	}
	matrix = make([][]float64, int(numCoins)+1)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]float64, int(change)+1)
		for j := range matrix[i] {
			matrix[i][j] = float64(0)
		}
	}
	MakeChange()
}
