package main

import (
	"fmt"
	"math"
	"strings"
)

// LCS recursive solution of lcs problem
func LCS() string {
	for i := range sequence {
		for j := range sequence[i] {
			if i == 0 || j == 0 {
				continue
			} else if string(str01[i]) == string(str02[j]) {
				sequence[i][j] = 1 + sequence[i-1][j-1]
			} else {
				sequence[i][j] = int(math.Max(float64(sequence[i-1][j]), float64(sequence[i][j-1])))
			}
		}
	}

	i := len(str01) - 1
	j := len(str02) - 1
	str := make([]string, sequence[i][j])
	for true {
		if sequence[i][j] == 0 {
			break
		} else if sequence[i-1][j] == sequence[i][j] {
			i = i - 1
		} else if sequence[i][j-1] == sequence[i][j] {
			j = j - 1
		} else {
			str[sequence[i][j]-1] += string(str02[j])
			i = i - 1
			j = j - 1
		}
	}
	return strings.Join(str, "")
}

var str01 string
var str02 string
var sequence [][]int

func main() {
	fmt.Println("\n-- Longest Common Subsequence --")
	fmt.Print("\nEnter the first string: ")
	fmt.Scanf("%s", &str01)
	fmt.Print("Enter the second string: ")
	fmt.Scanf("%s", &str02)
	str01 = "\n" + str01
	str02 = "\n" + str02
	sequence = make([][]int, len(str01))
	for i := range sequence {
		sequence[i] = make([]int, len(str02))
	}
	str := LCS()
	i := len(str01) - 1
	j := len(str02) - 1

	fmt.Println("The Longest Common Subsequence is of length", sequence[i][j], "which is", "'"+str+"'.")
}
