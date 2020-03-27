package main

import "fmt"

func main() {
	fmt.Println("\n-- Interpolation Search --")

	var size, key int
	arr := make([]int, 0)

	fmt.Print("\nEnter the size of array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements in sorted order:")
	for i := 0; i < size; i++ {
		var value int
		fmt.Scanf("%d\n", &value)
		arr = append(arr, value)
	}
	fmt.Print("Enter the value for key: ")
	fmt.Scanf("%d\n", &key)

	pos := InterpolationSearch(arr, key)

	if pos == -1 {
		fmt.Println("\nKey", key, "not found in given array.")
	} else {
		fmt.Println("\nKey", key, "found at position", pos, "in given array.")
	}
}

// InterpolationSearch searches for a key in an sorted array
func InterpolationSearch(arr []int, key int) int {
	l := 0
	h := len(arr) - 1

	for arr[l] <= key && arr[h] >= key {
		pos := l + (((key - arr[l]) * (h - l)) / (arr[h] - arr[l]))

		if arr[pos] == key {
			return pos + 1
		} else if key < arr[pos] {
			h = pos - 1
		} else if key > arr[pos] {
			l = pos + 1
		}
	}
	return -1
}

// var arr = []int{10, 20, 25, 35, 50, 70, 85, 100, 110, 120, 125}
// var arr = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
