# Graphs

In computer science, a graph is an abstract data type that is meant to implement the undirected graph and directed graph concepts from mathematics; specifically, the field of graph theory.

A graph `( G = <V, E> )` data structure consists of a finite (and possibly mutable) set of vertices (`V`) or nodes, together with a set of unordered pairs of these vertices for an undirected graph or a set of ordered pairs for a directed graph. These pairs are known as edges (`E`), or arcs.

**Note**: The number of edges `|E|` possible in an undirected graph with `|V|` vertices and no loops: `0 ≤ |E| ≤ |V|(|V| − 1)/2`. We have to divide product `|V|(|V| - 1)` by `2`, however, because it includes every edge twice.

#### Representation

Graph can be represented by using adjacency matrix, adjacency list or incidence matrix. Each implementations has their own pros and cons associated with it either it could be time complexity on operations or space complexity.

#### Types

1. Simple graph
2. Undirected or Directed graphs
3. Cyclic or Acyclic graphs
4. Labeled graphs
5. Weighted graphs or Unweighted
6. Infinite graphs
7. Graph having Connected or Disconnected components

#### Operations

1. **Add node `(G, x)`** - adds a node to the graph
2. **Add edge `(G, x, y)`** - adds an edge from node x to y
3. **Remove node `(G, x)`** - removes a node from the graph
4. **Remove edge `(G, x, y)`** - removes an edge from node x to y
5. **Get node value `(G, x)`** - returns the value of node x
6. **Set node value `(G, x, v)`** - sets the value of node x
7. **Get edge value `(G, x, y)`** - returns the value of an edge going from node x to y
8. **Set edge value `(G, x, y, v)`** - sets the value of an edge going from node x to y
9. **Breadth First Traversal** - traverses the graph in bfs order
10. **Depth First Traversal** - traverses the graph in dfs order

#### BFS Traversal Psuedo code

1. Initialize a queue with starting node in it.
2. Initialize an empty visited set.
3. Repeat following until queue is empty:
   - Remove a node from the queue and name it as `currNode`.
   - If the `currNode` is visited, start step-2 over again.
   - Mark the `currNode` as visited.
   - Expand the `currNode`, add resulting nodes to the queue if they aren't already visited and aren't in the stack.

#### DFS Traversal Psuedo code

1. Initialize a stack with starting node in it.
2. Initialize an empty visited set.
3. Repeat following until stack is empty:
   - Remove a node from the stack and name it as `currNode`.
   - If the `currNode` is visited, start step-2 over again.
   - Mark the `currNode` as visited.
   - Expand the `currNode`, add resulting nodes to the queue if they aren't already visited and aren't in the stack.

**Note**: The way you insert explored nodes into the queue or stack will result into different traversal paths.

#### Applications

1. **Shortest Path**: what is the best route between two cities? Google Maps.

2. **Facebook**: users are considered to be the vertices and if they are friends then there is an edge running between them. Facebook’s Friend suggestion algorithm uses graph theory.

3. **Traveling Salesman Problem (TSP)**: Finding the shortest tour through n cities that visits every city exactly once. In addition to obvious applications involving route planning, it arises in such modern applications as circuit board and VLSI chip fabrication, X-ray crystallography, and genetic engineering.

4. **Topological Sorting**: Dependencies of packages can be represented using graphs and those packages are compiled first which have no dependenies or have dependencies but these have already been compiled.

#### Terminologies

A graph with every pair of its vertices connected by an edge is called **complete**.

A graph with relatively few possible edges missing is called **dense**.

A graph with few edges relative to the number of its vertices is called **sparse**.

The node is attached as a child to the node it is being reached from with an edge called a **tree edge** or in other words it is an edge which is present in tree obtained after applying DFS on the graph.

If an edge leading to a previously visited node other than its immediate predecessor is encountered, the edge is noted as a **cross edge** or in others words it is a edge which connects two node such that they do not have any ancestor and a descendant relationship between them.

Edges connecting to its ancestor is called **back edges** or in other words it is an edge (u, v) such that v is ancestor of edge u but not part of DFS tree.

## References

- [https://en.wikipedia.org/wiki/Graph\_(abstract_data_type)](<https://en.wikipedia.org/wiki/Graph_(abstract_data_type)>)
- [https://www.geeksforgeeks.org/tree-back-edge-and-cross-edges-in-dfs-of-graph/](https://www.geeksforgeeks.org/tree-back-edge-and-cross-edges-in-dfs-of-graph/)
