package priorityQueue

type priorityQueueDynamic struct {
	queue      []any
	comparator func(first, second any) bool
}

func NewPriorityQueue(comparator func(first, second any) bool) PriorityQueue {
	return &priorityQueueDynamic{
		queue:      make([]any, 0),
		comparator: comparator,
	}
}

func (pq *priorityQueueDynamic) Size() int {
	return len(pq.queue)
}
func (pq *priorityQueueDynamic) IsEmpty() bool {
	return pq.Size() == 0
}
func (pq *priorityQueueDynamic) Push(element any) {
	pq.queue = append(pq.queue, element)
	pq.trickleUp(pq.Size() - 1)
}
func (pq *priorityQueueDynamic) trickleUp(index int) {
	var parentIndex int = (index - 1) / 2
	if parentIndex < 0 {
		return
	}
	if pq.comparator(pq.queue[index], pq.queue[parentIndex]) {
		pq.queue[index], pq.queue[parentIndex] = pq.queue[parentIndex], pq.queue[index]
		pq.trickleUp(parentIndex)
		return
	}
}
func (pq *priorityQueueDynamic) Pop() {
	if pq.IsEmpty() {
		panic("Priority Queue Is already Empty")
	}
	pq.queue[0], pq.queue[pq.Size()-1] = pq.queue[pq.Size()-1], pq.queue[0]
	pq.queue = pq.queue[0 : pq.Size()-1]
	pq.trickleDown(0)
}
func (pq *priorityQueueDynamic) trickleDown(index int) {
	var leftChildIndex, rightChildIndex int = index*2 + 1, index*2 + 2
	if leftChildIndex >= pq.Size() {
		return
	}
	if !pq.comparator(pq.queue[index], pq.queue[leftChildIndex]) {
		pq.queue[index], pq.queue[leftChildIndex] = pq.queue[leftChildIndex], pq.queue[index]
		pq.trickleDown(leftChildIndex)
		return
	}
	if rightChildIndex >= pq.Size() {
		return
	}
	if !pq.comparator(pq.queue[index], pq.queue[rightChildIndex]) {
		pq.queue[index], pq.queue[rightChildIndex] = pq.queue[rightChildIndex], pq.queue[index]
		pq.trickleDown(rightChildIndex)
		return
	}
}
func (pq *priorityQueueDynamic) Top() any {
	if pq.IsEmpty() {
		panic("PriorityQueue Is empty ,NO TOP")
	}
	return pq.queue[0]
}
func (pq *priorityQueueDynamic) GetElements() []any {
	var elements []any = make([]any, pq.Size())
	copy(elements, pq.queue)
	return elements
}
