package main

import "fmt"

// Node is a node in a tree
type Node struct {
	value int
	freq  int
	left  *Node
	right *Node
}

// BinarySearchTree is bst implementation
type BinarySearchTree struct {
	root *Node
}

// Insert inserts a node preserving bst
func (bst *BinarySearchTree) Insert(value int) {
	node := &Node{value: value, freq: 1}
	if bst.root == nil {
		bst.root = node
	} else {
		current := bst.root
		for true {
			if value == current.value {
				fmt.Println("\n-- Same value was found. Freq. incremented. --")
				current.freq++
				return
			} else if value < current.value {
				if current.left == nil {
					current.left = node
					return
				}
				current = current.left

			} else if value > current.value {
				if current.right == nil {
					current.right = node
					return
				}
				current = current.right
			}
		}
	}
}

// Find finds a key from the bst
func (bst *BinarySearchTree) Find(value int) {
	current := bst.root
	for true {
		if value < current.value {
			if current.left == nil {
				fmt.Println("\n-- Key not found. --")
				return
			} else if value == current.left.value {
				fmt.Println("\n-- Key found. --")
				fmt.Println("Parent is: ", current.value)
				fmt.Println("Sibling is: ", current.right)
				fmt.Println("Left child is: ", current.left.left)
				fmt.Println("Right child is: ", current.left.right)
				return
			}
			current = current.left
		} else if value > current.value {
			if current.right == nil {
				fmt.Println("\n-- Key not found. --")
				return
			}
			if value == current.right.value {
				fmt.Println("\n-- Key found. --")
				fmt.Println("Parent is: ", current.value)
				fmt.Println("Sibling is: ", current.left)
				fmt.Println("Left child is: ", current.right.left)
				fmt.Println("Right child is: ", current.right.right)
				return
			}
			current = current.right
		} else if value == current.value {
			fmt.Println("Key is root itself")
			fmt.Println("Left child is: ", current.left)
			fmt.Println("Right child is: ", current.right)
			return
		}
	}
}

var bst *BinarySearchTree

func init() {
	bst = &BinarySearchTree{}
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
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			insNode()
			break
		case 2:
			fmt.Println("\n", bst.BFS())
			break
		case 3:
			fmt.Println("\n", bst.DFS("pre"))
			break
		case 4:
			fmt.Println("\n", bst.DFS("in"))
			break
		case 5:
			fmt.Println("\n", bst.DFS("pos"))
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
	fmt.Scanf("%d\n", &element)
	bst.Insert(element)
}

func findNode() {
	var element int
	fmt.Print("Enter the value of the key: ")
	fmt.Scanf("%d\n", &element)
	bst.Find(element)
}
