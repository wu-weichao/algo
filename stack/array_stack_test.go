package stack

import "testing"

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack()
	if stack.IsEmpty() {
		t.Log("stack is empty")
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
	stack.Pop()
	stack.Print()
}

func TestArrayStack_Top(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
	stack.Top()
	stack.Print()
}

func TestArrayStack_Flush(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
	stack.Flush()
	stack.Print()
}
