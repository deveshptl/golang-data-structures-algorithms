package main

import (
	"fmt"
	"math"
)

// Insert inserts a node in a heap
func Insert(value int) {
	heap = append(heap, value)
	minHeapify()
}

func minHeapify() {
	idx := len(heap) - 1
	element := heap[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := heap[parentIdx]
		if element >= parent {
			break
		}
		heap[parentIdx] = element
		heap[idx] = parent
		idx = parentIdx
	}
}

// ExtractMin will remove a node from heap
func ExtractMin() int {
	if len(heap) == 0 || heap[0] == 0 {
		fmt.Println("\n-- Heap is empty. --")
		return 0
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
		var leftChild, rightChild, swap int

		if leftChildIdx < length {
			leftChild = heap[leftChildIdx]
			if leftChild < element {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = heap[rightChildIdx]
			if (rightChild < element && swap == 0) || (rightChild < leftChild && swap != 0) {
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

var heap []int

func main() {
	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. REMOVE")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		var ch int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &ch)
		switch ch {
		case 1:
			insertNode()
			break
		case 2:
			ExtractMin()
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
	var ch int
	fmt.Print("Enter the element that you want to insert: ")
	fmt.Scanf("%d\n", &ch)
	Insert(ch)
}

func display() {
	fmt.Println("")
	fmt.Println(heap)
}
