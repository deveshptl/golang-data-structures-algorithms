package main

// Note - Lower priority value processes will be run first. Hence this is implemented using min heap.

import (
	"fmt"
	"math"
)

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
		if element.value >= parent.value {
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
			if leftChild.value < element.value {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = heap[rightChildIdx]
			if (rightChild.value < element.value && swap == 0) || (rightChild.value < leftChild.value && swap != 0) {
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
