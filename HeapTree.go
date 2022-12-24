package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type HeapTree struct {
	root *Node
}
