package queue

import (
	"fmt"
	"strings"
	"sync/atomic"
)

type ConcurrentAtomicCircularArrayQueue struct {
	data []interface{}
	n    int32 // 队列长度
	head int32 // 队头指针
	tail int32 // 队尾指针
	len  int32 // 队列实际数据长度
}

func NewConcurrentAtomicCircularArrayQueue(n int32) *ConcurrentAtomicCircularArrayQueue {
	return &ConcurrentAtomicCircularArrayQueue{
		data: make([]interface{}, n),
		n:    n,
	}
}

func (this *ConcurrentAtomicCircularArrayQueue) Enqueue(v interface{}) bool {
	for {
		// 队满
		if this.len == this.n {
			// 等待
			return false
		}
		// 入队 追加到队尾
		curTail := atomic.LoadInt32(&this.tail)
		newTail := (curTail + 1) % this.n
		// 比较队尾指针是否变化
		if atomic.CompareAndSwapInt32(&this.tail, curTail, newTail) {
			this.data[newTail] = v
			this.len++
			break
		}
	}
	return true
}

func (this *ConcurrentAtomicCircularArrayQueue) Dequeue() interface{} {
	var v interface{}
	for {
		// 队空
		if this.len == 0 {
			return nil
		}
		// 出队
		curHead := atomic.LoadInt32(&this.head)
		newHead := (curHead + 1) % this.n
		if atomic.CompareAndSwapInt32(&this.head, curHead, newHead) {
			v = this.data[newHead]
			this.len--
			break
		}
	}
	return v
}

func (this *ConcurrentAtomicCircularArrayQueue) Print() {
	if this.len == 0 {
		fmt.Println("queue is empty")
		return
	}
	formatStr := strings.Builder{}
	formatStr.WriteString("queue list: \r\n")
	for i := int32(0); i < this.len; i++ {
		if i > 0 {
			formatStr.WriteString("<-")
		}
		formatStr.WriteString(fmt.Sprintf("%v", this.data[(this.head+i)%this.n]))
	}
	fmt.Println(formatStr.String())
}
