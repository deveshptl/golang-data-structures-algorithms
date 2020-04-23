# Queue

In computer science, a queue is a collection in which the elements in the collection are kept in order and the principal operations on the collection are the addition of elements to the rear terminal position, known as enqueue, and removal of elements from the front terminal position, known as dequeue.

This makes the queue a First-In-First-Out (FIFO) data structure. In a FIFO data structure, the first element added to the queue will be the first one to be removed. This is equivalent to the requirement that once a new element is added, all elements that were added before have to be removed before the new element can be removed.

A queue is an example of a linear data structure, or more abstractly a sequential collection.

Queue can be implemented using an Array, Stack or Linked List. The easiest way of implementing a queue is by using an Array.

#### Types of Queues

1. Simple Queue
2. Circular Queue
3. Double Ended Queue
4. Circular Double Ended Queue
5. Priority Queue

#### Applications

1. Scheduling of resources and multiple processes running on same cpu.
2. When data is transferred asynchronously between two processes, Queue is used for synchronization.
3. Queue is a part of implementation in many well known algorithms like Breadth First Search traversal in a Graph.
4. Handling of interrupts in real-time systems. The interrupts are handled in the same order as they arrive, First come first served.
5. Breadth First Traversal in tree and graph data structure uses queue to store explored nodes.

## Simple Queue

Simple Queue can be implemented using Array, Stack or Linked List. Easiest way to implement simple queue is using array.

#### Operations

1. Insert using tail(REAR) pointer.
2. Delete using head(FRONT) pointer.

## Circular Queue

A circular buffer, circular queue, cyclic buffer or ring buffer is a data structure that uses a single, fixed-size buffer as if it were connected end-to-end. This structure lends itself easily to buffering data streams.

#### Operations

1. Insert using tail(REAR) pointer.
2. Delete using head(FRONT) pointer.

## Double Ended Queue

In computer science, a double-ended queue (abbreviated to deque) is an abstract data type that generalizes a queue, for which elements can be added to or removed from either the front (head) or back (tail). It is also often called a head-tail linked list, though properly this refers to a specific data structure implementation of a deque.

#### Operations

1. Insert using tail(REAR) pointer.
2. Insert using head(FRONT) pointer.
3. Delete using tail(REAR) pointer.
4. Delete using head(FRONT) pointer.

## Circular Double Ended Queue

Cicural Double Ended Queue is a generalized version of Queue data structure that allows insert and delete at both ends.

#### Operations

1. Insert using tail(REAR) pointer.
2. Insert using head(FRONT) pointer.
3. Delete using tail(REAR) pointer.
4. Delete using head(FRONT) pointer.

## Priority Queue

In computer science, a priority queue is an abstract data type which is like a regular queue or stack data structure, but where additionally each element has a "priority" associated with it. In a priority queue, an element with high priority is served before an element with low priority. In some implementations, if two elements have the same priority, they are served according to the order in which they were enqueued, while in other implementations, ordering of elements with the same priority is undefined.

#### Operations

1. Insert - To insert a new element.
2. Remove - To remove an element depending upon it's priority. Returns the element which is removed.
3. Display - Display's the content of queue.

#### Applications

- **Dijkstra’s Shortest Path Algorithm**: When the graph is stored in the form of adjacency list or matrix, priority queue can be used to extract minimum efficiently when implementing Dijkstra’s algorithm.

- **Prim’s algorithm**: It is used to implement Prim’s Algorithm to store keys of nodes and extract minimum key node at every step.

- **Data compression**: It is used in Huffman codes which is used to compresses data.

- **A\* Search Algorithm**: The A\* search algorithm finds the shortest path between two vertices of a weighted graph, trying out the most promising routes first. The priority queue (also known as the fringe) is used to keep track of unexplored routes, the one for which the total path length is smallest is given highest priority.

- **Heap Sort**: Heap sort is typically implemented using Heap which is an implementation of Priority Queue.

- **Operating systems**: It is also used in Operating System for load balancing (load balancing on server), interrupt handling.

#### Time Complexities

Time complexities of priority queue depends on which sorting algorithm it uses internally to sort elements based on their priority. Here I have used heap sort algorithm, hence the time complexities will be:

- Best Case: nlog(n)
- Average Case: nlog(n)
- Worst Case: nlog(n)

## References

- [https://www.geeksforgeeks.org/queue-data-structure/](https://www.geeksforgeeks.org/queue-data-structure/)
- [https://en.wikipedia.org/wiki/Queue\_(abstract_data_type)](<https://en.wikipedia.org/wiki/Queue_(abstract_data_type)>)
- [https://www.studytonight.com/data-structures/queue-data-structure](https://www.studytonight.com/data-structures/queue-data-structure)
- [https://www.geeksforgeeks.org/circular-queue-set-1-introduction-array-implementation/](https://www.geeksforgeeks.org/circular-queue-set-1-introduction-array-implementation/)
- [https://en.wikipedia.org/wiki/Circular_buffer](https://en.wikipedia.org/wiki/Circular_buffer)
- [https://www.thecrazyprogrammer.com/2014/02/c-program-for-various-operations-on-a-dequeue-represented-using-a-circular-array.html](https://www.thecrazyprogrammer.com/2014/02/c-program-for-various-operations-on-a-dequeue-represented-using-a-circular-array.html)
- [https://en.wikipedia.org/wiki/Double-ended_queue](https://en.wikipedia.org/wiki/Double-ended_queue)
- [https://www.geeksforgeeks.org/implementation-deque-using-circular-array/](https://www.geeksforgeeks.org/implementation-deque-using-circular-array/)
- [https://en.wikipedia.org/wiki/Priority_queue](https://en.wikipedia.org/wiki/Priority_queue)
