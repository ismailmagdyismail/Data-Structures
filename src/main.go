package main

import (
	"GO/src/linkedList"
	"GO/src/priorityQueue"
	"GO/src/queue"
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
func testQueue(queue queue.Queue) {
	// Test normal use cases
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	fmt.Println("Front:", queue.Front()) // Should print 1
	fmt.Println("Back:", queue.Back())   // Should print 3

	queue.Pop()
	fmt.Println("Front after Pop:", queue.Front()) // Should print 2

	// Test edge cases
	fmt.Println("Is Empty:", queue.IsEmpty()) // Should print false

	queue.Pop()
	queue.Pop()
	fmt.Println("Is Empty after Pops:", queue.IsEmpty()) // Should print true

	// Test empty queue behavior
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Caught panic:", r) // Should print "Caught panic: Queue is Empty, No Front"
			}
		}()
		queue.Front()
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Caught panic:", r) // Should print "Caught panic: Queue is already Empty, Cannot Pop"
			}
		}()
		queue.Pop()
	}()
}
func testPriorityQueue(pq priorityQueue.PriorityQueue) {
	// Test normal use cases
	pq.Push(5)
	pq.Push(10)
	pq.Push(3)
	pq.Push(8)

	fmt.Println("Top:", pq.Top())              // Should print 10
	fmt.Println("Size:", pq.Size())            // Should print 4
	fmt.Println("Elements:", pq.GetElements()) // Should print [10 8 3 5]

	pq.Pop()
	fmt.Println("Top:", pq.Top())              // Should print 8
	fmt.Println("Size:", pq.Size())            // Should print 3
	fmt.Println("Elements:", pq.GetElements()) // Should print [8 5 3]

	// Test edge cases
	emptyHeap := priorityQueue.NewPriorityQueue(func(first, second interface{}) bool {
		return first.(int) > second.(int)
	})

	fmt.Println("Is Empty:", emptyHeap.IsEmpty()) // Should print true

	// Attempting to get the top of an empty heap should panic
	// Uncomment the next line to test the panic
	// fmt.Println("Top:", emptyHeap.Top())
	// Push elements and then pop until empty
	emptyHeap.Push(1)
	emptyHeap.Push(2)
	emptyHeap.Push(3)
	emptyHeap.Pop()
	emptyHeap.Pop()
	emptyHeap.Pop()
	fmt.Println("Is Empty:", emptyHeap.IsEmpty()) // Should print true
}
func main() {

	// Create a new doubly linked list
	//testLinkedList(linkedList.NewDoublyLinkedList())
	//testLinkedList(linkedList.NewSinglyLinkedList())
	//testStack(stack.NewStackDynamic())
	//testStack(stack.NewStackLinked())
	//testQueue(queue.NewQueueDynamic())
	testPriorityQueue(priorityQueue.NewPriorityQueue(func(first, second interface{}) bool {
		return first.(int) > second.(int)
	}))
}
