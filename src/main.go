package main

import (
	"GO/src/linkedList"
)

// Your doublyLinkedList implementation here...
func testLinkedList(list linkedList.LinkedList) {
	list.AddFront(3)
	list.AddFront(2)
	list.AddFront(1)
	list.Print() // Should print: 1 2 3

	// Test AddBack
	list.AddBack(4)
	list.AddBack(5)
	list.Print() // Should print: 1 2 3 4 5

	// Test RemoveFront
	list.RemoveFront()
	list.Print() // Should print: 2 3 4 5
}

func main() {

	// Create a new doubly linked list
	var list linkedList.LinkedList = linkedList.NewDoublyLinkedList()
	testLinkedList(list)
}
