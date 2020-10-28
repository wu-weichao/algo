package queue

import (
	"fmt"
	"strings"
)

type ListQueue struct {
	head *ListQueueNode
	tail *ListQueueNode
}

func NewListQueue() *ListQueue {
	return &ListQueue{nil, nil}
}

type ListQueueNode struct {
	data interface{}
	next *ListQueueNode
}

func NewListQueueNode(v interface{}) *ListQueueNode {
	return &ListQueueNode{v, nil}
}

func (this *ListQueue) Enqueue(v interface{}) bool {
	node := NewListQueueNode(v)
	if this.head == nil {
		this.head = node
	}
	if this.tail != nil {
		this.tail.next = node
	}
	this.tail = node

	//this.tail.next
	//
	//
	//oldHead := this.head
	//if oldHead != nil {
	//	node.next = oldHead
	//}
	//this.head = node
	return true
}

func (this *ListQueue) Dequeue() interface{} {
	if this.head == nil {
		return nil
	}

	//if this.head.next == nil {
	//	return nil
	//}
	v := this.head.data
	this.head = this.head.next
	return v
}

func (this *ListQueue) Print() {
	if this.head == nil {
		fmt.Println("queue is empty")
		return
	}
	formatStr := strings.Builder{}
	formatStr.WriteString("queue list: \r\n")
	cur := this.head
	for cur != nil {
		formatStr.WriteString(fmt.Sprintf("%v", cur.data))
		cur = cur.next
		if cur != nil {
			formatStr.WriteString("->")
		}
	}
	fmt.Println(formatStr.String())
}
