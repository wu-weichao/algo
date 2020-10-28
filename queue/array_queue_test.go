package queue

import "testing"

func TestArrayQueue_Enqueue(t *testing.T) {
	queue := NewArrayQueue(5)
	queue.Print()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Print()
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	queue.Enqueue(6)
	t.Log(queue.Enqueue(7))
	t.Log(queue.Enqueue(8))
	queue.Print()
}

func TestArrayQueue_Dequeue(t *testing.T) {
	queue := NewArrayQueue(5)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
	t.Log(queue.Dequeue())
	queue.Enqueue(4)
	t.Log(queue.Dequeue())
	queue.Print()
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	queue.Print()
}
