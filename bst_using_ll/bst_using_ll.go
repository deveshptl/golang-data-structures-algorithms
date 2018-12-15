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

// Delete deletes a node preserving bst
func (bst *BinarySearchTree) Delete(value int) {
	// TODO: implement delete function
}

// Find finds a key from the bst
func (bst *BinarySearchTree) Find(value int) {
	// TODO: implement find function
}

// BFS prints the tree using bfs traversal
func (bst *BinarySearchTree) BFS() []int {
	if bst.root == nil {
		fmt.Print("\n-- Tree is empty. --")
		return nil
	}
	var queue []*Node
	var data []int
	node := bst.root
	queue = append(queue, node)

	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		for node.freq != 0 {
			data = append(data, node.value)
			node.freq--
		}
		fmt.Println(node.left, node.right)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return data
}

// DFS prints the tree using dfs traversal
func (bst *BinarySearchTree) DFS() []int {
	// TODO: implement dfs traversal
	return nil
}

var bst *BinarySearchTree

func init() {
	bst = &BinarySearchTree{}
}

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. DELETE")
		fmt.Println("3. DISPLAY using BFS")
		fmt.Println("4. DISPLAY using BFS")
		fmt.Println("5. FIND")
		fmt.Println("5. EXIT")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			insNode()
			break
		case 2:
			deleteNode()
			break
		case 3:
			data := bst.BFS()
			if data != nil {
				fmt.Println(data)
			}
			break
		case 4:
			data := bst.DFS()
			if data != nil {
				fmt.Println(data)
			}
			break
		case 5:
			findNode()
			break
		case 6:
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
	bst.Insert(element)
}

func deleteNode() {
	var element int
	fmt.Print("Enter the value that you want to delete: ")
	fmt.Scanf("%d", &element)
	bst.Delete(element)
}

func findNode() {
	var element int
	fmt.Print("Enter the value of the key: ")
	fmt.Scanf("%d", &element)
	bst.Find(element)
}
