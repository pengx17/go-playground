package minmaxstack

import (
	"github.com/golang-collections/collections/stack"
)

// MinmaxStack keep tracks of the min/max of the stack
type MinmaxStack struct {
	origin *stack.Stack
	mm     *stack.Stack
}

type Minmax struct {
	min int
	max int
}

func New() *MinmaxStack {
	return &MinmaxStack{stack.New(), stack.New()}
}

func (ms *MinmaxStack) Len() int {
	return ms.origin.Len()
}

func (ms *MinmaxStack) Peek() int {
	return ms.origin.Peek().(int)
}

func (ms *MinmaxStack) Push(val int) {
	ms.origin.Push(val)
	minmax, ok := ms.GetMinmax()

	if !ok {
		minmax = Minmax{val, val}
	}

	if minmax.max < val {
		minmax.max = val
	}
	if minmax.min > val {
		minmax.min = val
	}

	ms.mm.Push(minmax)
}

func (ms *MinmaxStack) Pop() int {
	ms.mm.Pop()
	return ms.origin.Pop().(int)
}

func (ms *MinmaxStack) GetMinmax() (Minmax, bool) {
	minmax, ok := ms.mm.Peek().(Minmax)
	return minmax, ok
}
