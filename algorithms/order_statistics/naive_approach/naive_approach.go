package main

import "fmt"

var arr []int

func main() {
	fmt.Println("\n-- Order Statistics Naive Approach --")

	// initialization
	var size int
	newArr := make([]int, 0)
	var k int

	// user inputs
	fmt.Print("\nEnter the size of array: ")
	fmt.Scanf("%d", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		var value int
		fmt.Scanf("%d", &value)
		newArr = append(newArr, value)
	}
	fmt.Print("Enter the value of k: ")
	fmt.Scanf("%d", &k)

	naiveSelect(newArr, k)

	fmt.Println(k, "smallest element is:", arr[0])
}

func naiveSelect(newArr []int, k int) {
	for i := 0; i < k; i++ {
		arr = append(arr, newArr[i])
		maxHeapify()
	}
	for i := k; i < len(newArr); i++ {
		if arr[0] < newArr[i] {
			continue
		}
		arr = append(arr, newArr[i])
		maxHeapify()
		ExtractMax()
	}
}

// sorted - {1, 2, 3, 5, 6, 7, 8, 9, 10, 12, 14, 15, 16, 19}
// var newArr = []int{5, 15, 8, 3, 1, 9, 10, 6, 7, 12, 14, 2, 16, 19}
