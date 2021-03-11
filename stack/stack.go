package stack

type Stack struct {
	items []int
}

func (stack *Stack) Push(value int) {
	stack.items = append(stack.items, value)
}

func (stack *Stack) Pop() int {
	length := len(stack.items) - 1
	last := stack.items[length]
	stack.items = stack.items[:length]
	return last
}

func MakeStack() Stack {
	s := Stack{}

	return s
}
