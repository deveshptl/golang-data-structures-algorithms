package main

import "fmt"

var arr []int

func main() {
	var size int
	var element int
	fmt.Println("\n-- Bubble Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)
		arr = append(arr, element)
	}

	// sorting
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted Array is:")
	fmt.Println(arr)
}
