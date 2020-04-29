# Linked Lists

In computer science, a Linked list is a linear collection of data elements, whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence.

In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. This structure allows for efficient insertion or removal of elements from any position in the sequence during iteration. More complex variants add additional links, allowing more efficient insertion or removal of nodes at arbitrary positions.

A drawback of linked lists is that access time is linear (and difficult to pipeline). Faster access, such as random access, is not feasible.

### Types

1. Singly Linked List
2. Circular Linked List
3. Doubly Linked List
4. Circular Doubly Linked List

### Applications

1. Linked Lists can be used to implement Stacks , Queues.
2. Linked Lists can also be used to implement Graphs. (Adjacency list representation of Graph).
3. Implementing Hash Tables :- Each Bucket of the hash table can itself be a linked list. (Open chain hashing).
4. Linked lists are useful for dynamic memory allocation.
5. The real life application where the circular linked list is used is our Personal Computers, where multiple applications are running. All the running applications are kept in a circular linked list and the OS gives a fixed time slot to all for running. The Operating System keeps on iterating over the linked list until all the applications are completed.
6. Previous and next page in web browser – We can access previous and next url searched in web browser by pressing back and next button since, they are linked as linked list.
7. Circular Doubly Linked Lists are used for implementation of advanced data structures like Fibonacci Heap.

### Time Complexities

- Indexing: Θ(n)
- Insert/Delete at beginning: Θ(1)
- Insert/Delete at end:
  - Θ(1) when last element is known
  - Θ(n) when last element is unknown
- Insert/Delete in middle: search time + Θ(1)

### Operations

Following operations can be performed on each type of linked lists:

1. Insertion − Adds an element at the beginning of the list.
2. Deletion − Deletes an element at the beginning of the list.
3. Display − Displays the complete list.
4. Search − Searches an element using the given key.
5. Delete − Deletes an element using the given key.
6. Reverse - Reverses the linked list.

## Singly Linked List

Singly linked lists contain nodes which have a data field as well as 'next' field, which points to the next node in line of nodes. Operations that can be performed on singly linked lists include insertion, deletion and traversal.

## Cicular Linked List

In the last node of a list, the link field often contains a null reference, a special value is used to indicate the lack of further nodes. A less common convention is to make it point to the first node of the list; in that case, the list is said to be 'circular' or 'circularly linked'; otherwise, it is said to be 'open' or 'linear'. It is a list where the last pointer points to the first node.

## Doubly Linked List

In a 'doubly linked list', each node contains, besides the next-node link, a second link field pointing to the 'previous' node in the sequence. The two links may be called 'forward('s') and 'backwards', or 'next' and 'prev'('previous').

Many modern operating systems use doubly linked lists to maintain references to active processes, threads, and other dynamic objects.

## Circular Doubly Linked List

It is the combine version of circular and doubly linked list where each node points to previous and next node as well as the tail next reference(pointer) points to head and head previous reference(pointer) points to tails forming a circular doubly linked list.

## References

- [https://en.wikipedia.org/wiki/Linked_list](https://en.wikipedia.org/wiki/Linked_list)
- [https://www.quora.com/What-is-an-application-of-linear-linked-list-data-structures](https://www.quora.com/What-is-an-application-of-linear-linked-list-data-structures)
- [https://www.geeksforgeeks.org/applications-of-linked-list-data-structure/](https://www.geeksforgeeks.org/applications-of-linked-list-data-structure/)
- [https://www.tutorialspoint.com/data_structures_algorithms/linked_list_algorithms.htm](https://www.tutorialspoint.com/data_structures_algorithms/linked_list_algorithms.htm)
- [https://www.geeksforgeeks.org/circular-singly-linked-list-insertion/](https://www.geeksforgeeks.org/circular-singly-linked-list-insertion/)
- [https://www.geeksforgeeks.org/doubly-linked-list/](https://www.geeksforgeeks.org/doubly-linked-list/)
- [https://www.geeksforgeeks.org/doubly-circular-linked-list-set-1-introduction-and-insertion/](https://www.geeksforgeeks.org/doubly-circular-linked-list-set-1-introduction-and-insertion/)
