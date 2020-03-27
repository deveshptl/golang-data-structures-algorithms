package main

import "fmt"

var arr []int

func main() {
	var size int
	var element int
	fmt.Println("\n-- Merge Sort --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)
		arr = append(arr, element)
	}
	mergeSort(0, size-1)
	fmt.Println("Sorted Array is:")
	fmt.Println(arr)
}

func merge(l, m, r int) {
	var i, j, k int
	nl := m - l + 1
	nr := r - m
	var left, right []int

	for i = 0; i < nl; i++ {
		left = append(left, arr[l+i])
	}

	for j = 0; j < nr; j++ {
		right = append(right, arr[m+1+j])
	}

	i = 0
	j = 0
	k = l

	for i < nl && j < nr {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
			k++
		} else {
			arr[k] = right[j]
			j++
			k++
		}
	}

	for i < nl {
		arr[k] = left[i]
		i++
		k++
	}
	for j < nr {
		arr[k] = right[j]
		j++
		k++
	}
}

func mergeSort(l, r int) {
	var mid int
	if l < r {
		mid = (l + r) / 2
		mergeSort(l, mid)
		mergeSort(mid+1, r)
		merge(l, mid, r)
	}
}
