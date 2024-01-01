package main

import "fmt"

func main() {
	list := LinkedList{}

	list.traverseList()
	list.insertAtHead(30)
	list.insertAtHead(20)
	list.insertAtHead(10)
	list.traverseList()
	list.insertAtTail(40)
	list.traverseList()
	temp1 := list.deleteAtHead()
	fmt.Printf("Deleted value at head is: %d\n", temp1)
	list.traverseList()
	temp2 := list.deleteAtTail()
	fmt.Printf("Deleted value at tail is: %d\n", temp2)
	list.traverseList()
}
