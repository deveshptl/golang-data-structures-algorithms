package main

import (
	"fmt"
)

// r - rear
// n - length
// f - front
// q - queue

var r, n, f int
var q [6]string

// SimpleQueue implements simple queue
func main() {
	r = 0
	f = 0
	n = 5
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
			display()
		case 4:
			i = 1
		default:
			fmt.Println("\nCommand not recognized")
		}
	}

}

func insert() {
	if r >= n {
		fmt.Println("Overflow")
		return
	}

	r = r + 1
	var y string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s", &y)
	q[r] = y

	if f == 0 {
		f = 1
	}

	return
}

func delete() {
	if f == 0 {
		fmt.Println("Underflow")
		return
	}

	y := q[f]

	if f == r {
		f = 0
		r = 0
	} else {
		f = f + 1
	}

	fmt.Println("Element deleted", y)

	return
}

func display() {
	for i := f; i <= r; i++ {
		fmt.Print(q[i], " ")
	}
}
