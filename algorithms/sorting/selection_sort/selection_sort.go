package main

import "fmt"

var arr []int

func main() {
	var length int
	var element int
	fmt.Println("\n-- Selection Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d", &length)
	fmt.Println("Start entering the elements: ")
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &element)
		arr = append(arr, element)
	}

	// Sorting
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("Sorted Array is:")
	fmt.Println(arr)
}
