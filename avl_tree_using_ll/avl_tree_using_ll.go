package main

import (
	"fmt"
	"math"
)

// Node is a node in a tree
type Node struct {
	value  int
	freq   int
	height int
	left   *Node
	right  *Node
}

// AVLTree is Avl tree implementation
type AVLTree struct {
	root *Node
}

// Height returs the height of a node
func Height(node *Node) int {
	if node == nil {
		return 0
	}

	return node.height
}

// BalancingFactor returns a balancing factor of a node
func BalancingFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return Height(node.left) - Height(node.right)
}

// Insert inserts a node preserving avl
func Insert(node *Node, value int) *Node {
	if node == nil {
		nn := &Node{value: value, height: 1, freq: 1}
		return nn
	}

	if value > node.value {
		node.right = Insert(node.right, value)
	} else if value < node.value {
		node.left = Insert(node.left, value)
	} else if value == node.value {
		node.freq++
	}
	node.height = int(math.Max(float64(Height(node.left)), float64(Height(node.right))) + 1)

	bf := BalancingFactor(node)

	// LL case
	if bf > 1 && value < node.left.value {
		return rightRotate(node)
	}

	// RR case
	if bf < -1 && value > node.right.value {
		return leftRotate(node)
	}

	// LR case
	if bf > 1 && value > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// RL case
	if bf < -1 && value < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

// leftRotate rotates a given node to left
func leftRotate(node *Node) *Node {
	b := node.right
	t := b.left

	b.left = node
	node.right = t

	node.height = int(math.Max(float64(Height(node.left)), float64(Height(node.right))) + 1)
	b.height = int(math.Max(float64(Height(b.left)), float64(Height(b.right))) + 1)

	return b
}

// rightRotate rotates a given node to right
func rightRotate(node *Node) *Node {
	b := node.left
	t := b.right

	b.right = node
	node.left = t

	node.height = int(math.Max(float64(Height(node.left)), float64(Height(node.right))) + 1)
	node.height = int(math.Max(float64(Height(b.left)), float64(Height(b.right))) + 1)

	return b
}

// Find finds a key from the avl
func (avl *AVLTree) Find(value int) {
	current := avl.root
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
			fmt.Println("\n-- Key is root itself. --")
			fmt.Println("Left child is: ", current.left)
			fmt.Println("Right child is: ", current.right)
			return
		}
	}
}

var avl *AVLTree

func init() {
	avl = &AVLTree{}
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
			fmt.Println("\n", avl.BFS())
			break
		case 3:
			fmt.Println("\n", avl.DFS("pre"))
			break
		case 4:
			fmt.Println("\n", avl.DFS("in"))
			break
		case 5:
			fmt.Println("\n", avl.DFS("pos"))
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
	avl.root = Insert(avl.root, element)
}

func findNode() {
	var element int
	fmt.Print("Enter the value of the key: ")
	fmt.Scanf("%d", &element)
	avl.Find(element)
}
