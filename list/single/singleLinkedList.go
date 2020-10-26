package single

import (
	"fmt"
	"strings"
)

// 单链表结点
type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}

func NewLinkedListNode(value interface{}) *LinkedListNode {
	return &LinkedListNode{value, nil}
}

func (this *LinkedListNode) GetValue() interface{} {
	return this.value
}

func (this *LinkedListNode) GetNext() *LinkedListNode {
	return this.next
}

// 单链表
type LinkedList struct {
	head *LinkedListNode // 哨兵结点
	len  uint            // 链表长度，哨兵结点不记录
}

func NewLinkedList() *LinkedList {
	// 创建单链表时，创建一个哨兵结点，统一头尾结点的操作
	return &LinkedList{NewLinkedListNode(0), 0}
}

// 在指定结点后追加结点
// 修改指定结点的后继指针为新结点，将新结点的后继指针修改为原结点的后继指针
func (this *LinkedList) InsertAfter(node *LinkedListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	// 记录node原来的下一个结点
	oldNext := node.next
	// 修改后继指针
	node.next = NewLinkedListNode(v)
	node.next.next = oldNext
	// 修改链表长度
	this.len++
	return true
}

// 在指定结点前添加结点
// 获取指点结点的前驱结点，修改此结点的后继指针为新结点，新结点的后继指针在指定结点
func (this *LinkedList) InsertBefore(node *LinkedListNode, v interface{}) bool {
	if node == nil || node == this.head {
		return false
	}
	// 查找前驱结点
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 尾结点的后驱指针指向nil，当为找到匹配的结点是cur为nil
	if cur == nil {
		return false
	}
	// 插入结点
	pre.next = NewLinkedListNode(v)
	pre.next.next = node
	// 修改链表长度
	this.len++
	return true
}

// 添加头结点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 添加尾结点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 删除结点
func (this *LinkedList) DeleteNode(node *LinkedListNode) bool {
	if node == nil {
		return false
	}
	// 找到前驱结点
	cur := this.head.next
	pre := this.head
	for cur != node {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 未找到结点
	if cur == nil {
		return false
	}
	// 删除结点操作
	pre.next = node.next
	node = nil
	this.len--
	return true
}

// 根据索引查找结点
func (this *LinkedList) FindByIndex(index uint) *LinkedListNode {
	// 查找的索引不能超出链表的长度
	if index > this.len {
		return nil
	}
	// 查找
	cur := this.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 打印链表信息
func (this *LinkedList) Print() {
	cur := this.head.next
	formatStr := strings.Builder{}
	formatStr.WriteString(fmt.Sprintf("len: %d，nodes: ", this.len))
	for cur != nil {
		formatStr.WriteString(fmt.Sprintf("%v", cur.GetValue()))
		cur = cur.next
		if cur != nil {
			formatStr.WriteString("->")
		}
	}
	fmt.Println(formatStr.String())
}
