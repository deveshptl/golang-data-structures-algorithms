package main

import "fmt"

// Node is a node of bt
type Node struct {
	value  int
	height int
}

// Insert inserts a node in a bt
func Insert(value, index int) {
	flag := 0
	newNode := &Node{value: value}
	for i := range bt {
		if bt[i] == nil {
			bt[i] = newNode
			flag = 1
			return
		}
		if i == 99 && flag == 0 {
			fmt.Println("\n-- Not enough space in array. --")
			return
		}
	}
}

// Find finds a key and prints its information
func Find(key int) {
	for i := range bt {
		if bt[i] == nil {
			fmt.Println("\n-- Key NOT found. --")
			return
		} else if bt[(i*2)+1] != nil && bt[(i*2)+1].value == key {
			index := (i * 2) + 1
			fmt.Println("\n-- Key found. --")
			fmt.Println("Key info is:", bt[index])
			fmt.Println("Parent is:", bt[i])
			fmt.Println("Sibling is:", bt[(i*2)+2])
			fmt.Println("Left Child is:", bt[(index*2)+1])
			fmt.Println("Right Child is:", bt[(index*2)+2])
			return
		} else if bt[(i*2)+2] != nil && bt[(i*2)+2].value == key {
			index := (i * 2) + 2
			fmt.Println("\n-- Key found. --")
			fmt.Println("Key info is:", bt[index])
			fmt.Println("Parent is:", bt[i])
			fmt.Println("Sibling is:", bt[(i*2)+1])
			fmt.Println("Left Child is:", bt[(index*2)+1])
			fmt.Println("Right Child is:", bt[(index*2)+2])
			return
		}
	}
}

// BFS is a bfs traversal technique
func BFS() []int {
	var data []int
	for i := range bt {
		if bt[i] != nil {

			data = append(data, bt[i].value)
		}
	}
	return data
}

// DFS is a dfs traversal technique
func DFS(traversalType string) []int {
	var data []int
	if traversalType == "pre" {
		data = traversePre(data, 0)
	} else if traversalType == "in" {
		data = traverseIn(data, 0)
	} else if traversalType == "pos" {
		data = traversePos(data, 0)
	}
	return data
}

// traverses a tree in pre-order and recursively calls itself
func traversePre(data []int, index int) []int {
	if bt[index] != nil {
		data = append(data, bt[index].value)
		data = traversePre(data, (index*2)+1)
		data = traversePre(data, (index*2)+2)
	}
	return data
}

// traverses a tree in in-order and recursively calls itself
func traverseIn(data []int, index int) []int {
	if bt[index] != nil {
		data = traverseIn(data, (index*2)+1)
		data = append(data, bt[index].value)
		data = traverseIn(data, (index*2)+2)
	}
	return data
}

// traverses a tree in post-order and recursively calls itself
func traversePos(data []int, index int) []int {
	if bt[index] != nil {
		data = traversePos(data, (index*2)+1)
		data = traversePos(data, (index*2)+2)
		data = append(data, bt[index].value)
	}
	return data
}

var bt [100]*Node
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
