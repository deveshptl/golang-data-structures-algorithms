package main

import "fmt"

// PosToIn will convert a given Postfix notation to Infix
func PosToIn() {
	for i := range expression {
		char := string(expression[i])
		if (expression[i] >= 65 && expression[i] <= 90) || (expression[i] >= 97 && expression[i] <= 122) {
			stack = append(stack, char)
		} else {
			operand1 := stack[len(stack)-1] // read top of stack
			stack = stack[0 : len(stack)-1] // pop from stack
			operand2 := stack[len(stack)-1] // read top of stack
			stack = stack[0 : len(stack)-1] // pop from stack
			temp := "(" + operand2 + char + operand1 + ")"
			stack = append(stack, temp)
		}
	}
	infix := stack[len(stack)-1] // read top of stack which is the final infix notation
	fmt.Println("\n" + infix)
}

var stack []string
var expression string

func main() {
	fmt.Println("\n-- Postfix to Infix expression conversion --")
	i := 0
	for i == 0 {
		fmt.Print("\nEnter the expression: ")
		fmt.Scanf("%s", &expression)
		PosToIn()
		var choice string
		fmt.Print("\nDo you want to continue y/n? ")
		fmt.Scanf("%s", &choice)
		if choice == "N" || choice == "n" {
			break
		} else if choice == "Y" || choice == "y" {
			continue
		} else {
			fmt.Println("Command not recognized.")
		}
	}
}
