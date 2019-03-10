## Priority Queue

In computer science, a priority queue is an abstract data type which is like a regular queue or stack data structure, but where additionally each element has a "priority" associated with it. In a priority queue, an element with high priority is served before an element with low priority. In some implementations, if two elements have the same priority, they are served according to the order in which they were enqueued, while in other implementations, ordering of elements with the same priority is undefined.


#### Operations

1. Insert - To insert a new element.
2. Remove - To remove an element depending upon it's priority. Returns the element which is removed.
3. Display - Display's the content of queue.


#### Applications

* **Dijkstra’s Shortest Path Algorithm**: When the graph is stored in the form of adjacency list or matrix, priority queue can be used to extract minimum efficiently when implementing Dijkstra’s algorithm.

* **Prim’s algorithm**: It is used to implement Prim’s Algorithm to store keys of nodes and extract minimum key node at every step.

* **Data compression**: It is used in Huffman codes which is used to compresses data.

* **Artificial Intelligence**: A* Search Algorithm : The A* search algorithm finds the shortest path between two vertices of a weighted graph, trying out the most promising routes first. The priority queue (also known as the fringe) is used to keep track of unexplored routes, the one for which a lower bound on the total path length is smallest is given highest priority.

* **Heap Sort**: Heap sort is typically implemented using Heap which is an implementation of Priority Queue.

* **Operating systems**: It is also use in Operating System for load balancing (load balancing on server), interrupt handling.


#### Time Complexities
Time complexities of priority queue depends on which sorting algorithm it uses internally to sort elements based on their priority. Here I have used heap sort algorithm, hence the time complexities will be:

* Best Case: nlog(n)
* Average Case: nlog(n)
* Worst Case: nlog(n)

*References*:
* [https://en.wikipedia.org/wiki/Priority_queue](https://en.wikipedia.org/wiki/Priority_queue)