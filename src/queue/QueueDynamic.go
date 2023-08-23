package queue

type queueDynamic struct {
	size  int
	queue []any
}

func NewQueueDynamic() Queue {
	return &queueDynamic{
		size:  0,
		queue: make([]any, 0),
	}
}

func (queue *queueDynamic) Size() int {
	return queue.size
}

func (queue *queueDynamic) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *queueDynamic) Push(element any) {
	queue.queue = append(queue.queue, element)
	queue.size++
}
func (queue *queueDynamic) Pop() {
	if queue.IsEmpty() {
		panic("Queue is already Empty ,Cannot Pop")
	}
	queue.queue = queue.queue[1:]
	queue.size--
}

func (queue *queueDynamic) Back() any {
	if queue.IsEmpty() {
		panic("Queue is Empty, No Back")
	}
	return queue.queue[queue.size-1]
}

func (queue *queueDynamic) Front() any {
	if queue.IsEmpty() {
		panic("Queue is Empty, No Front")
	}
	return queue.queue[0]
}
