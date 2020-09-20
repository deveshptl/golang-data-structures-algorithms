package main

import (
	"fmt"
	"math"
)

func shellSort(size int) {
	gap := int(math.Floor(float64(size / 2)))
	for {
		for i := 0; i < size-((size-1)%gap); i++ {

			// swap the elements
			if arr[i] > arr[gap+i] {
				arr[i], arr[gap+i] = arr[gap+i], arr[i]

				j := i
				// start comparing with previous elements
				for {
					// swap the elements
					if j-gap >= 0 && arr[j-gap] > arr[j] {
						arr[j-gap], arr[j] = arr[j], arr[j-gap]
						j--
					} else {
						break
					}
				}
			}

			// check whether to break the loop
			if i == (size-1)-gap {
				break
			}
		}

		gap = int(math.Floor(float64(gap / 2))) // update the gap

		// break the main loop when gap is 0
		if gap == 0 {
			break
		}
	}
}

var arr []int

func main() {
	var size, element int
	fmt.Println("\n-- Shell Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)
		arr = append(arr, element)
	}
	shellSort(size)
	fmt.Println("Sorted array is:")
	fmt.Println(arr)
}
