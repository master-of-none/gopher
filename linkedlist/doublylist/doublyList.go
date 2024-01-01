package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtHead(val int) {
	newNode := &Node{value: val}

	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
}

func (list *LinkedList) insertAtTail(val int) {
	newNode := &Node{value: val}

	if list.head == nil {
		list.head = newNode
	} else {
		cur := list.head

		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newNode
		newNode.prev = cur
	}
}

func (list *LinkedList) traverseList() {
	cur := list.head

	if cur == nil {
		fmt.Println("Empty List")
		return
	}

	for cur != nil {
		fmt.Printf("%d -> ", cur.value)
		cur = cur.next
	}
	fmt.Println("-|")
}
