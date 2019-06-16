package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("\n-- Worst Case Linear Time Order Statistics --")
	arr := make([]float64, 0)
	var value int
	var n int
	var k int
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Start entering elements:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &value)
		arr = append(arr, float64(value))
	}
	fmt.Print("Enter the value for k: ")
	fmt.Scanf("%d", &k)
	if k < 1 || k >= len(arr) {
		fmt.Println("You must enter a valid value for k.")
		return
	}
	fmt.Println(k, "smallest element is:", kthSmallest(arr, k))
}

// WCLT Worst Case Linear Time Order Statistics
func kthSmallest(arr []float64, k int) float64 {
	if len(arr) == 1 {
		return arr[0]
	}

	median := make([]float64, 0)
	upto := int(math.Ceil(float64(len(arr)) / 5))
	// dividing array in n/5 subgroups
	for i := 0; i < upto; i++ {
		if (i*5)+5 > len(arr) {
			newArr := arr[i*5:]
			sort.Float64s(newArr)
			median = append(median, findMedian(newArr, len(newArr)))
		} else {
			newArr := arr[i*5 : (i*5)+5]
			sort.Float64s(newArr)
			median = append(median, findMedian(newArr, 5))
		}
	}

	// finding median of medians
	medianOfMedian := kthSmallest(median, len(median)/2)

	rank := findRank(arr, medianOfMedian)

	pos := partition(arr, rank)

	if pos == k-1 {
		return arr[pos]
	} else if pos > k-1 {
		return kthSmallest(arr[:pos], k)
	}

	return kthSmallest(arr[pos+1:], k-pos-1)
}

func partition(arr []float64, rank int) int {
	arr[rank], arr[len(arr)-1] = arr[len(arr)-1], arr[rank]
	pivot := arr[len(arr)-1]
	index := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[len(arr)-1] = arr[len(arr)-1], arr[index]
	return index
}

func findMedian(arr []float64, n int) float64 {
	return arr[n/2]
}

func findRank(arr []float64, value float64) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

// Array
// {5, 15, 8, 3, 1, 9, 10, 6, 7, 12, 14, 2, 16, 19}
// sorted - {1, 2, 3, 5, 6, 7, 8, 9, 10, 12, 14, 15, 16, 19}
