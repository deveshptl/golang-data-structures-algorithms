package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("\n-- Order Statistics Quick Select --")

	// initializations
	var size, k int
	arr := make([]int, 0)

	// user inputs
	fmt.Print("\nEnter the size of array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		var value int
		fmt.Scanf("%d\n", &value)
		arr = append(arr, value)
	}
	fmt.Print("Enter the value for k: ")
	fmt.Scanf("%d\n", &k)

	value := QuickSelect(arr, k)
	fmt.Println(k, "smallest element is:", value)
}

// QuickSelect finds kth smallest element in an array using quick sort - partially
func QuickSelect(arr []int, k int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	pos := partition(arr)

	if pos == k-1 {
		return arr[pos]
	} else if k-1 < pos {
		return QuickSelect(arr[:pos], k)
	}

	return QuickSelect(arr[pos+1:], k-pos-1)
}

func partition(arr []int) int {

	// generate a random number, time as a seed value
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(arr) - 1)

	// push the element to last
	arr[n], arr[len(arr)-1] = arr[len(arr)-1], arr[n]
	n = len(arr) - 1

	pivot := arr[n]
	index := 0
	for i := 0; i < n; i++ {
		if arr[i] < pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[n] = arr[n], arr[index]
	return index
}

// sorted - {1, 2, 3, 5, 6, 7, 8, 9, 10, 12, 14, 15, 16, 19}
// var arr = []int{5, 15, 8, 3, 1, 9, 10, 6, 7, 12, 14, 2, 16, 19}
