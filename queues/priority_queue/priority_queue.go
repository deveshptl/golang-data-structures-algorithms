package main

// Note - Lower priority value processes will be run first. Hence this is implemented using min heap.

import (
	"fmt"
	"math"
)

// Node is a single node in a queue
type Node struct {
	name     string
	priority int
}

// Enqueue inserts a node in a heap
func Enqueue(node *Node) {
	heap = append(heap, node)
	minHeapify()
}

func minHeapify() {
	idx := len(heap) - 1
	element := heap[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := heap[parentIdx]
		if element.priority >= parent.priority {
			break
		}
		heap[parentIdx] = element
		heap[idx] = parent
		idx = parentIdx
	}
}

// Dequeue will remove a node from heap
func Dequeue() *Node {
	if len(heap) == 0 || heap[0] == nil {
		fmt.Println("\n-- Heap is empty. --")
		return nil
	}
	min := heap[0]
	end := heap[len(heap)-1]
	heap = heap[0 : len(heap)-1]
	if len(heap) > 0 {
		heap[0] = end
		bubbleDown()
	}
	return min
}

func bubbleDown() {
	idx := 0
	length := len(heap)
	element := heap[0]
	for true {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := (2 * idx) + 2
		var leftChild, rightChild *Node
		var swap int

		if leftChildIdx < length {
			leftChild = heap[leftChildIdx]
			if leftChild.priority < element.priority {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = heap[rightChildIdx]
			if (rightChild.priority < element.priority && swap == 0) || (rightChild.priority < leftChild.priority && swap != 0) {
				swap = rightChildIdx
			}
		}

		if swap == 0 {
			break
		}
		heap[idx] = heap[swap]
		heap[swap] = element
		idx = swap
	}
}

var heap []*Node

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. REMOVE")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		var ch int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &ch)
		switch ch {
		case 1:
			insertNode()
			break
		case 2:
			node := Dequeue()
			fmt.Println("Process removed: ", node)
			break
		case 3:
			display()
			break
		case 4:
			i = 1
			break
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func insertNode() {
	var name string
	var priority int
	fmt.Print("Enter the name of the process: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter the priority of the process: ")
	fmt.Scanf("%d", &priority)
	fmt.Println(name, priority)
	newNode := &Node{name: name, priority: priority}
	Enqueue(newNode)
}

func display() {
	if len(heap) == 0 {
		fmt.Println("\n-- Heap is empty. --")
		return
	}
	for i := range heap {
		fmt.Println(heap[i].name, heap[i].priority)
	}
}
