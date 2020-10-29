package queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentCircularArrayQueue_queue(t *testing.T) {
	queue := NewConcurrentCircularArrayQueue(5)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			wg.Add(1)
			queue.Enqueue(fmt.Sprintf("enqueue-%d", i))
			fmt.Println(fmt.Sprintf("enqueue-%d", i))
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			wg.Add(1)
			fmt.Println(fmt.Sprintf("dequeue-%v", queue.Dequeue()))
		}()
	}
	wg.Wait()
}
