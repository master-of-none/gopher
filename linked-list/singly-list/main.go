package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertHead(val int) {
	newNode := &Node{value: val}

	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
}

func main() {
	list :=
}
