package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("\n-- Boyer Moore Algorithm --")

	var text, pattern string

	fmt.Print("\nEnter the text string: ")
	fmt.Scanf("%s", &text)
	fmt.Print("Enter the pattern string: ")
	fmt.Scanf("%s", &pattern)

	// shift table
	st := ShiftTable(strings.Split(pattern, ""))

	// offset table
	ot := OffSetTable(strings.Split(pattern, ""))
	fmt.Println(st, ot)

	index := Search(strings.Split(text, ""), strings.Split(pattern, ""), st, ot)

	if len(index) == 0 {
		fmt.Println("Pattern not found in text!")
	} else {
		fmt.Println("Pattern found at index/ices:")
		fmt.Println(index)
	}
}

// ShiftTable creates shift table for given pattern
func ShiftTable(pt []string) map[string]int {
	st := make(map[string]int)
	n := len(pt)
	for i := 0; i < n; i++ {
		if i == n-1 {
			if st[pt[i]] == 0 {
				st[pt[i]] = n
				continue
			}
			continue
		}
		st[pt[i]] = n - i - 1
	}

	return st
}

// OffSetTable table creates offset table for given pattern string
func OffSetTable(pt []string) []int {

	n := len(pt)
	ot := make([]int, n)
	index := len(pt)

	for i := n; i > 0; i-- {
		if IsPrefix(pt, i) {
			index = i
		}
		ot[n-i] = index - i + n
	}

	for i := 0; i < n-1; i++ {
		l := SuffixLength(pt, i)
		ot[l] = n - 1 - i + l
	}

	return ot
}

// IsPrefix checks whether pt[k:] is prefix of string pt
func IsPrefix(pt []string, k int) bool {
	i := 0
	for j := k; j < len(pt); j++ {
		if pt[i] != pt[j] {
			return false
		}
		i = i + 1
	}
	return true
}

// SuffixLength computes maximum length of substring that ends at k and is suffix of string pt
func SuffixLength(pt []string, k int) int {
	l := 0
	j := len(pt) - 1

	for i := k; i >= 0 && pt[i] == pt[j]; i-- {
		l++
		j--
	}

	return l
}

// Search searches for pattern string in given text string and returns the indices values if pattern is found
// else returns empty array
func Search(txt, pt []string, st map[string]int, ot []int) []int {

	pos := make([]int, 0)

	var j int
	for i := len(pt) - 1; i >= 0 && i < len(txt); {
		skip := 0

		for j = len(pt) - 1; j >= 0 && pt[j] == txt[i]; j-- {
			i--
			if j == 0 {
				pos = append(pos, i+2)
				i++
				skip = 0
				break
			}
		}

		if skip == 1 {
			continue
		}

		otVal := float64(ot[len(pt)-1-j])
		stVal := float64(st[txt[i]])

		if stVal == 0 {
			stVal = float64(len(pt) - 1)
		}

		i += int(math.Max(otVal, stVal))

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
