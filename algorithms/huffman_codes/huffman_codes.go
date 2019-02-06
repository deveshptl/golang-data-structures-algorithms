package main

import (
	"fmt"
	"strings"
)

func (h heap) genHeap() *Node {
	var curr, prev *Node
	for len(h) > 1 {
		prev, h = h.ExtractMin()
		curr, h = h.ExtractMin()
		newNode := Node{left: prev, right: curr, freq: (prev.freq + curr.freq)}
		h = h.Insert(newNode)
	}
	return &h[0]
}

func (ht *Node) printHuffman() {
	var code []string
	code = make([]string, 0)
	ht.printHuffmanUtil(ht, code)
}

func (ht *Node) printHuffmanUtil(node *Node, code []string) {
	if node.left == nil && node.right == nil {
		// print
		fmt.Println("\n", node.char, "=>", strings.Join(code, ""))
		code = code[0:0]
	}
	if node.left != nil {
		code = append(code, "0")
		ht.printHuffmanUtil(node.left, code)
		code = code[0 : len(code)-1]
	}
	if node.right != nil {
		code = append(code, "1")
		ht.printHuffmanUtil(node.right, code)
		code = code[0 : len(code)-1]
	}
}

var n int
var nodes []Node
var huffTree *Node

func main() {
	fmt.Println("\n-- Huffman Codes --")

	// initialize min heap
	h := make(heap, 0)

	fmt.Print("\nEnter the number of characters: ")
	fmt.Scanf("%d", &n)
	nodes = make([]Node, 0)

	fmt.Println("Start entering character and frequency:")
	for i := 0; i < n; i++ {
		h = h.insertNodes()
	}

	huffTree = h.genHeap()
	huffTree.printHuffman()
}

// a, 5
// b, 9
// c, 12
// d, 13
// e, 16
// f, 45
