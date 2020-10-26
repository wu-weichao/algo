package doubly

import (
	"fmt"
	"strings"
)

// 双向链表结点
type DoublyLinkedListNode struct {
	value interface{}
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

func NewDoublyLinkedListNode(v interface{}) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{v, nil, nil}
}

func (this *DoublyLinkedListNode) GetValue() interface{} {
	return this.value
}

func (this *DoublyLinkedListNode) GetPrev() *DoublyLinkedListNode {
	return this.prev
}

func (this *DoublyLinkedListNode) GetNext() *DoublyLinkedListNode {
	return this.next
}

// 双向链表
type DoublyLinkedListNodeList struct {
	head *DoublyLinkedListNode
	len  uint
}

func NewDoublyLinkedListNodeList() *DoublyLinkedListNodeList {
	return &DoublyLinkedListNodeList{NewDoublyLinkedListNode(0), 0}
}

// 在指定结点后插入
func (this *DoublyLinkedListNodeList) InsertAfter(node *DoublyLinkedListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	// 指定节点的原后继结点
	oldNext := node.next
	// 新结点
	newNode := NewDoublyLinkedListNode(v)
	// 新结点前驱及后继结点
	newNode.prev = node
	newNode.next = oldNext
	// 修改指定结点的后继结点
	node.next = newNode
	// 修改原后继结点的前驱结点
	if oldNext != nil {
		oldNext.prev = newNode
	}
	this.len++
	return true
}

// 在指定结点前插入
func (this *DoublyLinkedListNodeList) InsertBefore(node *DoublyLinkedListNode, v interface{}) bool {
	if node == nil && node != this.head {
		return false
	}
	// 指定结点的前驱结点
	oldPrev := node.prev
	// 新结点
	newNode := NewDoublyLinkedListNode(v)
	newNode.prev = oldPrev
	newNode.next = node
	// 修改指点结点前驱结点的后继结点
	oldPrev.next = newNode
	// 修改指定的前驱结点
	node.prev = newNode
	this.len++
	return true
}

// 在头部插入
func (this *DoublyLinkedListNodeList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在尾部插入
func (this *DoublyLinkedListNodeList) InsertToTail(v interface{}) bool {
	// 查询最后一个节点
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 查找通过索引结点
func (this *DoublyLinkedListNodeList) FindByIndex(index uint) *DoublyLinkedListNode {
	if index > this.len {
		return nil
	}
	cur := this.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除结点
func (this *DoublyLinkedListNodeList) DeleteNode(node *DoublyLinkedListNode) bool {
	if node == nil {
		return false
	}
	// 待删除结点的前驱结点的后继结点调整为待删除结点的后继结点
	node.prev.next = node.next
	// 待删除结点的后继结点的前驱结点调整为待删除结点的前驱结点
	if node.next != nil {
		node.next.prev = node.prev
	}
	// 删除结点
	node = nil
	this.len--
	return true
}

// 打印链表
func (this *DoublyLinkedListNodeList) Print() {
	formatStr := strings.Builder{}
	formatStr.WriteString(fmt.Sprintf("len: %d，values: ", this.len))
	cur := this.head.next
	for i := uint(0); i < this.len; i++ {
		formatStr.WriteString(fmt.Sprintf("%v", cur.GetValue()))
		if cur.GetNext() != nil {
			formatStr.WriteString("->")
		}
		cur = cur.GetNext()
	}
	fmt.Println(formatStr.String())
}
