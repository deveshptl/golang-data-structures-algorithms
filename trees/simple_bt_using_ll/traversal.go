package main

import "fmt"

// BFS prints the tree using bfs traversal
func (bt *BinaryTree) BFS() []int {
	if bt.root == nil {
		fmt.Print("\n-- Tree is empty. --")
		return nil
	}
	var queue []*Node
	var data []int
	node := bt.root
	queue = append(queue, node)

	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		data = append(data, node.value)
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
func (bt *BinaryTree) DFS(traverseType string) []int {
	if bt.root == nil {
		fmt.Println("\n-- Tree is empty. --")
		return nil
	}

	var data []int
	current := bt.root
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
	data = append(data, node.value)
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
	data = append(data, node.value)
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
	data = append(data, node.value)
	return data
}
