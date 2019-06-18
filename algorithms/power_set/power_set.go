package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\n-- Generating Power Set (Set of Subsets)--")

	var value int

	fmt.Print("\nEnter a value: ")
	fmt.Scanf("%d", &value)

	list := PowerSet(value)

	fmt.Println("All possible subsets upto", value, "are:")
	for i := 0; i < len(list); i++ {
		str := strings.Split(list[i], "")
		fmt.Println("{" + strings.Join(str, ", ") + "}")
	}
}

// PowerSet generates a set of all subsets of a given integer
func PowerSet(n int) []string {

	if n == 0 {
		return []string{"0"}
	}

	list := PowerSet(n - 1)

	newList := make([]string, 0)

	str := strconv.Itoa(n)
	newList = append(newList, list...)
	newList = append(newList, str)
	for i := 1; i < len(list); i++ {
		newList = append(newList, list[i]+str)
	}
	return newList
}
