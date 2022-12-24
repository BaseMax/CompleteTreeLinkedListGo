/**
 *
 * @Name: Heap Tree Implementation in Go By Linked List
 * @Author: Max Base
 * @Repository: https://github.com/BaseMax/HeapTreeLinkedListGo
 * @License: GPL-3.0
 * @Date: 2022/12/24
 *
 **/

package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type HeapTree struct {
	root *Node
}

/**
 * build a heap tree
 */
func NewHeapTree() *HeapTree {
	return &HeapTree{}
}

/**
 * insert a new node to heap tree
 */
func (h *HeapTree) Insert(value int) {
	if h.root == nil {
		h.root = &Node{value: value}
		return
	}

	h.insert(h.root, value)
}

func (h *HeapTree) insert(n *Node, value int) {
	if n.left == nil {
		n.left = &Node{value: value}
		return
	}

	if n.right == nil {
		n.right = &Node{value: value}
		return
	}

	h.insert(n.left, value)
}

/**
 * print heap tree
 */
func (h *HeapTree) Print() {
	// print the heap tree by using ASCII art
	h.print(h.root, 0)
}
func (h *HeapTree) print(n *Node, level int) {
	if n == nil {
		return
	}

	h.print(n.right, level+1)

	for i := 0; i < level; i++ {
		print("    ")
	}

	print(n.value)
	print("  ")
	print(" ")
	print("\n")

	h.print(n.left, level+1)
}

func main() {
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
}
