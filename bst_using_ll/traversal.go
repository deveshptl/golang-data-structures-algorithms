package main

import "fmt"

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
		count := node.freq
		for count != 0 {
			data = append(data, node.value)
			count--
		}
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
func (bst *BinarySearchTree) DFS(traverseType string) []int {
	if bst.root == nil {
		fmt.Println("\n-- Tree is empty. --")
		return nil
	}

	var data []int
	current := bst.root
	if traverseType == "pre" {
		data = traversePre(data, current)
	} else if traverseType == "in" {
		data = traverseIn(data, current)
	} else if traverseType == "pos" {
		data = traversePos(data, current)
	}
	return data
}

func traversePre(data []int, node *Node) []int {
	count := node.freq
	for count != 0 {
		data = append(data, node.value)
		count--
	}
	if node.left != nil {
		data = traversePre(data, node.left)
	}
	if node.right != nil {
		data = traversePre(data, node.right)
	}
	return data
}

func traverseIn(data []int, node *Node) []int {
	if node.left != nil {
		data = traverseIn(data, node.left)
	}
	count := node.freq
	for count != 0 {
		data = append(data, node.value)
		count--
	}
	if node.right != nil {
		data = traverseIn(data, node.right)
	}
	return data
}

func traversePos(data []int, node *Node) []int {
	if node.left != nil {
		data = traversePos(data, node.left)
	}
	if node.right != nil {
		data = traversePos(data, node.right)
	}
	count := node.freq
	for count != 0 {
		data = append(data, node.value)
		count--
	}
	return data
}
