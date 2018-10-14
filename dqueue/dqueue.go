package main

import (
	"fmt"
)

// r - rear
// f - front
// n - size of queue
// q - queue

var r, f, n int
var q [5]string

func init() {
	for i := range q {
		q[i] = "nil"
	}
}

func main() {
	r = -1
	f = -1
	n = 4
	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT REAR")
		fmt.Println("2. INSERT FRONT")
		fmt.Println("3. DELETE REAR")
		fmt.Println("4. DELETE FRONT")
		fmt.Println("5. DISPLAY")
		fmt.Println("6. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			insertRear()
			fmt.Println("\nValues after insert from rear FRONT:", f, "REAR:", r)
			display()
			break
		case 2:
			insertFront()
			fmt.Println("\nValues after insert from front FRONT:", f, "REAR:", r)
			display()
			break
		case 3:
			deleteRear()
			fmt.Println("\nValues after delete from rear FRONT:", f, "REAR:", r)
			display()
			break
		case 4:
			deleteFront()
			fmt.Println("\nValues after delete from front FRONT:", f, "REAR:", r)
			display()
			break
		case 5:
			fmt.Println("\nValues are FRONT:", f, "REAR:", r)
			display()
			break
		case 6:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func insertRear() {

	// check for overflow
	if r >= n {
		fmt.Println("\n-- Overflow --")
		return
	}

	// increment rear pointer
	r = r + 1

	// scan the element
	var y string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s", &y)

	// insert the element
	q[r] = y

	// set front pointer
	if f == -1 {
		f = 0
	}

	return
}

func insertFront() {

	// check for initial insertion
	if f == -1 {
		f = 0
		r = 0
		var element string
		fmt.Print("Enter the element that you want to insert: ")
		fmt.Scanf("%s", &element)
		q[f] = element
		return
	}

	// check if element can be inserted
	if f == 0 {
		fmt.Println("\n-- Element cannot be inserted --")
		return
	}

	// decrement front pointer
	f = f - 1

	// scan the element
	var element string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s", &element)

	// insert the element
	q[f] = element

	return
}

func deleteRear() {

	// check for initial condition
	if r == -1 {
		fmt.Println("\n-- Element cannot be deleted --")
		return
	}

	// delete element
	y := q[r]
	q[r] = "nil"

	// check if queue is empty
	if f == r {
		f = -1
		r = -1
	} else {
		r = r - 1
	}

	// print the deleted value
	fmt.Println("\nElement deleted:", y)

	return
}

func deleteFront() {

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
	fmt.Println("\nElement deleted:", y)

	return
}

func display() {
	for i := range q {
		fmt.Print(q[i], " ")
	}
	fmt.Println("")
}
