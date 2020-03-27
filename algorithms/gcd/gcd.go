package main

import "fmt"

func main() {
	var n, m int
	fmt.Println("-- GCD --")
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d\n", &n)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d\n", &m)
	var gcd int
	if n < m {
		gcd = findGCD(m, n)
	} else {
		gcd = findGCD(n, m)
	}
	fmt.Println("GCD of given numbers is:", gcd)
}

func findGCD(m, n int) int {
	if n == 0 {
		return m
	}
	return findGCD(n, m%n)
}
