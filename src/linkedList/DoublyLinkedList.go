package linkedList

import "fmt"

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (linkedList *doublyLinkedList) IsEmpty() bool {
	return linkedList.size == 0
}
func (linkedList *doublyLinkedList) Size() int {
	return linkedList.size
}
func (linkedList *doublyLinkedList) AddFront(value any) {
	defer linkedList.incrementSize()
	if linkedList.IsEmpty() {
		linkedList.head = &node{
			next:  nil,
			prev:  nil,
			value: value,
		}
		linkedList.tail = linkedList.head
		return
	}
	newNode := &node{
		prev:  nil,
		next:  linkedList.head,
		value: value,
	}
	linkedList.head.prev = newNode
	linkedList.head = newNode
}
func (linkedList *doublyLinkedList) AddBack(value any) {
	if linkedList.IsEmpty() {
		linkedList.AddFront(value)
		return
	}
	newNode := &node{
		prev:  linkedList.tail,
		next:  nil,
		value: value,
	}
	linkedList.tail.next = newNode
	linkedList.tail = newNode
	linkedList.size++
}

func (linkedList *doublyLinkedList) GetAt(index int) any {
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

func (linkedList *doublyLinkedList) SetAt(index int, value any) {
	requiredNode := linkedList.getNode(index)
	requiredNode.value = value
}
func (linkedList *doublyLinkedList) RemoveFront() {
	defer linkedList.decrementSize()
	if linkedList.IsEmpty() {
		panic("List is already Empty")
		return
	}
	if linkedList.Size() == 1 {
		linkedList.head = nil
		linkedList.tail = nil
		return
	}
	nextNode := linkedList.head.next
	nextNode.prev = nil
	linkedList.head.next = nil
	linkedList.head = nextNode
}

func (linkedList *doublyLinkedList) RemoveBack() {
	if linkedList.IsEmpty() {
		panic("List is already Empty")
		return
	}
	if linkedList.Size() == 1 {
		linkedList.RemoveFront()
		return
	}
	prevNode := linkedList.tail.prev
	prevNode.next = nil
	linkedList.tail.prev = nil
	linkedList.tail = prevNode
	linkedList.decrementSize()
}

func (linkedList *doublyLinkedList) Print() {
	tempNode := linkedList.head
	for tempNode != nil {
		fmt.Print(tempNode.value, " ")
		tempNode = tempNode.next
	}
	fmt.Println()
}

func (linkedList *doublyLinkedList) getNode(index int) *node {
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

func (linkedList *doublyLinkedList) incrementSize() {
	linkedList.size++
}
func (linkedList *doublyLinkedList) decrementSize() {
	linkedList.size--
}
