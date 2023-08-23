package stack

import "GO/src/linkedList"

type stackLinked struct {
	size  int
	stack linkedList.LinkedList
}

func NewStackLinked() Stack {
	return &stackLinked{
		size:  0,
		stack: linkedList.NewSinglyLinkedList(),
	}
}
func (stack *stackLinked) Size() int {
	return stack.size
}
func (stack *stackLinked) IsEmpty() bool {
	return stack.Size() == 0
}
func (stack *stackLinked) Top() any {
	if stack.IsEmpty() {
		panic("Stack is Empty , No Top Elements")
	}
	return stack.stack.GetAt(0)
}
func (stack *stackLinked) Push(element any) {
	stack.stack.AddFront(element)
	stack.size++
}
func (stack *stackLinked) Pop() {
	if stack.IsEmpty() {
		panic("Stack is already empty,Cannot Pop")
	}
	stack.size--
	stack.stack.RemoveFront()
}
