package queue

import (
	"fmt"
	"strings"
)

type ArrayQueue struct {
	data []interface{} // 队列
	n    int           // 队列长度
	head int           // 队头指针
	tail int           // 队尾指针
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (this *ArrayQueue) Enqueue(v interface{}) bool {
	// 队满
	if this.tail == this.n {
		// 队满处理
		if this.head == 0 {
			return false
		}
		// 存在空闲的空间 数据前移
		// 向前移动head个位置
		for i := this.head; i < this.tail; i++ {
			this.data[i-this.head] = this.data[i]
		}
		// 调整队头和队尾指针
		this.tail -= this.head
		this.head = 0
	}
	// 入队 追加到队尾
	this.data[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) Dequeue() interface{} {
	// 队空
	if this.tail == this.head {
		return nil
	}
	// 出队
	v := this.data[this.head]
	this.head++
	return v
}

func (this *ArrayQueue) Print() {
	if this.tail == this.head {
		fmt.Println("queue is empty")
		return
	}
	formatStr := strings.Builder{}
	formatStr.WriteString("queue list: \r\n")
	for i := this.head; i < this.tail; i++ {
		if i != this.head {
			formatStr.WriteString("<-")
		}
		formatStr.WriteString(fmt.Sprintf("%v", this.data[i]))
	}
	fmt.Println(formatStr.String())
}
