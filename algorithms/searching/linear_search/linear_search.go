package main

import "fmt"

func linearSearch(size, key int) []int {
	var index []int
	for i := 0; i < size; i++ {
		if key == arr[i] {
			index = append(index, i+1)
		}
	}
	return index
}

var arr []int

func main() {
	var size, element, key int
	fmt.Println("\n-- Linear Search --")
	fmt.Print("\nEnter the size of an array: ")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Start entering elements in sorted manner:")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d\n", &element)
		arr = append(arr, element)
	}
	fmt.Print("Enter the search key: ")
	fmt.Scanf("%d\n", &key)
	index := linearSearch(size, key)
	if len(index) == 0 {
		fmt.Println("Key not fount")
	} else {
		fmt.Println("Key found at following index/indices:")
		fmt.Println(index)
	}
}
