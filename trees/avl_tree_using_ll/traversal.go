package main

import "fmt"

// BFS prints the tree using bfs traversal
func (avl *AVLTree) BFS() []int {
	if avl.root == nil {
		fmt.Print("\n-- Tree is empty. --")
		return nil
	}
	var queue []*Node
	var data []int
	node := avl.root
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
func (avl *AVLTree) DFS(traverseType string) []int {
	if avl.root == nil {
		fmt.Println("\n-- Tree is empty. --")
		return nil
	}

	var data []int
	current := avl.root
	if traverseType == "pre" {
		data = traversePre(data, current)
	} else if traverseType == "in" {
		data = traverseIn(data, current)
	} else if traverseType == "pos" {
		data = traversePos(data, current)
	}
	return data
}

// traverses a node in pre-order and recursively calls itself
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

// traverses a node in in-order and recursively calls itself
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

// traverses a node in pos-order and recursively calls itself
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
