package queue

type Queue interface {
	Push(element any)
	Pop()
	Size() int
	IsEmpty() bool
	Front() any
	Back() any
}
