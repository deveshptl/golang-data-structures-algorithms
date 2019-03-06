package main

import "fmt"

// size - size of stack
// top - top pointer
// stack - stack array

type stack []string

var s stack
var size, top int

func init() {
	size = 4
	top = -1
	s = make(stack, 5)
	for i := range s {
		s[i] = "nil"
	}
}

func main() {
	i := 0
	for i == 0 {
		var choice int
		fmt.Println("\n1. PUSH")
		fmt.Println("2. POP")
		fmt.Println("3. PEEK")
		fmt.Println("4. DISPLAY")
		fmt.Println("5. isEmpty")
		fmt.Println("6. EXIT")
		fmt.Print("\nEnter you choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			s.push()
			fmt.Println("\nValues after push TOP:", top)
			s.display()
		case 2:
			s.pop()
			fmt.Println("\nValues after pop TOP:", top)
			s.display()
		case 3:
			s.peek()
			fmt.Println("\nValues after peek TOP:", top)
			s.display()
		case 4:
			fmt.Println("\nTOP:", top)
			s.display()
		case 5:
			s.isEmpty()
		case 6:
			i = 1
		default:
			fmt.Println("Choice not recognized")
		}

	}
}

func (s stack) push() {

	// check for overflow
	if size == top {
		fmt.Println("\n-- Stack overflow --")
		return
	}

	// scan for element
	var val string
	fmt.Print("Enter the value that you want to PUSH: ")
	fmt.Scanf("%s", &val)

	// increment top pointer
	top = top + 1

	// insert the element
	s[top] = val
	return
}

func (s stack) pop() {

	// check for underflow
	if top == -1 {
		fmt.Println("\n-- Stack underflow --")
		return
	}

	// print popped value
	y := s[top]
	fmt.Println("\nValue popped:", y)

	// initialize to nil
	s[top] = "nil"

	// decrement top pointer
	top = top - 1
	return
}

func (s stack) peek() {
	// check for underflow
	if top == -1 {
		fmt.Println("\n-- Stack underflow --")
		return
	}

	// print the peeked value
	y := s[top]
	fmt.Println("\nValue peeked:", y)
	return
}

func (s stack) isEmpty() {
	if top == -1 {
		fmt.Println("\ntrue")
		return
	}
	fmt.Println("\nfalse")
	return
}

func (s stack) display() {
	for i := range s {
		fmt.Print(s[i], " ")
	}
	fmt.Println("")
}
