package main

import "fmt"

var arr []int

func main() {
	var size int
	var element int
	fmt.Println("\n-- Insertion Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d", &size)
	fmt.Println("Start entering the elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &element)
		arr = append(arr, element)
	}

	for i := 0; i < size-1; i++ {
		j := i
		key := arr[i+1]
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	fmt.Println("Sorted Array is:")
	fmt.Println(arr)
}
