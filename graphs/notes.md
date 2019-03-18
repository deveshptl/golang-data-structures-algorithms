### Graphs

In computer science, a graph is an abstract data type that is meant to implement the undirected graph and directed graph concepts from mathematics; specifically, the field of graph theory.

A graph `( G = <V, E> )` data structure consists of a finite (and possibly mutable) set of vertices (`V`) or nodes or points, together with a set of unordered pairs of these vertices for an undirected graph or a set of ordered pairs for a directed graph. These pairs are known as edges (`E`), arcs, or lines for an undirected graph and as arrows, directed edges, directed arcs, or directed lines for a directed graph.

**Note**: The number of edges `|E|` possible in an undirected graph with `|V|` vertices and no loops: `0 ≤ |E| ≤ |V|(|V| − 1)/2`. We have to divide product `|V|(|V| - 1)` by `2`, however, because it includes every edge twice. 

A graph with every pair of its vertices connected by an edge is called **complete**.

A graph with relatively few possible edges missing is called **dense**.

A graph with few edges relative to the number of its vertices is called **sparse**.

**Graph Representation**:
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

* **adjacent `(G, x, y)`**: Tests whether there is an edge from the vertex x to the vertex y.
* **neighbors `(G, x, y)`**: Lists all vertices y such that there is an edge from the vertex x to the vertex y.
* **add_vertex `(G, x)`**: Adds the vertex x, if it is not there.
* **remove_vertex `(G, x)`**: Removes the vertex x, if it is there.
* **add_edge `(G, x, y)`**: Adds the edge from the vertex x to the vertex y, if it is not there.
* **remove_edge `(G, x, y)`**: Removes the edge from the vertex x to the vertex y, if it is there.
* **get_vertex_value `(G, x)`**: Returns the value associated with the vertex x.
* **set_vertex_value `(G, x, v)`**: Sets the value associated with the vertex x to v.

Structures that associate values to the edges usually also provide:
* **get_edge_value `(G, x, y)`**: Returns the value associated with the edge (x, y).
* **set_edge_value `(G, x, y, v)`**: Sets the value associated with the edge (x, y) to v.


#### Applications

1. **Shortesh Path**: what is the best route between two cities? Google Maps.

2. **Facebook**: users are considered to be the vertices and if they are friends then there is an edge running between them. Facebook’s Friend suggestion algorithm uses graph theory.

3. **Traveling Salesman Problem (TSP)**: Finding the shortest tour through n cities that visits every city exactly once. In addition to obvious applications involving route planning, it arises in such modern applications as circuit board and VLSI chip fabrication, X-ray crystallography, and genetic engineering.

4. **Topological Sorting**: Dependencies of packages can be represented using graphs and those packages are compiled first which have no dependenies or have dependencies but these have already been compiled.

<br>
*References*:
* [https://en.wikipedia.org/wiki/Graph_(abstract_data_type)](https://en.wikipedia.org/wiki/Graph_(abstract_data_type))