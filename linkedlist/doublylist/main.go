package main

func main() {
	list := LinkedList{}

	list.traverseList()
	list.insertAtHead(30)
	list.insertAtHead(20)
	list.insertAtHead(10)
	list.traverseList()
	list.insertAtTail(40)
	list.traverseList()
	list.deleteAtHead()
	list.traverseList()
	list.deleteAtTail()
	list.traverseList()
}
