package main

import "fmt"

// PreToPos will convert a given Prefix notation to Postfix
func PreToPos() {
	if !validate(expression) {
		fmt.Println("\n-- Invalid expression --")
		return
	}
	str := reverse(expression)
	for i := range str {
		char := string(str[i])
		if (str[i] >= 65 && str[i] <= 90) || (str[i] >= 97 && str[i] <= 122) {
			stack = append(stack, char)
		} else {
			operand1 := stack[len(stack)-1] // read top of stack
			stack = stack[0 : len(stack)-1] // pop from stack
			operand2 := stack[len(stack)-1] // read top of stack
			stack = stack[0 : len(stack)-1] // pop from stack
			temp := operand1 + operand2 + char
			stack = append(stack, temp)
		}
	}
	infix := stack[len(stack)-1] // read top of stack which is the final infix notation
	fmt.Println("\n" + infix)
}

var stack []string
var expression string

func main() {
	fmt.Println("\n-- Prefix to Postfix expression conversion --")
	i := 0
	for i == 0 {
		fmt.Print("\nEnter the expression: ")
		fmt.Scanf("%s", &expression)
		PreToPos()
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
