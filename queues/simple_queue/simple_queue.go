package main

import (
	"fmt"
)

// r - rear
// n - length
// f - front
// q - queue

type queue []string

var r, n, f int
var q queue

func init() {
	r = -1
	f = -1
	n = 4
	q = make(queue, 5)
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
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			q.insert()
			fmt.Println("\nValues after insertion FRONT:", f, "REAR:", r)
			display()
		case 2:
			q.delete()
			fmt.Println("\nValues after deletion FRONT:", f, "REAR:", r)
			display()
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

func (q queue) insert() {

	// check for overflow
	if r >= n {
		fmt.Println("\n-- Overflow --")
		return
	}

	// increment rear pointer
	r++

	// insert the element
	var y string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s\n", &y)
	q[r] = y

	// set front pointer
	if f == -1 {
		f = 0
	}
}

func (q queue) delete() {

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
		f++
	}

	// print deleted element
	fmt.Println("Element deleted", y)
}

func display() {
	for i := range q {
		fmt.Print(q[i], " ")
	}
	fmt.Println("")
}
