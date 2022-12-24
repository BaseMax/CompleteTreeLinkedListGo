# Complete Tree Linked List (Go)

This is a Go program for creating a complete tree using linked list. You can easily insert nodes to the tree and print it. This library provide two inserting functions. One of them is for inserting nodes to the tree by filling the tree from left to right in per row. The other one is for inserting nodes to the tree by filling the tree to the left side of the tree.

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

------------------

Also, there are another inserting function that it will insert the value to left and the right of node by filling the tree from left to right in per row.

```go
// Create another complete tree
CompleteTree2 := NewCompleteTree()

// Insert some values
CompleteTree2.InsertClean(10)
CompleteTree2.InsertClean(9)
CompleteTree2.InsertClean(8)
CompleteTree2.InsertClean(7)
CompleteTree2.InsertClean(6)
CompleteTree2.InsertClean(5)
CompleteTree2.InsertClean(4)
CompleteTree2.InsertClean(3)
CompleteTree2.InsertClean(2)
CompleteTree2.InsertClean(1)

// Print the tree
CompleteTree2.Print()
```

And the output will be:

```text
                2
                    1
            4
                3
                    1
        6
                2
                    1
            5
                3
                    1
    8
                2
                    1
            4
                3
                    1
        7
                2
                    1
            5
                3
                    1
10
                2
                    1
            4
                3
                    1
        6
                2
                    1
            5
                3
                    1
    9
                2
                    1
            4
                3
                    1
        7
                2
                    1
            5
                3
                    1
```

Copyright (c) 2022, Max Base
