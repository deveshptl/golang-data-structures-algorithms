package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-- Sieve of Eratosthenes: Generating consecutive primes numbers not exceeding n --")
	var n int
	fmt.Print("Enter the number greater than 1: ")
	fmt.Scanf("%d", &n)
	arr := Sieve(n)
	fmt.Println("Consecutive prime number not exceeding", n, "are: ")
	fmt.Println(arr)
}

// Sieve implements Sieve of Eratosthenes algorithm
func Sieve(n int) []int {
	arr := make([]int, 0)
	limit := int(math.Floor(math.Sqrt(float64(n))))

	for i := 0; i <= n; i++ {
		arr = append(arr, i)
	}

	arr[1] = 0

	for i := 2; i <= limit; i++ {
		if arr[i] != 0 {
			j := i * i
			for j <= n {
				arr[j] = 0
				j = j + i
			}
		}
	}

	newArr := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			newArr = append(newArr, arr[i])
		}
	}

	return newArr
}
