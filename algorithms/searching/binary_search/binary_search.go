package main

import "fmt"

func binarySearch(size, key int) int {
	var low, middle, high, index int
	low = 0
	high = size - 1

	for low <= high {
		middle = (low + high) / 2
		if arr[middle] < key {
			low = middle + 1
		} else if arr[middle] > key {
			high = middle - 1
		} else if arr[middle] == key {
			index = middle + 1
			return index
		}
	}
	return -1
}

var arr []int

func main() {
	var size, element, key int
	fmt.Println("\n-- Binary Search --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements in sorted manner:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)
		arr = append(arr, element)
	}
	fmt.Print("Enter the search key: ")
	fmt.Scanf("%d\n", &key)

	index := binarySearch(size, key)
	if index == -1 {
		fmt.Println("Search key not found.")
	} else {
		fmt.Println("Search key found at index:")
		fmt.Println(index)
	}
}
