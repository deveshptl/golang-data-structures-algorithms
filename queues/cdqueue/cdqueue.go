package main

import (
	"fmt"
)

// r- rear pointer
// f - front pointer
// n - size of circular double ended queue
// q - circular double ended queue

type queue []string

var r, f, n int
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
		fmt.Println("\n1. INSERT REAR")
		fmt.Println("2. INSERT FRONT")
		fmt.Println("3. DELETE REAR")
		fmt.Println("4. DELETE FRONT")
		fmt.Println("5. DISPLAY")
		fmt.Println("6. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			q.insertRear()
			fmt.Println("\nValues after insert from rear FRONT:", f, "REAR:", r)
			display()
		case 2:
			q.insertFront()
			fmt.Println("\nValues after insert from front FRONT:", f, "REAR:", r)
			display()
		case 3:
			q.deleteRear()
			fmt.Println("\nValues after delete from rear FRONT:", f, "REAR:", r)
			display()
		case 4:
			q.deleteFront()
			fmt.Println("\nValues after delete from front FRONT:", f, "REAR:", r)
			display()
		case 5:
			fmt.Println("\nValues are FRONT:", f, "REAR:", r)
			display()
		case 6:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (q queue) insertRear() {

	// reset rear pointer
	if r == n {
		r = 0
	} else {
		r++
	}

	// check for overflow
	if f == r {
		fmt.Println("\n-- Overflow --")
		if r == 0 {
			r = n
		} else {
			r--
		}
		return
	}

	// insert element
	var y string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s\n", &y)
	q[r] = y

	// set front pointer
	if f == -1 {
		f = 0
	}

}

func (q queue) insertFront() {

	// check for initial insertion
	if f == -1 {
		f = 0
		r = 0
		var element string
		fmt.Print("Enter the element that you want to insert: ")
		fmt.Scanf("%s\n", &element)
		q[f] = element
		return
	}

	// reset front pointer
	if f == 0 {
		f = n
	} else {
		f--
	}

	// check for overflow
	if f == r {
		if f == n {
			f = 0
		} else {
			f++
		}
		fmt.Println("\n-- Overflow --")
		return
	}

	// insert element
	var element string
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%s\n", &element)
	q[f] = element

}

func (q queue) deleteRear() {

	// check for initial insertion
	if r == -1 {
		fmt.Println("\n-- Underflow --")
		return
	}

	// delete element
	y := q[r]
	q[r] = "nil"

	// check if queue is empty
	if f == r {
		f = -1
		r = -1
		fmt.Println("\nElement deleted:", y)
		return
	}

	// set rear pointer
	if r == 0 {
		r = n
	} else {
		r--
	}

	fmt.Println("\nElement deleted:", y)

}

func (q queue) deleteFront() {

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
		f++
	}
	fmt.Println("\nElement deleted: ", y)

}

func display() {
	for i := range q {
		fmt.Print(q[i], " ")
	}
	fmt.Println("")
}
