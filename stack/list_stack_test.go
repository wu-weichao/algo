package stack

import "testing"

func TestListStack_Push(t *testing.T) {
	stack := NewListStack()
	stack.Print()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
}

func TestListStack_Pop(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
	t.Log(stack.Pop())
	stack.Pop()
	stack.Print()
}

func TestListStack_Top(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
	t.Log(stack.Top())
	stack.Print()
}
