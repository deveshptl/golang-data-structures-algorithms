package main

import "fmt"

// size - size of stack
// top - top pointer
// stack - stack array

var stack [5]string
var size, top int

func init() {
	size = 4
	top = -1
	for i := range stack {
		stack[i] = "nil"
	}
}

func main() {
	i := 0
	for i == 0 {
		var choice int
		fmt.Println("\n1. PUSH")
		fmt.Println("2. POP")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		fmt.Print("\nEnter you choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			push()
			fmt.Println("\nValues after push TOP:", top)
			display()
			break
		case 2:
			pop()
			fmt.Println("\nValues after pop TOP:", top)
			display()
			break
		case 3:
			fmt.Println("\nTOP:", top)
			display()
		case 4:
			i = 1
			break
		default:
			fmt.Println("Choice not recognized")
		}

	}
}

func push() {

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
	stack[top] = val
}

func pop() {

	// check for underflow
	if top == -1 {
		fmt.Println("\n-- Stack underflow --")
		return
	}

	// Print popped value
	y := stack[top]
	fmt.Println("\nValue popped:", y)

	// initialize to nil
	stack[top] = "nil"

	// decrement top pointer
	top = top - 1
}

func display() {
	for i := range stack {
		fmt.Print(stack[i], " ")
	}
	fmt.Println("")
}
