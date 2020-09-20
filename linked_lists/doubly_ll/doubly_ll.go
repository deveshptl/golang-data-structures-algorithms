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
	prev  *Node
	value int
	next  *Node
}

// Push pushes new node at the end of the list
func (l *List) Push(value int) {
	node := &Node{value: value}

	if l.head == nil { // checks if the list is empty
		l.head = node
	} else { // assign next pointer of tail to new node
		node.prev = l.tail
		l.tail.next = node
	}

	// assign tail to newly inserted node
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
		return
	}

	if l.head.next == nil { // check if only one node is present
		l.head = nil
		l.tail = nil
		return
	}

	// change head pointer and head's prev pointer
	l.head = l.head.next
	l.head.prev = nil
}

// Unshift inserts a node at the begining of the list
func (l *List) Unshift(value int) {
	node := &Node{value: value}

	if l.head == nil { // check if list is empty
		l.head = node
		l.tail = node
	} else { // add new node to the list and assign it as head
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
}

// InsMiddle inserts a node in the middle of the list
func (l *List) InsMiddle(value, place int) {

	if (l.head == nil) && place > 1 { // check if insertion place greater than 1 and list is empty or only 1 node is present
		fmt.Println("\n-- Index out of range --")
	} else if place == 1 || l.head == nil { // check if list is empty and place is 1
		l.Unshift(value)
	} else { // loop over the list
		list := l.head
		node := &Node{value: value}
		for i := 0; i < place-2; i++ {
			if list == l.tail { // check if next pointer of list is pointing to tail and value of i and place are such that index out of range is satisfied
				fmt.Println("\n-- Index out of range --")
				return
			}

			// increment next pointer
			list = list.next
		}

		// check if list is pointing to tail
		if list == l.tail {
			l.tail.next = node
			node.prev = l.tail
			l.tail = node
		} else { // insert the node
			temp := list.next
			list.next = node
			node.prev = list
			node.next = temp
			node.next.prev = node
		}
	}
}

// DelMiddle deletes a node from the middle of the list
func (l *List) DelMiddle(place int) {

	// check if deletion place is 1 or list is empty
	if place == 1 || l.head == nil {
		l.Shift()
	} else if l.head.next == nil && place > 1 { // check if only one node is present and deletion place is greater than 1
		fmt.Println("\n-- Index out of range --")
	} else { // loop over the list
		list := l.head
		for i := 0; i < place-2; i++ {

			if list.next == l.tail { // check if next pointer of list is pointing to tail
				fmt.Println("\n-- Index out of range --")
				return
			}

			// increment next pointer
			list = list.next

		}

		if list.next == l.tail { // check if next pointer of list is pointing to tail
			list.next = nil
			l.tail = list
		} else { // delete the node
			list.next = list.next.next
			list.next.prev = list
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
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			insStart()
		case 2:
			insMiddle()
		case 3:
			insEnd()
		case 4:
			l.Shift()
		case 5:
			delMiddle()
		case 6:
			l.Pop()
		case 7:
			display()
		case 8:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func insEnd() {
	var element int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d\n", &element)
	l.Push(element)
}

func insStart() {
	var element int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d\n", &element)
	l.Unshift(element)
}

func insMiddle() {
	var element, place int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d\n", &element)
	fmt.Print("Enter the index where you want to insert the element: ")
	fmt.Scanf("%d\n", &place)
	l.InsMiddle(element, place)
}

func delMiddle() {
	var place int
	fmt.Print("Enter the index from where you want to delete the element: ")
	fmt.Scanf("%d\n", &place)
	l.DelMiddle(place)
}

func display() {
	fmt.Println("\nHEAD:", l.head, "TAIL:", l.tail)
	if l.head != nil {
		for list := l.head; list != nil; list = list.next {
			fmt.Println(list.value, list.prev, list.next)
		}
	} else {
		fmt.Println("-- List is empty --")
	}
	fmt.Println("")
}
