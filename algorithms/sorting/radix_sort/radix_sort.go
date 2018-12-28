package main

import (
	"fmt"
	"math"
)

var arr []int

func main() {
	var size int
	var element int
	fmt.Println("\n-- Radix Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &element)
		arr = append(arr, element)
	}
	radixSort()
	fmt.Println("Sorted Array is:")
	fmt.Println(arr)
}

func radixSort() {
	maxDigitCount := mostDigits(arr)

	for k := 0; k < maxDigitCount; k++ {
		var bucket [10][]int
		for i := 0; i < len(arr); i++ {
			digit := getDigit(arr[i], k)
			bucket[digit] = append(bucket[digit], arr[i])
		}

		// concat numbers back to original array
		arr = arr[0:0]
		for j := 0; j < 10; j++ {
			arr = append(arr, bucket[j]...)
		}
	}
}

// getDigit returns an integer whose value is equivalent to index position in given number
func getDigit(value, index int) int {
	return int(math.Floor(math.Abs(float64(value)/math.Pow(float64(10), float64(index))))) % 10
}

// digitCount returns an integer which is the count of number of digits in a given number
func digitCount(value int) int {
	if value == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(value)))) + 1
}

// mostDigits returns a digits count whose length is maximum or say whose value is maximum
func mostDigits(nums []int) int {
	maxDigits := 0
	for i := 0; i < len(nums); i++ {
		maxDigits = int(math.Max(float64(maxDigits), float64(digitCount(nums[i]))))
	}
	return maxDigits
}
