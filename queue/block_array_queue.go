package queue

import (
	"fmt"
	"strings"
	"sync"
)

type BlockCircularArrayQueue struct {
	data []interface{}
	n    int         // 队列长度
	head int         // 队头指针
	tail int         // 队尾指针
	len  int         // 队列实际数据长度
	lock *sync.Mutex // 锁
}

func NewBlockCircularArrayQueue(n int) *BlockCircularArrayQueue {
	var lock sync.Mutex
	return &BlockCircularArrayQueue{
		data: make([]interface{}, n),
		n:    n,
		lock: &lock,
	}
}

func (this *BlockCircularArrayQueue) Enqueue(v interface{}) bool {
	this.lock.Lock()
	// 队满
	for this.len == this.n {
		// 等待
		this.lock.Unlock()
		this.lock.Lock()
	}
	// 入队 追加到队尾
	this.data[this.tail] = v
	this.tail = (this.tail + 1) % this.n
	this.len++
	// 解锁
	this.lock.Unlock()
	return true
}

func (this *BlockCircularArrayQueue) Dequeue() interface{} {
	this.lock.Lock()
	// 队空
	for this.len == 0 {
		this.lock.Unlock()
		this.lock.Lock()
	}
	// 出队
	v := this.data[this.head]
	this.head = (this.head + 1) % this.n
	this.len--
	// 解锁
	this.lock.Unlock()
	return v
}

func (this *BlockCircularArrayQueue) Print() {
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
