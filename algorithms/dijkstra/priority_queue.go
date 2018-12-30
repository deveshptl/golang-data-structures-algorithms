package main

import (
	"fmt"
	"math"
)

// QueueNode is a single node in a queue
type QueueNode struct {
	name     string
	priority int
}

// Enqueue inserts a node in a heap
func Enqueue(node *QueueNode) {
	pqueue = append(pqueue, node)
	minHeapify()
}

func minHeapify() {
	idx := len(pqueue) - 1
	element := pqueue[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := pqueue[parentIdx]
		if element.priority >= parent.priority {
			break
		}
		pqueue[parentIdx] = element
		pqueue[idx] = parent
		idx = parentIdx
	}
}

// Dequeue will remove a node from heap
func Dequeue() *QueueNode {
	if len(pqueue) == 0 || pqueue[0] == nil {
		fmt.Println("\n-- Heap is empty. --")
		return nil
	}
	min := pqueue[0]
	end := pqueue[len(pqueue)-1]
	pqueue = pqueue[0 : len(pqueue)-1]
	if len(pqueue) > 0 {
		pqueue[0] = end
		bubbleDown()
	}
	return min
}

func bubbleDown() {
	idx := 0
	length := len(pqueue)
	element := pqueue[0]
	for true {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := (2 * idx) + 2
		var leftChild, rightChild *QueueNode
		var swap int

		if leftChildIdx < length {
			leftChild = pqueue[leftChildIdx]
			if leftChild.priority < element.priority {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = pqueue[rightChildIdx]
			if (rightChild.priority < element.priority && swap == 0) || (rightChild.priority < leftChild.priority && swap != 0) {
				swap = rightChildIdx
			}
		}

		if swap == 0 {
			break
		}
		pqueue[idx] = pqueue[swap]
		pqueue[swap] = element
		idx = swap
	}
}

func insertNodeInQueue(name string, priority int) {
	newNode := &QueueNode{name: name, priority: priority}
	Enqueue(newNode)
}

func display() {
	if len(pqueue) == 0 {
		fmt.Println("\n-- Heap is empty. --")
		return
	}
	for i := range pqueue {
		fmt.Println(pqueue[i].name, pqueue[i].priority)
	}
}
