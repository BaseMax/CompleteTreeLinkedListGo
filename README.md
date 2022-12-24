# Heap Tree Go

This is a simple Go program and by using it you can easily insert your nodes to the heap tree without worring about it to not be a heap tree. This program will automatically fix the heap tree for you and you are sure that your heap tree is always a heap tree.

Heap Tree Go is a Go implementation of a heap tree. It is a binary tree where the parent node is always greater than or equal to its children. It is a complete binary tree, meaning that all levels of the tree are fully filled except for possibly the last level. The last level is filled from left to right.

## Using

In the following we will create an empty Heap tree and add some values to it.

Finally we will print the heap tree.

```go
// create a new heap tree
heapTree := NewHeapTree()

// insert some values
heapTree.Insert(1)
heapTree.Insert(2)
heapTree.Insert(3)
heapTree.Insert(4)
heapTree.Insert(5)
heapTree.Insert(6)
heapTree.Insert(7)
heapTree.Insert(8)
heapTree.Insert(9)
heapTree.Insert(10)

// print heap tree
heapTree.Print()
```

The output of above code will be:

```text
    3
1
        5
    2
            7
        4
                9
            6
                8
                    10
```

Copyright (c) 2022, Max Base
