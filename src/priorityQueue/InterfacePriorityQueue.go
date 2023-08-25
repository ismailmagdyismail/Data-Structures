package priorityQueue

type PriorityQueue interface {
	Size() int
	IsEmpty() bool
	Push(element any)
	Pop()
	Top() any
	GetElements() []any
}
