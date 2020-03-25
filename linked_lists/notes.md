## Linked Lists

In computer science, a Linked list is a linear collection of data elements, whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence.

In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. This structure allows for efficient insertion or removal of elements from any position in the sequence during iteration. More complex variants add additional links, allowing more efficient insertion or removal of nodes at arbitrary positions.

A drawback of linked lists is that access time is linear (and difficult to pipeline). Faster access, such as random access, is not feasible.

#### Types

1. Singly Linked List
2. Circular Linked List
3. Doubly Linked List
4. Circular Doubly Linked List

#### Applications

1. Linked Lists can be used to implement Stacks , Queues.
2. Linked Lists can also be used to implement Graphs. (Adjacency list representation of Graph).
3. Implementing Hash Tables :- Each Bucket of the hash table can itself be a linked list. (Open chain hashing).
4. Linked lists are useful for dynamic memory allocation.
5. The real life application where the circular linked list is used is our Personal Computers, where multiple applications are running. All the running applications are kept in a circular linked list and the OS gives a fixed time slot to all for running. The Operating System keeps on iterating over the linked list until all the applications are completed.
6. Previous and next page in web browser – We can access previous and next url searched in web browser by pressing back and next button since, they are linked as linked list.
7. Circular Doubly Linked Lists are used for implementation of advanced data structures like Fibonacci Heap.

#### Time Complexities

- Indexing: Θ(n)
- Insert/Delete at beginning: Θ(1)
- Insert/Delete at end:
  - Θ(1) when last element is known
  - Θ(n) when last element is unknown
- Insert/Delete in middle: search time + Θ(1)

<br/>

_References_:

- [https://en.wikipedia.org/wiki/Linked_list](https://en.wikipedia.org/wiki/Linked_list)
- [https://www.quora.com/What-is-an-application-of-linear-linked-list-data-structures](https://www.quora.com/What-is-an-application-of-linear-linked-list-data-structures)
- [https://www.geeksforgeeks.org/applications-of-linked-list-data-structure/](https://www.geeksforgeeks.org/applications-of-linked-list-data-structure/)
