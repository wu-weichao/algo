package queue

import "testing"

func TestCircularArrayQueue_Enqueue(t *testing.T) {
	queue := NewCircularArrayQueue(5)
	queue.Print()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
	queue.Enqueue(4)
	t.Log(queue.Enqueue(5))
	t.Log(queue.Enqueue(6))
	queue.Print()
}

func TestCircularArrayQueue_Dequeue(t *testing.T) {
	queue := NewCircularArrayQueue(5)
	queue.Print()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	t.Log(queue.Enqueue(5))
	t.Log(queue.Enqueue(6))
	queue.Print()
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	queue.Print()
	t.Log(queue.Enqueue(7))
	t.Log(queue.Enqueue(8))
	t.Log(queue.Enqueue(9))
	queue.Print()
}
