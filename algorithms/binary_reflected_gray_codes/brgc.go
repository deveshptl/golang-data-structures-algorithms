package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("\n-- Binary Reflected Gray Codes --")
	var value int
	fmt.Print("\nEnter a value: ")
	fmt.Scanf("%d\n", &value)
	if value < 1 {
		log.Fatal("Invalid value. Value must be greater than 0.")
	}
	l := BRGC(value)
	fmt.Println("Gray codes for", value, "bits are:")
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
}

// BRGC generates binary gray codes of given integer
func BRGC(n int) []string {

	// return 0 and 1 as list if n == 1
	if n == 1 {
		return []string{"0", "1"}
	}

	// create list l1 by recursively calling on n-1
	l1 := BRGC(n - 1)
	l2 := make([]string, 0)

	// create list l2 in reverse order and add prefix 1 on each item
	for i := len(l1) - 1; i >= 0; i-- {
		l2 = append(l2, "1"+l1[i])
	}

	// add prefix 0 on each item of list l2
	for i := 0; i < len(l1); i++ {
		l1[i] = "0" + l1[i]
	}

	// append l2 to l1
	l1 = append(l1, l2...)

	// return generated list
	return l1
}
