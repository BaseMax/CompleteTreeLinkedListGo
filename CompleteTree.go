/**
 *
 * @Name: Complete Tree Implementation in Go By Linked List
 * @Author: Max Base
 * @Repository: https://github.com/BaseMax/CompleteTreeLinkedListGo
 * @License: GPL-3.0
 * @Date: 2022/12/24
 *
 **/

package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type CompleteTree struct {
	root *Node
}

/**
 * build a complete tree
 */
func NewCompleteTree() *CompleteTree {
	return &CompleteTree{}
}

/**
 * insert a new node to complete tree
 */
func (h *CompleteTree) Insert(value int) {
	if h.root == nil {
		h.root = &Node{value: value}
		return
	}

	h.insert(h.root, value)
}

func (h *CompleteTree) insert(n *Node, value int) {
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
 * insert a new node to complete tree by filling from left to right (per row)
 */
func (h *CompleteTree) InsertClean(value int) {
	if h.root == nil {
		h.root = &Node{value: value}
		return
	}

	h.insertClean(h.root, value)
}
func (h *CompleteTree) insertClean(n *Node, value int) {
	if n.left == nil {
		n.left = &Node{value: value}
		return
	}

	if n.right == nil {
		n.right = &Node{value: value}
		return
	}

	h.insertClean(n.left, value)
	if n.left != nil {
		h.insertClean(n.right, value)
	}
}

/**
 * print complete tree
 */
func (h *CompleteTree) Print() {
	// print the complete tree by using ASCII art
	h.print(h.root, 0)
}
func (h *CompleteTree) print(n *Node, level int) {
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
	// Create a new complete tree
	CompleteTree := NewCompleteTree()

	// Insert some values
	CompleteTree.Insert(1)
	CompleteTree.Insert(2)
	CompleteTree.Insert(3)
	CompleteTree.Insert(4)
	CompleteTree.Insert(5)
	CompleteTree.Insert(6)
	CompleteTree.Insert(7)
	CompleteTree.Insert(8)
	CompleteTree.Insert(9)
	CompleteTree.Insert(10)

	// Print complete tree
	CompleteTree.Print()

	fmt.Println("======================================")
	fmt.Println("======================================")

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
}
