package main

import (
	"fmt"
)

// List implements single linked list
type List struct {
	head *Node
	tail *Node
}

// Node is a single node of a linked list
type Node struct {
	value int
	next  *Node
}

// Insert inserts a node preserving the order of the list
func (l *List) Insert(value int) {
	node := &Node{value: value}
	if l.head == nil { // check if list is empty
		l.head = node
		l.tail = node
	} else if l.head.value >= value { // check if head value is greater than insertion value
		node.next = l.head
		l.head = node
	} else if l.tail.value <= value { // check if tail value is less than insertion value
		l.tail.next = node
		l.tail = node
	} else { // loop over the list
		list := l.head
		for !(list.value <= value && list.next.value >= value) {
			// increment list pointer
			list = list.next
		}

		// insert the node
		temp := list.next
		list.next = node
		node.next = temp
	}
}

// Delete deletes a node preserving the order of the list
func (l *List) Delete(place int) {
	if l.head == nil { // check if list is empty
		fmt.Println("\n-- Underflow --")
		return
	} else if l.head.next == nil && place == 1 { // check if only one node is present and index value is 1
		l.head = nil
		l.tail = nil
		return
	} else if place == 1 { // check if index value is 1
		l.head = l.head.next
		return
	} else if l.head.next == nil && place > 1 { // check if only one node is present and index value is greater than 1
		fmt.Println("\n-- Index out of range --")
		return
	}
	list := l.head
	for i := 0; i < place-2; i++ { // loop over the list
		if list.next == l.tail {
			fmt.Println("\n-- Index out of range --")
			return
		}

		// increment next pointer
		list = list.next
	}

	if list.next == l.tail { // check if value of list next pointer is pointing to tail
		l.tail = list
	}

	// change the next pointer value
	list.next = list.next.next
}

var l *List

func init() {
	l = &List{}
}

func main() {
	i := 0
	var choice int

	for i == 0 {
		fmt.Println("\n1. INSERT PRESERVING ORDER")
		fmt.Println("2. DELETE PRESERVING ORDER")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			nodeInsert()
			break
		case 2:
			nodeDelete()
			break
		case 3:
			display()
			break
		case 4:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func nodeInsert() {
	var element int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d", &element)
	l.Insert(element)
}

func nodeDelete() {
	var place int
	fmt.Print("Enter the index place from where you to delete the node: ")
	fmt.Scanf("%d", &place)
	l.Delete(place)
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
