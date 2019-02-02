## Data Structures and Algorithms

### Instructions

1. To run any of the file make sure that you have go (golang) programming language installed on your machine. Follow how to install instructions from [golang download instructions](https://golang.org/doc/install).
2. Once golang is installed on your machine, just clone or download this repository.
3. Now `cd <folder-name>` into the folder where the file you want to run is located.
4. Now run `go run <file-name>`.

**Note**: If a folder contains multiple `.go` extension files then use `go run <file-name> [<file-name>...]`. For e.g `bst_using_arr` folder contains two files: `bst_using_arr.go` and `traversal.go`. So use the command `go run bst_using_arr.go traversal.go`.

### Description

This repo is for learning data structures and algorithms which is the base of any programming language.

**FOLDER NAMES**

01. **algorithms** -
    * *01knapsack_dp* - 0-1 Knapsack Problem using Dynamic Programming
    * *articulation_point_detection* - Detecting Articulation Points in an undirected graph
    * *assembly_line_scheduling* - Assembly Line Scheduling algorithm using Dynamic Programming
    * *bellman_ford* - Bellman Ford Algorithm
    * *bridge_detection* - Bridge Detection/Cut Edge Detection in an undirected graph
    * *cd_directed_graph_traversals* - Cycle Detection in Directed Graphs using Traversals techniques &nbsp;&nbsp;:point_left:
    * *cd_undirected_graph_traversals* - Cycle Detection in Undirected Graphs using Traversals techniques &nbsp;&nbsp;:point_left:
    * *dijkstra* - Dijkstra's Algorithm
    * *expression_conversions* - 
        * *infix_postfix* - Infix to Postfix Conversion
        * *infix_prefix* - Infix to Prefix Conversion
        * *postfix_infix* - Postfix to Infix Conversion
        * *postfix_prefix* - Postfix to Prefix Conversion
        * *prefix_infix* - Prefix to Infix Conversion
        * *prefix_postfix* - Prefix to Postfix Conversion
    * *floyd_warshall* - Floyd Warshall Algorithm (All points shortest path)
    * *kruskals* - Kruskal's Algorithm (Finding Minimum Spanning Tree MST using Greedy Approach)
    * *making_change_dp* - Making Change Problem using Dynamic Programming
    * *naive_pattern_matching* - Naive Pattern/String Matching
    * *prims* - Prim's Algorithm (Finding Minimum Spanning Tree MST using Greedy Approach)
    * *searching* -
        * *binary_search* - Binary Search - O(log n)
        * *linear_search* - Linear Search - O(n)
    * *sorting* - 
        * *bubble_sort* - Bubble Sort - O(n<sup>2</sup>)
        * *counting_sort* - Counting Sort - O(n + k)
        * *heap_sort* - Heap Sort - O(nlog(n)
        * *insertion_sort* - Insertion Sort - O(n<sup>2</sup>)
        * *merge_sort* - Merge Sort - O(nlog(n))
        * *quick_sort* - Quick Sort - Θ(nlog(n))
        * *radix_sort* - Radix Sort - O(n+k)
        * *selection_sort* - Selection Sort - (O(n<sup>2</sup>))
        * *shell_sort* - Shell Sort - О(n)
    * *toh* - Tower of Hanoi
    * *topological_sort* - Topological Sort
    * *union_find* - Union Find / Disjoint Sets (Detecting cycles in an undirected graph)
02. **graphs** -
    * *directed_unweighted* - Directed Unweighted graph
    * *directed_weighted* - Directed Weighted graph
    * *undirected_unweighted* - Undirected Unweighted graph
    * *undirected_weighted* - Undirected Weighted graph
03. **heaps** -
    * *max_binary_heap* - Max Binary Heap
    * *min_binary_heap* - Min Binary Heap
04. **linked_lists** -
    * *circular_doubly_ll* - Circular Doubly Linked List
    * *circular_ll* - Circular Linked List
    * *doubly_ll* - Doubly Linked List
    * *pres_rev_single_ll* - Preserve order during insertion on Single Linked List and Reversing Single Linked List
    * *single_ll* - Single Linked List
05. **queues** - 
    * *cdqueue* - Circular Double ended Queue
    * *cqueue* - Circular Queue
    * *dqueue* - Double ended Queue
    * *priority_queue* - Priority Queue with the use of Min Heap
    * *simple_queue* - Simple Queue
06. **stack** - stack
07. **trees** - 
    * *avl_tree_using_ll* - AVL Tree using linked list with BFS and DFS (Pre, In, Post) order traversals.
    * *bst_using_arr* - Binary Search Tree using array with BFS and DFS (Pre, In, Post) order traversals.
    * *bst_using_ll* - Binary Search Tree using linked list with BFS and DFS (Pre, In, Post) order traversals.
    * *simple_bt_using_arr* - Simple Binary Tree using array with BFS and DFS (Pre, In, Post) order traversals.
    * *simple_bt_using_ll* - Simple Binary Tree using linked list with BFS and DFS (Pre, In, Post) order traversals.


**Note**: The pointer &nbsp;"&nbsp;:point_left:&nbsp;"&nbsp; indicates incomplete implementation and is in todo list.

### Contribution

Feel free to make PR if you find a better way to solve/implement any algorithm or data structure. I will be more than happy to merge your PR. :smile:

### License
This repository is released under the [MIT license](https://opensource.org/licenses/MIT). In short, this means you are free to use this software in any personal, open-source or commercial projects. Attribution is optional but appreciated.

```
HAPPY CODING
HAPPY LEARNING
```