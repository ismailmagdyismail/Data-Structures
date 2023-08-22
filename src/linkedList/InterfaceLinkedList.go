package linkedList

type LinkedList interface {
	AddFront(val any)
	AddBack(val any)
	RemoveFront()
	RemoveBack()
	Size() int
	IsEmpty() bool
	GetAt(index int) any
	SetAt(index int, val any)
	Print()
}
