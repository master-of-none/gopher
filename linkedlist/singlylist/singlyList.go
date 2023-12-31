package main

import "fmt"

type Node struct {
	value int
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
	}
}

func (list *LinkedList) deleteAtHead() int {
	if list.head == nil {
		return 0
	}

	temp := list.head
	list.head = list.head.next
	return temp.value
}

func (list *LinkedList) deleteAtTail() int {
	if list.head == nil {
		return 0
	}

	cur := list.head
	prev := list.head

	if cur.next == nil {
		list.head = nil
		return cur.value
	}

	for cur.next != nil {
		prev = cur
		cur = cur.next
	}
	prev.next = nil
	return cur.value
}

func (list *LinkedList) traverseList() {
	cur := list.head

	if cur == nil {
		fmt.Println("Empty List")
		return
	}

	fmt.Println("The Linked List is: ")
	for cur != nil {
		fmt.Printf("%d -> ", cur.value)
		cur = cur.next
	}
	fmt.Println("-|")
}
