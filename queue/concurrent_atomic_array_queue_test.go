package queue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConcurrentAtomicCircularArrayQueue_queue(t *testing.T) {
	queue := NewConcurrentAtomicCircularArrayQueue(int32(20))

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()

		for i := 0; i < 30; i++ {
			fmt.Println("enqueue-", i, queue.Enqueue(i))
		}
	}()
	time.Sleep(time.Millisecond * 20)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 5; i++ {
				fmt.Println("dequeue-", i, queue.Dequeue())
			}
		}()
	}
	wg.Wait()
}
