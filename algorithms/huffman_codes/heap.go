package main

import (
	"fmt"
	"math"
)

// Node represents a single node in a heap
type Node struct {
	char  string
	freq  int
	left  *Node
	right *Node
}

type heap []Node

// Insert inserts a node in a heap
func (h heap) Insert(node Node) heap {
	h = append(h, node)
	h = h.minHeapify()
	return h
}

func (h heap) minHeapify() heap {
	idx := len(h) - 1
	element := h[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := h[parentIdx]
		if element.freq >= parent.freq {
			break
		}
		h[parentIdx] = element
		h[idx] = parent
		idx = parentIdx
	}
	return h
}

// ExtractMin will remove a node from heap
func (h heap) ExtractMin() (*Node, heap) {
	if len(h) == 0 {
		fmt.Println("\n-- Heap is empty. --")
		return &Node{}, h
	}
	min := h[0]
	end := h[len(h)-1]
	h = h[0 : len(h)-1]
	if len(h) > 0 {
		h[0] = end
		h.bubbleDown()
	}
	return &min, h
}

func (h heap) bubbleDown() heap {
	idx := 0
	length := len(h)
	element := h[0]
	for true {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := (2 * idx) + 2
		var leftChild, rightChild Node
		var swap int

		if leftChildIdx < length {
			leftChild = h[leftChildIdx]
			if leftChild.freq < element.freq {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = h[rightChildIdx]
			if (rightChild.freq < element.freq && swap == 0) || (rightChild.freq < leftChild.freq && swap != 0) {
				swap = rightChildIdx
			}
		}

		if swap == 0 {
			break
		}
		h[idx] = h[swap]
		h[swap] = element
		idx = swap
	}
	return h
}

func (h heap) insertNodes() heap {
	var char string
	var freq int
	fmt.Print("\nChar: ")
	fmt.Scanf("%s", &char)
	fmt.Print("Freq: ")
	fmt.Scanf("%d", &freq)
	h = h.Insert(Node{char: char, freq: freq})
	return h
}

func (h heap) display() {
	fmt.Println("")
	fmt.Println(h)
}
