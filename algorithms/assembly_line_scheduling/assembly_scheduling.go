package main

import (
	"fmt"
	"math"
)

var assembly [][]int
var transport [][]int
var result [][]int
var stations int

func init() {
	assembly = make([][]int, 2)
	transport = make([][]int, 2)
	result = make([][]int, 2)
}

func main() {
	fmt.Println("\n-- Assembly Line Scheduling --")
	fmt.Print("Enter the number of stations in assembly line: ")
	fmt.Scanf("%d", &stations)

	fmt.Println("Start entering the costs:")
	var temp int
	for i := 0; i < 2; i++ {
		assembly[i] = make([]int, stations+2)
		for j := 0; j < stations+2; j++ {
			fmt.Print(i, "-", j, ": ")
			fmt.Scanf("%d", &temp)
			assembly[i][j] = temp
		}
	}

	fmt.Println("Start entering the transport costs:")
	for i := 0; i < 2; i++ {
		transport[i] = make([]int, stations+2)
		for j := 0; j < stations-1; j++ {
			fmt.Print(i+1, "-", j+1, ": ")
			fmt.Scanf("%d", &temp)
			transport[i][j+1] = temp
		}
	}

	assemblySchedule()

	fmt.Println("\nThe matrix formed during calculations:")
	fmt.Println(result[0])
	fmt.Println(result[1])

	if result[0][stations] < result[1][stations] {
		fmt.Println("\nCheapest way to assemble the product is:", result[0][stations])
	} else {
		fmt.Println("\nCheapest way to assemble the product is:", result[1][stations])
	}
}

func assemblySchedule() {
	i := 0
	result[i] = make([]int, stations+1)
	result[i+1] = make([]int, stations+1)
	for j := 0; j < stations; j++ {
		if j == 0 {
			result[i][j] = assembly[i][j] + assembly[i][j+1]
			result[i+1][j] = assembly[i+1][j] + assembly[i+1][j+1]
		} else {
			result[i][j] = int(math.Min(float64(result[i][j-1]+assembly[i][j+1]), float64(result[i+1][j-1]+transport[i+1][j]+assembly[i][j+1])))
			result[i+1][j] = int(math.Min(float64(result[i+1][j-1]+assembly[i+1][j+1]), float64(result[i][j-1]+transport[i][j]+assembly[i+1][j+1])))
		}
		check := j + 1
		if check == stations {
			result[i][j+1] = result[i][j] + assembly[i][j+2]
			result[i+1][j+1] = result[i+1][j] + assembly[i+1][j+2]
		}
	}
}

// assembly[0] = []int{2, 7, 9, 3, 4, 8, 4, 3}
// assembly[1] = []int{4, 8, 5, 6, 4, 5, 7, 2}

// transport[0] = []int{0, 2, 3, 1, 3, 4, 0, 0}
// transport[1] = []int{0, 2, 1, 2, 2, 1, 0, 0}
