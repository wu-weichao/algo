package queue

import "testing"

func TestListQueue_Enqueue(t *testing.T) {
	queue := NewListQueue()
	queue.Print()
	queue.Enqueue(1)
	queue.Print()
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
}

func TestListQueue_Dequeue(t *testing.T) {
	queue := NewListQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
	queue.Dequeue()
	queue.Print()
	queue.Dequeue()
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	queue.Print()
}
