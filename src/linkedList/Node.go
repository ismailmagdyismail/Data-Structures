package linkedList

type doublyNode struct {
	next  *doublyNode
	prev  *doublyNode
	value any
}

type singlyNode struct {
	next *singlyNode
	val  any
}
