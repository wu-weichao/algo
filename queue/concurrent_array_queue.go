package queue

import (
	"fmt"
	"strings"
	"sync"
)

type ConcurrentCircularArrayQueue struct {
	data   []interface{}
	n      int           // 队列长度
	head   int           // 队头指针
	tail   int           // 队尾指针
	len    int           // 队列实际数据长度
	lock   *sync.RWMutex // 锁
	enCond *sync.Cond
	deCond *sync.Cond
}

func NewConcurrentCircularArrayQueue(n int) *ConcurrentCircularArrayQueue {
	var lock sync.RWMutex
	return &ConcurrentCircularArrayQueue{
		data:   make([]interface{}, n),
		n:      n,
		lock:   &lock,
		enCond: sync.NewCond(&lock),
		deCond: sync.NewCond(lock.RLocker()),
	}
}

func (this *ConcurrentCircularArrayQueue) Enqueue(v interface{}) bool {
	this.lock.Lock()
	// 队满
	for this.len == this.n {
		// 等待
		this.enCond.Wait()
	}
	// 入队 追加到队尾
	this.data[this.tail] = v
	this.tail = (this.tail + 1) % this.n
	this.len++
	// 解锁
	this.lock.Unlock()
	// 通知阻塞的出队操作
	this.deCond.Broadcast()
	return true
}

func (this *ConcurrentCircularArrayQueue) Dequeue() interface{} {
	this.lock.RLock()
	// 队空
	for this.len == 0 {
		this.deCond.Wait()
	}
	// 出队
	v := this.data[this.head]
	this.head = (this.head + 1) % this.n
	this.len--
	// 解锁
	this.lock.RUnlock()
	// 通知阻塞的入队操作
	this.enCond.Broadcast()
	return v
}

func (this *ConcurrentCircularArrayQueue) Print() {
	if this.len == 0 {
		fmt.Println("queue is empty")
		return
	}
	formatStr := strings.Builder{}
	formatStr.WriteString("queue list: \r\n")
	for i := 0; i < this.len; i++ {
		if i > 0 {
			formatStr.WriteString("<-")
		}
		formatStr.WriteString(fmt.Sprintf("%v", this.data[(this.head+i)%this.n]))
	}
	fmt.Println(formatStr.String())
}
