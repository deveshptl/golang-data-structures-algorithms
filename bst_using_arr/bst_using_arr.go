package main

import "fmt"

// Node is a node of bst
type Node struct {
	value  int
	freq   int
	height int
}

// Insert inserts a node in a bst
func Insert(value, index int) {
	newNode := &Node{value: value, freq: 1}
	if index >= size {
		fmt.Println("\n-- Not enough space in array. --")
		return
	}
	if bst[index] == nil {
		bst[index] = newNode
	} else if bst[index].value == value {
		bst[index].freq++
	} else {
		if value < bst[index].value {
			Insert(value, (index*2)+1)
		} else {
			Insert(value, (index*2)+2)
		}
	}
}

// Find finds a key and prints its information
func Find(key int) {
	i := 0
	for true {
		if bst[i] != nil {
			if bst[i].value == key {
				index := i
				var sibling *Node
				if i%2 == 0 {
					i = (i / 2) - 1
					sibling = bst[(i*2)+1]
				} else {
					i = i / 2
					sibling = bst[(i*2)+2]
				}
				fmt.Println("\n-- Key found. --")
				fmt.Println("Key info is: ", bst[index])
				fmt.Println("Parent is: ", bst[i])
				fmt.Println("Sibling is: ", sibling)
				fmt.Println("Left Child is: ", bst[(index*2)+1])
				fmt.Println("Right Child is: ", bst[(index*2)+2])
				return
			} else if bst[i].value < key {
				i = (i * 2) + 2
			} else if bst[i].value > key {
				i = (i * 2) + 1
			}
		} else {
			fmt.Println("\n-- Key NOT found. --")
			return
		}
	}
}

var bst [100]*Node
var size int

func init() {
	size = 100
}

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. DISPLAY using BFS")
		fmt.Println("3. DISPLAY using DFS Pre-order")
		fmt.Println("4. DISPLAY using DFS In-order")
		fmt.Println("5. DISPLAY using DFS Post-order")
		fmt.Println("6. FIND")
		fmt.Println("7. EXIT")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			insNode()
			break
		case 2:
			fmt.Println("\n", BFS())
			break
		case 3:
			fmt.Println("\n", DFS("pre"))
			break
		case 4:
			fmt.Println("\n", DFS("in"))
			break
		case 5:
			fmt.Println("\n", DFS("pos"))
			break
		case 6:
			findNode()
			break
		case 7:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func insNode() {
	var element int
	fmt.Print("Enter the node value that you want to insert: ")
	fmt.Scanf("%d", &element)
	Insert(element, 0)
}

func findNode() {
	var element int
	fmt.Print("Enter the value of the key: ")
	fmt.Scanf("%d", &element)
	Find(element)
}
