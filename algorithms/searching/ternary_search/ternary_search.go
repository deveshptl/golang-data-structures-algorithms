package main

import "fmt"

func main() {
	fmt.Println("\n-- Ternary Search --")

	var size, key int
	arr := make([]int, 0)

	fmt.Print("\nEnter the size of array: ")
	fmt.Scanf("%d", &size)
	fmt.Println("Start entering elements in sorted order:")
	for i := 0; i < size; i++ {
		var value int
		fmt.Scanf("%d", &value)
		arr = append(arr, value)
	}
	fmt.Print("Enter the value for key: ")
	fmt.Scanf("%d", &key)

	pos := TernarySearch(arr, 0, size-1, key)

	if pos == -1 {
		fmt.Println("\nKey", key, "not found in given array.")
	} else {
		fmt.Println("\nKey", key, "found at position", pos, "in given array.")
	}

}

// TernarySearch searches for a key in a given array
func TernarySearch(arr []int, l, h, key int) int {
	if h < l {
		return -1
	}

	mid1 := l + (h-l)/3
	mid2 := h - (h-l)/3

	if arr[mid1] == key {
		return mid1 + 1
	} else if arr[mid2] == key {
		return mid2 + 1
	} else if key < arr[mid1] {
		return TernarySearch(arr, l, mid1-1, key)
	} else if key > arr[mid1] && key < arr[mid2] {
		return TernarySearch(arr, mid1+1, mid2-1, key)
	} else if key > arr[mid2] {
		return TernarySearch(arr, mid2+1, h, key)
	}
	return -1
}

// var arr = []int{1, 3, 4, 6, 8, 9, 12, 15, 17, 18, 23, 25, 27, 31, 32}
