package main

import "fmt"

// InsertionSort performs insertion sort while finding the right position for the key element to be inserted
// using BinarySearch
func InsertionSort() {
	var lock int
	for i := 0; i < n-1; i++ {
		j := i
		key := elements[i+1]
		lock = BinarySearch(elements, key, 0, j)
		for j >= lock {
			elements[j+1] = elements[j]
			j--
		}
		elements[j+1] = key
	}
}

var n int
var elements []int

func main() {
	fmt.Println("\n-- Binary Insertion Sort --")
	fmt.Print("\nEnter the number of elements: ")
	fmt.Scanf("%d", &n)
	elements = make([]int, n)
	var temp int
	fmt.Println("Start entering the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		elements[i] = temp
	}
	InsertionSort()
	fmt.Println("Sorted Array is: ")
	fmt.Println(elements)
}

// BinarySearch performs binary search on given array
func BinarySearch(arr []int, item, low, high int) int {
	if high <= low {
		if item > arr[low] {
			return low + 1
		}
		return low
	}
	mid := (low + high) / 2
	if item == arr[mid] {
		return mid + 1
	}
	if item > arr[mid] {
		return BinarySearch(arr, item, mid+1, high)
	}
	return BinarySearch(arr, item, low, mid-1)
}
