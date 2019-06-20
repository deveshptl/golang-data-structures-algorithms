package main

import "fmt"

func naiveMatch() []int {
	var occurance []int
	for i := range mainString {
		for j := range pattern {
			if mainString[i+j] != pattern[j] {
				break
			} else if j == (len(pattern) - 1) {
				occurance = append(occurance, i+1)
			}
		}
	}
	return occurance
}

var mainString, pattern string

func main() {
	fmt.Println("\n-- Naive Pattern Matching --")
	fmt.Print("\nEnter the main string: ")
	fmt.Scanf("%s", &mainString)
	fmt.Print("Enter the pattern: ")
	fmt.Scanf("%s", &pattern)
	occurance := naiveMatch()
	if len(occurance) == 0 {
		fmt.Println("-- No pattern found --")
	} else {
		for i := range occurance {
			fmt.Println("Pattern found at index", occurance[i])
		}
	}
}
