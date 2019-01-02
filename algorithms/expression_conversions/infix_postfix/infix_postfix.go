package main

import "fmt"

// InToPos converts an infix expression to postfix expression
func InToPos() {
	if !validate(expression) {
		fmt.Println("\n-- Invalid expression --")
		return
	}
	postfix := ""
	stack = append(stack, "N")
	for i := range expression {
		char := string(expression[i])
		if (expression[i] >= 65 && expression[i] <= 90) || (expression[i] >= 97 && expression[i] <= 122) {
			postfix += char
		} else if char == "(" {
			stack = append(stack, char)
		} else if char == ")" {
			for stack[len(stack)-1] != "N" && stack[len(stack)-1] != "(" {
				newChar := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				postfix += newChar
			}
			if stack[len(stack)-1] == "(" {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			for stack[len(stack)-1] != "N" && Prec(stack[len(stack)-1]) >= Prec(char) {
				newChar := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				postfix += newChar
			}
			stack = append(stack, char)
		}
	}
	for stack[len(stack)-1] != "N" {
		newChar := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		postfix += newChar
	}
	fmt.Println("\n" + postfix)
}

var stack []string
var expression string

func main() {
	fmt.Println("\n-- Infix to Postfix expression conversion --")
	i := 0
	for i == 0 {
		fmt.Print("\nEnter the expression: ")
		fmt.Scanf("%s", &expression)
		InToPos()
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

// Prec will return a precedence of a given operator
func Prec(c string) int {
	if c == "^" {
		return 3
	} else if c == "*" || c == "/" {
		return 2
	} else if c == "+" || c == "-" {
		return 1
	}
	return -1
}

func isOperator(s string) bool {
	if s == "^" || s == "*" || s == "/" || s == "+" || s == "-" {
		return true
	}
	return false
}

func validate(c string) bool {
	var check int
	for i := range c {
		if isOperator(string(c[i])) || string(c[i]) == ")" {
			check += -1
		} else {
			check++
		}
	}
	if check == 1 {
		return true
	}
	return false
}
