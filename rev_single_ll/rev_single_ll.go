package main

import "fmt"

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

// Push pushes a new node at the end
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil { // check if the list is empty
		l.head = node
	} else { // assign a new node at the end of the list
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

// Reverse reverses a list
func (l *List) Reverse() {
	var prev, next *Node
	curr := l.head

	tail := l.head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	l.tail = tail
	l.tail.next = nil
}

var l *List

func init() {
	l = &List{}
}

func main() {
	i := 0
	var choice int
	for i == 0 {
		fmt.Println("\n1. INSERT AT END")
		fmt.Println("2. DELETE FROM END")
		fmt.Println("3. REVERSE")
		fmt.Println("4. DISPLAY")
		fmt.Println("5. EXIT")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			insEnd()
			break
		case 2:
			l.Pop()
			break
		case 3:
			l.Reverse()
			break
		case 4:
			display()
			break
		case 5:
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
