package main

import "fmt"

// Knp implements Knuth Morris Pratt Algorithms
func Knp() {
	computePi()
	i := 0
	j := 0
	m := len(pattern)
	n := len(text)
	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}
		if j == m {
			fmt.Println("Pattern found at", i-j)
			j = pi[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = pi[j-1]
			} else {
				i++
			}
		}
	}
}

var text string
var pattern string
var pi []int

func main() {
	fmt.Println("\n-- Knuth Morris Pratt Algorithm --")
	fmt.Print("\nEnter the initial text: ")
	fmt.Scanf("%s\n", &text)
	fmt.Print("Enter the pattern: ")
	fmt.Scanf("%s\n", &pattern)
	Knp()
}

func computePi() {
	pi = make([]int, len(pattern))
	pi[0] = 0
	k := 0
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[k] {
			k++
			pi[i] = k
		} else {
			if k != 0 {
				k = pi[k-1]
			} else {
				pi[i] = k
			}
		}
	}
}
