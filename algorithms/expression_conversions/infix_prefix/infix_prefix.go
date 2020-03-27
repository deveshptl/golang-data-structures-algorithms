package main

import (
	"fmt"
)

// InToPos converts an infix expression to prefix expression
func InToPos(expression string) string {
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
	return postfix
}

// InToPre will reverse the expression and will call postfix on it
// will then reverse again the returned expression from postfix call
func InToPre() {
	if !validate(inputExp) {
		fmt.Println("\n-- Invalid expression --")
		return
	}
	reverseStr := reverse(inputExp)
	postfix := InToPos(reverseStr)
	prefix := reverse(postfix)
	fmt.Println("\n" + prefix)
}

var stack []string
var inputExp string

func main() {
	fmt.Println("\n-- Infix to Prefix expression conversion --")
	i := 0
	for i == 0 {
		fmt.Print("\nEnter the expression: ")
		fmt.Scanf("%s\n", &inputExp)
		InToPre()
		var choice string
		fmt.Print("\nDo you want to continue y/n? ")
		fmt.Scanf("%s\n", &choice)
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
	return 0
}

func isOperator(s string) bool {
	if s == "^" || s == "*" || s == "/" || s == "+" || s == "-" {
		return true
	}
	return false
}

func reverse(s string) string {
	reverseStr := ""
	j := len(s) - 1
	for j != -1 {
		if string(s[j]) == ")" {
			reverseStr += "("
		} else if string(s[j]) == "(" {
			reverseStr += ")"
		} else {
			reverseStr += string(s[j])
		}
		j--
	}
	return reverseStr
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
