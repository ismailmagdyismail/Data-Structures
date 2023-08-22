package linkedList

import "fmt"

type singleLinkedList struct {
	head *singlyNode
	tail *singlyNode
	size int
}

func NewSinglyLinkedList() LinkedList {
	return &singleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}
func (linkedList *singleLinkedList) Size() int {
	return linkedList.size
}
func (linkedList *singleLinkedList) IsEmpty() bool {
	return linkedList.Size() == 0
}
func (linkedList *singleLinkedList) AddBack(val any) {
	defer linkedList.incrementSize()
	if linkedList.IsEmpty() {
		linkedList.head = &singlyNode{
			next:  nil,
			value: val,
		}
		linkedList.tail = linkedList.head
		return
	}
	linkedList.tail.next = &singlyNode{
		next:  nil,
		value: val,
	}
	linkedList.tail = linkedList.tail.next
}
func (linkedList *singleLinkedList) AddFront(val any) {
	if linkedList.IsEmpty() {
		linkedList.AddBack(val)
		return
	}
	newNode := &singlyNode{
		next:  linkedList.head,
		value: val,
	}
	linkedList.head = newNode
	linkedList.incrementSize()
}
func (linkedList *singleLinkedList) RemoveFront() {
	defer linkedList.decrementSize()
	if linkedList.IsEmpty() {
		panic("List is Already Empty")
		return
	}
	if linkedList.Size() == 1 {
		linkedList.head = nil
		linkedList.tail = nil
		return
	}
	nextNode := linkedList.head.next
	linkedList.head.next = nil
	linkedList.head = nextNode
}
func (linkedList *singleLinkedList) RemoveBack() {
	if linkedList.IsEmpty() || linkedList.Size() == 1 {
		linkedList.RemoveFront()
		return
	}
	beforeLastNode := linkedList.head
	for beforeLastNode.next != linkedList.tail {
		beforeLastNode = beforeLastNode.next
	}
	beforeLastNode.next = nil
	linkedList.tail = beforeLastNode
	linkedList.incrementSize()
}
func (linkedList *singleLinkedList) GetAt(index int) any {
	if index < 0 || index >= linkedList.size {
		panic("Invalid Index")
	}
	tempNode := linkedList.head
	i := 0
	for tempNode != nil && i < index {
		i++
		tempNode = tempNode.next
	}
	return tempNode.value
}
func (linkedList *singleLinkedList) SetAt(index int, value any) {
	requiredNode := linkedList.getNode(index)
	requiredNode.value = value
}
func (linkedList *singleLinkedList) Print() {
	tempNode := linkedList.head
	for tempNode != nil {
		fmt.Print(tempNode.value, " ")
		tempNode = tempNode.next
	}
	fmt.Println()
}
func (linkedList *singleLinkedList) getNode(index int) *singlyNode {
	if index < 0 || index >= linkedList.size {
		panic("Invalid Index")
	}
	tempNode := linkedList.head
	i := 0
	for tempNode != nil && i < index {
		i++
		tempNode = tempNode.next
	}
	return tempNode
}
func (linkedList *singleLinkedList) incrementSize() {
	linkedList.size++
}
func (linkedList *singleLinkedList) decrementSize() {
	linkedList.size--
}
