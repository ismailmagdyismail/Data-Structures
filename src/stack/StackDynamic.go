package stack

type stackDynamic struct {
	size  int
	stack []any
}

func NewStackDynamic() Stack {
	return &stackDynamic{
		size:  0,
		stack: make([]any, 0),
	}
}

func (stack *stackDynamic) IsEmpty() bool {
	return stack.Size() == 0
}
func (stack *stackDynamic) Size() int {
	return stack.size
}
func (stack *stackDynamic) Top() any {
	if stack.IsEmpty() {
		panic("Stack is Empty , No elements Exist")
	}
	return stack.stack[stack.size-1]
}
func (stack *stackDynamic) Push(element any) {
	stack.stack = append(stack.stack, element)
	stack.size++
}
func (stack *stackDynamic) Pop() {
	if stack.IsEmpty() {
		panic("Stack is already Empty ,cannot POP")
	}
	stack.stack = stack.stack[:stack.size-1]
	stack.size--
}
