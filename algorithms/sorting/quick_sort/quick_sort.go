package main

import "fmt"

var arr []int

func main() {
	var size int
	var element int
	fmt.Println("\n-- Quick Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &element)
		arr = append(arr, element)
	}

	quicksort(0, size-1)

	fmt.Println("Sorted Array is:")
	fmt.Println(arr)
}

func partition(first, last int) int {
	pivot := arr[first]
	i := first
	j := last + 1

	for i < j {
		i++
		j--
		for arr[i] < pivot && i < last {
			i++
		}
		for arr[j] > pivot && j > first {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[first] = arr[first], arr[j]
	return j
}

func quicksort(first, last int) {
	var sp int
	if first < last {
		sp = partition(first, last)
		quicksort(first, sp-1)
		quicksort(sp+1, last)
	}
}
