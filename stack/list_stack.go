package stack

import (
	"fmt"
	"strings"
)

type ListStack struct {
	head *ListStackNode
}

type ListStackNode struct {
	data interface{}
	next *ListStackNode
}

func NewListStack() *ListStack {
	return &ListStack{NewListStackNode(0)}
}

func NewListStackNode(v interface{}) *ListStackNode {
	return &ListStackNode{v, nil}
}

func (this *ListStack) IsEmpty() bool {
	if this.head.next == nil {
		return true
	}
	return false
}

func (this *ListStack) Push(v interface{}) {
	node := NewListStackNode(v)
	if this.head.next != nil {
		tem := this.head.next
		this.head.next = node
		node.next = tem
	} else {
		this.head.next = node
	}
}

func (this *ListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.head.next
	this.head.next = v.next
	return v.data
}

func (this *ListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.head.next
	return v.data
}

func (this *ListStack) Flush() {
	this.head.next = nil
}

func (this *ListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}
	formatStr := strings.Builder{}
	cur := this.head.next
	for cur != nil {
		formatStr.WriteString(fmt.Sprintf("%v\r\n", cur.data))
		cur = cur.next
	}
	fmt.Println(formatStr.String())
}
