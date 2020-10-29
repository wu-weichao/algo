package queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestBlockCircularArrayQueue_queue(t *testing.T) {
	queue := NewBlockCircularArrayQueue(5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			queue.Enqueue(fmt.Sprintf("enqueue-%d", i))
			fmt.Println(fmt.Sprintf("enqueue-%d", i))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			fmt.Println(fmt.Sprintf("dequeue-%v", queue.Dequeue()))
		}
	}()

	wg.Wait()
}
