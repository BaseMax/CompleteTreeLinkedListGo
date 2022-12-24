/**
 *
 * @Name: Heap Tree in Go
 * @Author: Max Base
 * @Repository:
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
