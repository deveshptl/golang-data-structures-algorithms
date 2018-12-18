package main

import (
	"fmt"
)

// r - rear
// n - length
// f - front
// q - queue

var r, n, f int
var q [5]string

func init() {
	r = -1
	f = -1
	n = 4
	for i := range q {
		q[i] = "nil"
	}
}

func main() {
	i := 0
	var choice int
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. DELETE")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			insert()
			fmt.Println("\nValues after insertion FRONT:", f, "REAR:", r)
			display()
			break
		case 2:
			delete()
			fmt.Println("\nValues after deletion FRONT:", f, "REAR:", r)
			display()
			break
		case 3:
			fmt.Println("\nFRONT:", f, "REAR:", r)
			display()
		case 4:
			i = 1
		default:
			fmt.Println("\nCommand not recognized")
		}
	}

}

func insert() {

	// check for overflow
	if r >= n {
		fmt.Println("\n-- Overflow --")
		return
	}

	// increment rear pointer
	r = r + 1

	// insert the element
	var y string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s", &y)
	q[r] = y

	// set front pointer
	if f == -1 {
		f = 0
	}

	return
}

func delete() {

	// check for underflow
	if f == -1 {
		fmt.Println("\n-- Underflow --")
		return
	}

	// delete element
	y := q[f]
	q[f] = "nil"

	// check if queue is empty
	if f == r {
		f = -1
		r = -1
	} else {
		f = f + 1
	}

	// print deleted element
	fmt.Println("Element deleted", y)

	return
}

func display() {
	for i := range q {
		fmt.Print(q[i], " ")
	}
	fmt.Println("")
}
