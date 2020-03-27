package main

import "fmt"

// Node is a node in a tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// BinaryTree is bst implementation
type BinaryTree struct {
	root *Node
}

// Insert inserts a node preserving bst
func (bt *BinaryTree) Insert(value int) {
	newNode := &Node{value: value}
	if bt.root == nil {
		bt.root = newNode
		return
	}

	var queue []*Node
	current := bt.root
	queue = append(queue, current)

	for len(queue) != 0 {
		current = queue[0]
		queue = queue[1:]
		if current.left != nil {
			queue = append(queue, current.left)
		} else {
			current.left = newNode
			return
		}
		if current.right != nil {
			queue = append(queue, current.right)
		} else {
			current.right = newNode
			return
		}
	}
	return
}

// Find finds a key from the bt
func (bt *BinaryTree) Find(value int) {
	if bt.root == nil {
		fmt.Print("\n-- Tree is empty. --")
		return
	}
	flag := 0
	var queue []*Node
	node := bt.root
	queue = append(queue, node)

	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		if node.left != nil && node.left.value == value {
			flag = 1
			fmt.Println("\n-- Key found. --")
			fmt.Println("Key info is:", node.left)
			fmt.Println("Parent is:", node)
			fmt.Println("Sibling is:", node.right)
			fmt.Println("Left Child is:", node.left.left)
			fmt.Println("Right Child is:", node.left.right)
		} else if node.right != nil && node.right.value == value {
			flag = 1
			fmt.Println("\n-- Key found. --")
			fmt.Println("Key info is:", node.right)
			fmt.Println("Parent is:", node)
			fmt.Println("Sibling is:", node.left)
			fmt.Println("Left Child is:", node.right.left)
			fmt.Println("Right Child is:", node.right.right)
		}
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	if flag == 0 {
		fmt.Println("\n-- Key NOT found. --")
	}
}

var bt *BinaryTree

func init() {
	bt = &BinaryTree{}
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
			fmt.Println("\n", bt.BFS())
			break
		case 3:
			fmt.Println("\n", bt.DFS("pre"))
			break
		case 4:
			fmt.Println("\n", bt.DFS("in"))
			break
		case 5:
			fmt.Println("\n", bt.DFS("pos"))
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
	bt.Insert(element)
}

func findNode() {
	var element int
	fmt.Print("Enter the value of the key: ")
	fmt.Scanf("%d\n", &element)
	bt.Find(element)
}
