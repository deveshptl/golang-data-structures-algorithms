package main

import (
	"fmt"
	"math"
)

func maxHeapify() {
	idx := len(arr) - 1
	element := arr[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := arr[parentIdx]
		if element <= parent {
			break
		}
		arr[parentIdx] = element
		arr[idx] = parent
		idx = parentIdx
	}
}

// ExtractMax will remove a node from heap
func ExtractMax() int {
	if len(arr) == 0 || arr[0] == 0 {
		fmt.Println("\n-- Heap is empty. --")
		return 0
	}
	max := arr[0]
	end := arr[len(arr)-1]
	arr = arr[0 : len(arr)-1]
	if len(arr) > 0 {
		arr[0] = end
		bubbleDown()
	}
	return max
}

func bubbleDown() {
	idx := 0
	length := len(arr)
	element := arr[0]
	for {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := (2 * idx) + 2
		var leftChild, rightChild, swap int

		if leftChildIdx < length {
			leftChild = arr[leftChildIdx]
			if leftChild > element {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = arr[rightChildIdx]
			if (rightChild > element && swap == 0) || (rightChild > leftChild && swap != 0) {
				swap = rightChildIdx
			}
		}

		if swap == 0 {
			break
		}
		arr[idx] = arr[swap]
		arr[swap] = element
		idx = swap
	}
}

// HeapSort sorts the arr using heap sort technique and returns the sorted arr
func HeapSort(size int) []int {
	sorted := make([]int, size)
	for i := 0; i < size; i++ {
		max := ExtractMax()
		sorted[size-1-i] = max
	}
	return sorted
}

var arr []int

func main() {
	var size, element int
	fmt.Println("\n-- Heap Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements:")

	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)
		arr = append(arr, element)
		maxHeapify()
	}

	sorted := HeapSort(size)
	fmt.Println("Sorted array is:")
	fmt.Println(sorted)
}
