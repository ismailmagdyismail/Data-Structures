package linkedList

type singleLinkedList struct {
	head *node
	tail *node
	size int
}

func NewSingleList() singleLinkedList {
	return singleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}
