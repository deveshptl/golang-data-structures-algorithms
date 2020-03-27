package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n-- Boyer Moore Horspool Algorithm --")

	var text, pattern string

	fmt.Print("\nEnter the text string: ")
	fmt.Scanf("%s\n", &text)
	fmt.Print("Enter the pattern string: ")
	fmt.Scanf("%s\n", &pattern)

	st := ShiftTable(strings.Split(pattern, ""))
	fmt.Println(st)
	index := Search(strings.Split(text, ""), strings.Split(pattern, ""), st)

	if len(index) == 0 {
		fmt.Println("\n-- Pattern not found in text! --")
	} else {
		fmt.Println("\nPattern found at index/ices:")
		fmt.Println(index)
	}

}

// ShiftTable creates shift table for given pattern
func ShiftTable(pt []string) map[string]int {
	st := make(map[string]int)
	n := len(pt)
	for i := 0; i < n-1; i++ {
		if i == n-1 {
			if st[pt[i]] == 0 {
				st[pt[i]] = n
				continue
			}
			continue
		}
		st[pt[i]] = n - 1 - i
	}

	return st
}

// Search searches for pattern string in given text string and returns the indices values if pattern is found
// else returns empty array
func Search(txt, pt []string, st map[string]int) []int {

	pos := make([]int, 0)

	var j int
	for i := len(pt) - 1; i >= 0 && i < len(txt); {
		skip := 0

		for j = len(pt) - 1; j >= 0 && pt[j] == txt[i]; j-- {
			i--
			if j == 0 {
				pos = append(pos, i+2)
				i += len(pt) + 1
				skip = 1
				break
			}
		}

		if skip == 1 {
			continue
		}

		stVal := st[txt[i]]

		if stVal == 0 {
			stVal = len(pt) - 1
		}

		i += stVal
	}

	return pos
}

// Sample test strings
//
// GCAATGCCTATGTGACC
// TATGTG
// 9
//
// AABAACAADAABAABA
// AABA
// [1 10 13]
//
// TRUSTHARDTOOTHBRUSHES
// TOOTH
// 10
//
// GCATATGTGATGCCTATGTGACCTATGTG
// TATGTG
// [4 15 24]
//
// JIM SAW ME IN A BARBER SHOP
// BARBER
// [7]
//
// JIMBARBER SAW ME IN A BARBER SHOP
// BARBER
// [4 23]
