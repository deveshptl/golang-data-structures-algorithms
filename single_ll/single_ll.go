package main

import (
	"fmt"
)

// List implements linked list
type List struct {
	head *Node
	tail *Node
}

// Node is a single node of linked list
type Node struct {
	value int
	next  *Node
}

// Push pushes new node at the end of the list
func (l *List) Push(value int) {
	node := &Node{value: value}

	if l.head == nil { // check if list is empty
		l.head = node
	} else { // assign next pointer of tail to new node
		l.tail.next = node
	}

	// assign new tail as
	l.tail = node
}

// Pop deletes a node from the end of the list
func (l *List) Pop() {

	if l.head == nil { // check if list is empty
		fmt.Println("\n-- Underflow --")
	} else if l.head.next == nil { // check if only one node is present
		l.head = nil
		l.tail = nil
	} else { // loop over the list
		for list := l.head; list != nil; list = list.next {
			if list.next.next == nil { // check if list pointer is pointing to second last node of the list
				list.next = nil
				l.tail = list
				return
			}
		}
	}
}

// Shift deletes a node from the begining of the list
func (l *List) Shift() {

	// check if list is empty
	if l.head == nil {
		fmt.Println("\n-- Underflow --")
	} else if l.head.next == nil { // check if only one node is present
		l.head = nil
		l.tail = nil
	} else { // assign new head to the list
		l.head = l.head.next
	}
}

// Unshift inserts a node at the begining of the list
func (l *List) Unshift(value int) {
	node := &Node{value: value}

	if l.head == nil { // check if list is empty
		l.head = node
		l.tail = node
	} else { // assign new node as head
		node.next = l.head
		l.head = node
	}
}

// InsMiddle inserts a node in the middle of the list
func (l *List) InsMiddle(value, place int) {

	// check if insertion place is 1 or list is empty
	if place == 1 || l.head == nil {
		l.Unshift(value)
	} else { // loop over list
		list := l.head
		node := &Node{value: value}
		for i := 0; i <= place; i++ {

			if list == l.tail && i != place-1 { // check if tail is encountered
				l.tail.next = node
				l.tail = node
				return
			} else if i+1 == place-1 { // check if list pointer is pointing to correct node
				temp := list.next
				list.next = node
				node.next = temp
				return
			}

			// increment list pointer to next node
			list = list.next
		}
	}
}

// DelMiddle deletes a node from the middle of the list
func (l *List) DelMiddle(place int) {

	// check if deletion place is 1 or list is empty
	if place == 1 || l.head == nil {
		l.Shift()
	} else if l.head.next == nil && place > 1 { // check if only one node is present and deletion place is greater than 1
		fmt.Println("\n-- Index out of range", place, "--")
	} else { // loop over the list
		list := l.head
		for i := 0; i <= place; i++ {

			// check if tail is encountered
			if list.next == nil {
				fmt.Println("\n-- Index out of range", place, "--")
				return
			} else if i+1 == place-1 { // check is list is pointing to correct node

				// check if list is pointing to second last element
				if list.next == l.tail {
					list.next = nil
					l.tail = list
				} else { // assign new next pointer to nodes
					list.next = list.next.next
				}
				return
			} else { // increment list pointer to next node
				list = list.next
			}
		}
	}
}

// Next - returns the next pointer of a node
func (n *Node) Next() *Node {
	return n.next
}

var l *List

func init() {
	l = &List{}
}

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT AT HEAD")
		fmt.Println("2. INSERT AT MIDDLE")
		fmt.Println("3. INSERT AT END")
		fmt.Println("4. DELETE FROM HEAD")
		fmt.Println("5. DELETE FROM MIDDLE")
		fmt.Println("6. DELETE FROM END")
		fmt.Println("7. DISPLAY")
		fmt.Println("8. EXIT")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			insStart()
			break
		case 2:
			insMiddle()
			break
		case 3:
			insEnd()
			break
		case 4:
			l.Shift()
			break
		case 5:
			delMiddle()
			break
		case 6:
			l.Pop()
			break
		case 7:
			display()
			break
		case 8:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func insEnd() {
	var element int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d", &element)
	l.Push(element)
}

func insStart() {
	var element int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d", &element)
	l.Unshift(element)
}

func insMiddle() {
	var element, place int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d", &element)
	fmt.Print("Enter the index where you want to insert the element: ")
	fmt.Scanf("%d", &place)
	l.InsMiddle(element, place)
}

func delMiddle() {
	var place int
	fmt.Print("Enter the index from where you want to delete the element: ")
	fmt.Scanf("%d", &place)
	l.DelMiddle(place)
}

func display() {
	fmt.Println("\nHEAD:", l.head, "TAIL:", l.tail)
	if l.head != nil {
		for list := l.head; list != nil; list = list.next {
			fmt.Println(list.value, list.next)
		}
	} else {
		fmt.Println("-- List is empty --")
	}
	fmt.Println("")
}
