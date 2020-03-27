package main

import (
	"fmt"
	"math"
)

// InsertionSort performs insertion sort on each bucket
func InsertionSort(a []float64, n int) []float64 {
	for i := 0; i < n-1; i++ {
		j := i
		key := a[i+1]
		for j >= 0 && key < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	return a
}

// BucketSort sort the array using bucket sort algorithm
func BucketSort() {
	bucket = make([][]float64, size)
	for i := 0; i < size; i++ {
		index := int(math.Floor((arr[i] / max) * (float64(size - 1))))
		bucket[index] = append(bucket[index], arr[i])
	}

	for i := 0; i < size; i++ {
		if i == 0 {
			bucket[0] = InsertionSort(bucket[i], len(bucket[i]))
		} else {
			bucket[0] = append(bucket[0], InsertionSort(bucket[i], len(bucket[i]))...)
		}
	}
}

var size int
var max float64
var arr []float64
var bucket [][]float64

func init() {
	arr = make([]float64, 0)
	bucket = make([][]float64, 0)
}

func main() {
	fmt.Println("\n-- Insertion Sort --")
	fmt.Print("Enter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering the elements:")
	var temp float64
	for i := 0; i < size; i++ {
		if i == 0 {
			max = temp
		}
		fmt.Scanf("%f\n", &temp)
		arr = append(arr, temp)
		if max < temp {
			max = temp
		}
	}

	BucketSort()

	fmt.Println(bucket[0])
}
