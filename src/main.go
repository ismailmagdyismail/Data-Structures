package main

import (
	"GO/src/linkedList"
	"GO/src/stack"
	"fmt"
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
func testStack(stack stack.Stack) {

	fmt.Println("Is the stack empty?", stack.IsEmpty())

	stack.Push(42)
	stack.Push(73)
	stack.Push(15)

	fmt.Println("Stack size:", stack.Size())
	fmt.Println("Top element:", stack.Top())

	fmt.Println("Popping elements:")
	for !stack.IsEmpty() {
		fmt.Println("Popped:", stack.Top())
		stack.Pop()
	}

	fmt.Println("Is the stack empty?", stack.IsEmpty())
}
func main() {

	// Create a new doubly linked list
	//testLinkedList(linkedList.NewDoublyLinkedList())
	//testLinkedList(linkedList.NewSinglyLinkedList())
	//testStack(stack.NewStackDynamic())
	testStack(stack.NewStackLinked())
}
