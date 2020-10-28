package queue

import (
	"fmt"
	"strings"
)

type CircularArrayQueue struct {
	data []interface{}
	n    int // 队列长度
	head int // 队头指针
	tail int // 队尾指针
	len  int // 队列实际数据长度
}

func NewCircularArrayQueue(n int) *CircularArrayQueue {
	return &CircularArrayQueue{make([]interface{}, n), n, 0, 0, 0}
}

func (this *CircularArrayQueue) Enqueue(v interface{}) bool {
	// 队满
	if this.len == this.n {
		return false
	}
	// 入队 追加到队尾
	this.data[this.tail] = v
	this.tail = (this.tail + 1) % this.n
	this.len++
	return true
}

func (this *CircularArrayQueue) Dequeue() interface{} {
	// 队空
	if this.len == 0 {
		return nil
	}
	// 出队
	v := this.data[this.head]
	this.head = (this.head + 1) % this.n
	this.len--
	return v
}

func (this *CircularArrayQueue) Print() {
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
