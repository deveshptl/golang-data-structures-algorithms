package main

// BFS is a bfs traversal technique
func BFS() []int {
	var data []int
	for i := range bst {
		if bst[i] != nil {
			count := bst[i].freq
			for count != 0 {
				data = append(data, bst[i].value)
				count--
			}
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
	if bst[index] != nil {
		count := bst[index].freq
		for count != 0 {
			data = append(data, bst[index].value)
			count--
		}
		data = traversePre(data, (index*2)+1)
		data = traversePre(data, (index*2)+2)
	}
	return data
}

// traverses a tree in in-order and recursively calls itself
func traverseIn(data []int, index int) []int {
	if bst[index] != nil {
		data = traverseIn(data, (index*2)+1)
		count := bst[index].freq
		for count != 0 {
			data = append(data, bst[index].value)
			count--
		}
		data = traverseIn(data, (index*2)+2)
	}
	return data
}

// traverses a tree in post-order and recursively calls itself
func traversePos(data []int, index int) []int {
	if bst[index] != nil {
		data = traversePos(data, (index*2)+1)
		data = traversePos(data, (index*2)+2)
		count := bst[index].freq
		for count != 0 {
			data = append(data, bst[index].value)
			count--
		}
	}
	return data
}
