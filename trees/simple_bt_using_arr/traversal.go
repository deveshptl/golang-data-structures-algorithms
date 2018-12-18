package main

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
