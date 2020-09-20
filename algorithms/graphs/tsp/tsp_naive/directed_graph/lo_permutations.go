package main

import (
	"strings"
)

// LOP performs Lexicographic Order Permutations on a given string.
// Returns all possibilities as a slice of string.
func LOP(str []string) []string {
	result := make([]string, 0)
	result = append(result, strings.Join(str, ""))
	for {
		var i int
		var j int

		for i = len(str) - 1; i >= 0; i-- {
			if i != 0 && str[i-1] < str[i] {
				break
			}
		}

		if i == -1 {
			break
		}

		for j = len(str) - 1; j >= 0; j-- {
			if str[i-1] < str[j] {
				str[i-1], str[j] = str[j], str[i-1]
				str = Reverse(str, i)
				result = append(result, strings.Join(str, ""))
				break
			}
		}

	}
	return result
}

// Reverse reverses a slice of string from it's length to given index
func Reverse(str []string, index int) []string {
	newStr := strings.Join(str[0:index], "")
	for i := len(str) - 1; i >= index; i-- {
		newStr += str[i]
	}
	return strings.Split(newStr, "")
}
