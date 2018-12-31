package main

import "fmt"

func main() {
	var n int
	fmt.Println("\n-- Tower of Hanoi --")
	fmt.Print("\nEnter the number of discs: ")
	fmt.Scanf("%d", &n)
	Toh("A", "B", "C", n)
}

// Toh recursively calls iteself to implment tower of hanoi algorithm
func Toh(origin, intermediate, destination string, n int) {
	if n == 1 {
		fmt.Println("Move disc", n, "from", origin, "->", destination)
		return
	}
	Toh(origin, destination, intermediate, n-1)
	fmt.Println("Move disc", n, "from", origin, "->", destination)
	Toh(intermediate, origin, destination, n-1)
}
