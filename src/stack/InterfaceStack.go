package stack

type Stack interface {
	Top() any
	Pop()
	Push(element any)
	IsEmpty() bool
	Size() int
}
