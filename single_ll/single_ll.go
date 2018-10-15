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
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

// Pop deletes a node from the end of the list
func (l *List) Pop() {
	if l.head == nil {
		fmt.Println("\n-- Underflow --")
	} else if l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		for list := l.head; list != nil; list = list.next {
			if list.next.next == nil {
				list.next = nil
				l.tail = list
				return
			}
		}
	}
}

// Shift deletes a node from the begining of the list
func (l *List) Shift() {
	if l.head == nil {
		fmt.Println("\n-- Underflow --")
	} else if l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
	}
}

// Unshift inserts a node at the begining of the list
func (l *List) Unshift(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
}

// InsMiddle inserts a node in the middle of the list
func (l *List) InsMiddle(value, place int) {
	if place == 1 || l.head == nil {
		l.Unshift(value)
	} else {
		list := l.head
		node := &Node{value: value}
		for i := 0; i <= place; i++ {
			if list == l.tail && i != place-1 {
				l.tail.next = node
				l.tail = node
				return
			} else if i+1 == place-1 {
				temp := list.next
				list.next = node
				node.next = temp
				return
			}
			list = list.next
		}
	}
}

// DelMiddle deletes a node from the middle of the list
func (l *List) DelMiddle(place int) {
	if place == 1 || l.head == nil {
		l.Shift()
	} else if l.head.next == nil && place > 1 {
		fmt.Println("\n-- Index out of range", place, "--")
	} else {
		list := l.head
		for i := 0; i <= place; i++ {
			if list.next == nil {
				fmt.Println("\n-- Index out of range", place, "--")
				return
			} else if i+1 == place-1 {
				if list.next == l.tail {
					list.next = nil
					l.tail = list
				} else {
					list.next = list.next.next
				}
				return
			} else {
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
