package main

import (
	"fmt"
	"math"
)

// HeapNode is a single node in a queue
type HeapNode struct {
	name     string
	priority float64
}

// Enqueue inserts a node in a heap
func Enqueue(node *HeapNode) {
	heap = append(heap, node)
	hashMap[node.name] = len(heap) - 1
	minHeapify(-1)
}

func minHeapify(idx int) {
	if idx == -1 {
		idx = len(heap) - 1
	}
	element := heap[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := heap[parentIdx]
		if element.priority >= parent.priority {
			break
		}
		heap[parentIdx] = element
		hashMap[element.name] = parentIdx
		heap[idx] = parent
		hashMap[parent.name] = idx
		idx = parentIdx
	}
}

// Dequeue will remove a node from heap
func Dequeue() *HeapNode {
	if len(heap) == 0 || heap[0] == nil {
		fmt.Println("\n-- Heap is empty. --")
		return nil
	}
	min := heap[0]
	end := heap[len(heap)-1]
	heap = heap[0 : len(heap)-1]
	hashMap[end.name] = 0
	if len(heap) > 0 {
		heap[0] = end
		bubbleDown()
	}
	delete(hashMap, min.name)
	return min
}

func bubbleDown() {
	idx := 0
	length := len(heap)
	element := heap[0]
	for {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := (2 * idx) + 2
		var leftChild, rightChild *HeapNode
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
		hashMap[heap[swap].name] = idx
		heap[swap] = element
		hashMap[element.name] = swap
		idx = swap
	}
}

// DecreaseKey will decrease the priority of a given node positioned by index in a heap
// will then call minHeapify to satisfy the min heap property
func DecreaseKey(index int, newValue float64) {
	heap[index].priority = newValue
	minHeapify(index)
}

var heap []*HeapNode

// hasMap maps heap values to heap indexes
var hashMap map[string]int

func init() {
	hashMap = make(map[string]int)
}

func insertNode(name string, priority float64) {
	newNode := &HeapNode{name: name, priority: priority}
	Enqueue(newNode)
}
