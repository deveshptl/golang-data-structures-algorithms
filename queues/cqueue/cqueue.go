package main

import (
	"fmt"
)

// r - rear pointer
// f - front pointer
// n - size of circular queue
// q - circular queue

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
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. DELETE")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			q.insert()
			fmt.Println("\nValues after insertion - FRONT:", f, "REAR:", r)
			display()
			break
		case 2:
			q.delete()
			fmt.Println("\nValues after deletion - FRONT:", f, "REAR:", r)
			display()
			break
		case 3:
			fmt.Println("\nFRONT:", f, "REAR:", r)
			display()
			break
		case 4:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (q queue) insert() {

	// reset rear pointer
	if r == n {
		r = 0
	} else {
		r = r + 1
	}

	// check for overflow
	if f == r {
		fmt.Println("\n-- Overflow --")
		if r == 0 {
			r = n
		} else {
			r = r - 1
		}
		return
	}

	// insert element
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

func (q queue) delete() {

	// check for underflow
	if f == -1 {
		fmt.Println("\n-- Underflow --")
		return
	}

	// delete the element
	y := q[f]
	q[f] = "nil"

	// check if queue is empty
	if f == r {
		f = -1
		r = -1
		fmt.Println("\nElement deleted: ", y)
		return
	}

	// increment front pointer
	if f == n {
		f = 0
	} else {
		f = f + 1
	}

	fmt.Println("\nElement deleted: ", y)

	return
}

func display() {
	fmt.Println("Displaying elements: ")
	for i := range q {
		fmt.Print(q[i], " ")
	}
	fmt.Println("")
}
