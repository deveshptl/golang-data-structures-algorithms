package main

import "fmt"

func main() {
	fmt.Println("\n-- n C r Combinations with repetition --")

	var n, r int
	arr := make([]string, 0)

	fmt.Print("\nEnter the value of n: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Start entering unique values of n:")
	for i := 0; i < n; i++ {
		var temp string
		fmt.Scanf("%s", &temp)
		arr = append(arr, temp)
	}
	fmt.Print("Enter the value of r: ")
	fmt.Scanf("%d", &r)

	data := make([]string, r)

	Combinations(arr, data, 0, n-1, 0, r)
}

// Combinations prints all possible nCr combinations of given array
func Combinations(arr, data []string, start, end, index, r int) {
	if index == r {
		fmt.Println(data)
		return
	}

	for i := start; i <= end; i++ {
		data[index] = arr[i]
		Combinations(arr, data, i, end, index+1, r)
	}

}
