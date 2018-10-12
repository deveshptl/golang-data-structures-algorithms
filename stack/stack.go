package main

import "fmt"

var len int

func main() {
	len = 5
	top := -1
	x := make([]string, 6)
	selection := 1
	for selection == 1 {
		var choice int
		fmt.Println("\n1. PUSH")
		fmt.Println("2. POP")
		fmt.Println("3. EXIT")
		fmt.Print("\nEnter you choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			x, top = push(x, top)
			fmt.Println("\nValues after PUSH ", x, top)
			break
		case 2:
			x, top = pop(x, top)
			fmt.Println("\nValues after POP ", x, top)
			break
		case 3:
			selection = 0
			break
		default:
			fmt.Println("Choice not recognized")
		}

	}

	fmt.Println("Final values: ", x, top)
	fmt.Println("Exited...")

}

func push(x []string, top int) ([]string, int) {
	if len == top {
		fmt.Println("Stack overflow")
		return x, top
	}
	var val string
	fmt.Print("Enter the value that you want to PUSH: ")
	fmt.Scanf("%s", &val)
	top = top + 1
	x[top] = val
	return x, top
}

func pop(x []string, top int) ([]string, int) {
	if top == -1 {
		fmt.Println("Stack underflow")
		return x, top
	}

	top = top - 1
	return x[:top+1], top
}
