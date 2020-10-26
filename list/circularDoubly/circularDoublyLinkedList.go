package circularDoubly

import (
	"fmt"
	"strings"
)

// 双向链表结点
type DoublyCLinkedListNode struct {
	value interface{}
	prev  *DoublyCLinkedListNode
	next  *DoublyCLinkedListNode
}

func NewDoublyCLinkedListNode(v interface{}) *DoublyCLinkedListNode {
	return &DoublyCLinkedListNode{v, nil, nil}
}

func (this *DoublyCLinkedListNode) GetValue() interface{} {
	return this.value
}

func (this *DoublyCLinkedListNode) GetPrev() *DoublyCLinkedListNode {
	return this.prev
}

func (this *DoublyCLinkedListNode) GetNext() *DoublyCLinkedListNode {
	return this.next
}

// 双向链表
type DoublyCLinkedListNodeList struct {
	head *DoublyCLinkedListNode
	len  uint
}

func NewDoublyCLinkedListNodeList() *DoublyCLinkedListNodeList {
	// 初始化哨兵结点，前驱结点和后继结点指向自己
	head := NewDoublyCLinkedListNode(0)
	head.next = head
	head.prev = head
	return &DoublyCLinkedListNodeList{head, 0}
}

// 在指定结点后插入
func (this *DoublyCLinkedListNodeList) InsertAfter(node *DoublyCLinkedListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	// 指定节点的原后继结点
	oldNext := node.next
	// 新结点
	newNode := NewDoublyCLinkedListNode(v)
	// 新结点前驱及后继结点
	newNode.prev = node
	newNode.next = oldNext
	// 修改指定结点的后继结点
	node.next = newNode
	// 修改原后继结点的前驱结点
	oldNext.prev = newNode
	this.len++
	return true
}

// 在指定结点前插入
func (this *DoublyCLinkedListNodeList) InsertBefore(node *DoublyCLinkedListNode, v interface{}) bool {
	if node == nil && node != this.head {
		return false
	}
	// 指定结点的前驱结点
	oldPrev := node.prev
	// 新结点
	newNode := NewDoublyCLinkedListNode(v)
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
func (this *DoublyCLinkedListNodeList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在尾部插入
func (this *DoublyCLinkedListNodeList) InsertToTail(v interface{}) bool {
	// 查询最后一个节点
	cur := this.head
	for cur.next != this.head {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 查找通过索引结点
func (this *DoublyCLinkedListNodeList) FindByIndex(index uint) *DoublyCLinkedListNode {
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
func (this *DoublyCLinkedListNodeList) DeleteNode(node *DoublyCLinkedListNode) bool {
	if node == nil {
		return false
	}
	// 待删除结点的前驱结点的后继结点调整为待删除结点的后继结点
	node.prev.next = node.next
	// 待删除结点的后继结点的前驱结点调整为待删除结点的前驱结点
	node.next.prev = node.prev
	// 删除结点
	node = nil
	this.len--
	return true
}

// 打印链表
func (this *DoublyCLinkedListNodeList) Print() {
	formatStr := strings.Builder{}
	formatStr.WriteString(fmt.Sprintf("len: %d，values: ", this.len))
	cur := this.head.next
	for i := uint(0); i < this.len; i++ {
		formatStr.WriteString(fmt.Sprintf("%v", cur.GetValue()))
		if cur.GetNext() != this.head {
			formatStr.WriteString("->")
		}
		cur = cur.GetNext()
	}
	fmt.Println(formatStr.String())
}
