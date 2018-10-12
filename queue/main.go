package main

import (
	"fmt"
)

func main() {
	i := 0
	var choice int
	for i == 0 {
		fmt.Println("1. SIMPLE QUEUE")
		fmt.Println("2. EXIT")
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			SimpleQueue()
		case 2:
			i = 1
		default:
			fmt.Println("Command not recognized")
		}
	}
}
