# Complete Tree Linked List (Go)



## Using

In the following we will create an empty **Complete Tree** and add some values to it.

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
