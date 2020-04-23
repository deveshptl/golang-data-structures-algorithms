#### Memory and Time requirements

The memory and time requirements to solve the input puzzle depends on the input. For example, consider the below input puzzle:

**Input**

```
1 2 3
0 4 6
7 5 8
```

**Required Output**

```
1 2 3
8 0 4
7 6 5
```

To transform the above input to required output takes subtantially long time and more memory, and may hang your computer, which happened to me 😜, if the algorithm continues to find the solution further.

#### Few methods to solve 8 Puzzle problem:

1. A<sup>\*</sup> Algorithm
2. BFS (Breadth First Search)
3. DFS (Depth First Search)
4. Best First Search
5. Iterative Deepening Search

#### Calculating Cost

The `calculateCost` function implemented in this implementation can be changed to calculate cost using different methods. Following are some ways to calculate the cost of any puzzle:

1. Number of mismatched tiles (Implemented)
2. Manhattan distance
3. Euclidean distance

#### References

Live Demo: [http://tristanpenman.com/demos/n-puzzle/](http://tristanpenman.com/demos/n-puzzle/)

List of sample inputs: [https://www.andrew.cmu.edu/course/15-121/labs/HW-7%20Slide%20Puzzle/lab.html](https://www.andrew.cmu.edu/course/15-121/labs/HW-7%20Slide%20Puzzle/lab.html)
