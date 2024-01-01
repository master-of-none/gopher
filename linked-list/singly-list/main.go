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

func main() {
	list := LinkedList{}

	list.traverseList()
	list.insertAtHead(30)
	list.insertAtHead(20)
	list.insertAtHead(10)
	list.traverseList()
}
