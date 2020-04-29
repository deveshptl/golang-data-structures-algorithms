# Heap

Heap is a tree based data structure where the tree is a complete binary tree. This binary tree satisfies the properties of heap:

1. For max heap, any given node C, if P is a parent node of C, then the key (the value) of P is greater than or equal to the key of C.
2. For min heap, any given node C, if P is a parent node of C, then the key (the value) of P is less than or equal to the key of C.

### Representation

Heap can be implemented using binary tree or array.

### Types

1. Min Binary Heap: Returns the least value node present in the heap.
2. Max Binary Heaps: Returns the max value node present in the heap.

### Operations

1. **Insert**: Inserts a node in the heap
2. **ExtractMin/ExtractMax**: Extracts the min or max value node depending on the type of the heap
3. **Delete**: Deletes a node from the heap
4. **Increase/Decrease Key**: Increases or decreases the key value of given node

Main functions of heap are:

1. **MaxHeapify/MinHeapify (aka Bubble Up)**: This is called after inserting an element, which is inserted in the end of the heap. This function traverses the heap in upward direction while making necessary shifts between nodes.
2. **Bubble down**: This is called after deleting an element, which is the root of the heap and the last node of the heap is placed there instead. This function traverses downward in direction while making necessary shifts between nodes.

### Applications

1. Heap Sort: Used to sort the elements in an array
2. Priority Queue: Used to select the highest priority job
3. Graph Algorithms: Used in graph algorithms like Dijkstra’s Shortest Path and Prim’s Minimum Spanning Tree

## References

- [https://en.wikipedia.org/wiki/Heap\_(data_structure)](<https://en.wikipedia.org/wiki/Heap_(data_structure)>)
- [https://en.wikipedia.org/wiki/Binary_heap](https://en.wikipedia.org/wiki/Binary_heap)
- [https://www.geeksforgeeks.org/binary-heap/](https://www.geeksforgeeks.org/binary-heap/)
