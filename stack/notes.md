## Stack

Stack is a linear data structure which follows a particular order in which the operations are performed. The order may be LIFO(Last In First Out) or FILO(First In Last Out).

#### Operations

1. **PUSH** - Adds an item in the stack and increments the `top` pointer. If the stack is full, then it is said to be an Overflow condition.
2. **POP** - Removes an item from the stack and decrements the `top`. The items are popped in the reversed order in which they are pushed. If the stack is empty, then it is said to be an Underflow condition.
3. **PEEK** - Returns top element of stack. Does not remove any elements from stack.
4. **isEmpty** - Returns true if stack is empty, else false.
5. **DISPLAY** - Displays content of stack.

#### Time Complexities

- push(), pop(), isEmpty() and peek() all take O(1) time. We do not run any loop in any of these operations.

#### Applications

1. Expression evaluation
2. Expression conversion
3. Syntax parsing
4. String reversal
5. Implementation of Tower of Hanoi Problem
6. Function call and Recursion
7. Parenthesis checking
8. Backtracking
9. File undo and redo operations

_References_:

- [https://www.geeksforgeeks.org/stack-data-structure/](https://www.geeksforgeeks.org/stack-data-structure/)
- [https://www.geeksforgeeks.org/stack-data-structure-introduction-program/](https://www.geeksforgeeks.org/stack-data-structure-introduction-program/)
- [https://www.thecrazyprogrammer.com/2016/04/applications-of-stack.html](https://www.thecrazyprogrammer.com/2016/04/applications-of-stack.html)
