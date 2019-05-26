## Queue

In computer science, a queue is a collection in which the entities in the collection are kept in order and the principal (or only) operations on the collection are the addition of entities to the rear terminal position, known as enqueue, and removal of entities from the front terminal position, known as dequeue.

This makes the queue a First-In-First-Out (FIFO) data structure. In a FIFO data structure, the first element added to the queue will be the first one to be removed. This is equivalent to the requirement that once a new element is added, all elements that were added before have to be removed before the new element can be removed.

Often a peek or front operation is also entered, returning the value of the front element without dequeuing it. A queue is an example of a linear data structure, or more abstractly a sequential collection.

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

_References_:

- [https://www.geeksforgeeks.org/queue-data-structure/](https://www.geeksforgeeks.org/queue-data-structure/)
- [https://en.wikipedia.org/wiki/Queue\_(abstract_data_type)](<https://en.wikipedia.org/wiki/Queue_(abstract_data_type)>)
