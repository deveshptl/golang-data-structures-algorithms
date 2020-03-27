package main

import (
	"fmt"
	"math"
)

func countingSort(size, max int) []int {
	count := make([]int, max)
	sorted := make([]int, size)
	for i := 0; i < size; i++ {
		count[arr[i]] = count[arr[i]] + 1
	}
	for i := 1; i < max; i++ {
		count[i] = count[i] + count[i-1]
	}
	for i := size - 1; i >= 0; i-- {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]] = count[arr[i]] - 1
	}
	return sorted
}

var arr []int

func main() {
	var size, element, min, max int
	fmt.Println("\n-- Counting Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)

		if i == 0 {
			min = element
			max = element
		}

		min = int(math.Min(float64(min), float64(element)))
		max = int(math.Max(float64(max), float64(element)))
		arr = append(arr, element)
	}
	sorted := countingSort(size, max+1)
	fmt.Println("Sorted array is:")
	fmt.Println(sorted)
}
