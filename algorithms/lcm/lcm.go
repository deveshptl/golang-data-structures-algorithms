package main

import "fmt"

func main() {
	fmt.Println("\n-- Least Common Multiple --")
	var m, n, lcm int
	fmt.Print("\nEnter the first number: ")
	fmt.Scanf("%d\n", &m)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d\n", &n)
	if n < m {
		lcm = findLCM(m, n)
	} else {
		lcm = findLCM(n, m)
	}
	fmt.Println("Least Common Multiple of", m, "and", n, "is:", lcm)
}

func findLCM(m, n int) int {
	val := findGCD(m, n)
	return (m * n) / val
}

func findGCD(m, n int) int {
	if n == 0 {
		return m
	}
	return findGCD(n, m%n)
}
