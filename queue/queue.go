package queue

type Queue interface {
	Enqueue(v interface{}) bool
	Dequeue() interface{}
}
