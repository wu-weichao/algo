package circularSingle

import (
	"fmt"
	"strings"
)

// 循环单链表结点
type CLinkedListNode struct {
	value interface{}
	next  *CLinkedListNode
}

func NewCLinkedListNode(value interface{}) *CLinkedListNode {
	return &CLinkedListNode{value, nil}
}

func (this *CLinkedListNode) GetValue() interface{} {
	return this.value
}

func (this *CLinkedListNode) GetNext() *CLinkedListNode {
	return this.next
}

// 循环单链表
type CLinkedList struct {
	head *CLinkedListNode // 哨兵结点
	len  uint             // 链表长度，哨兵结点不记录
}

func NewCLinkedList() *CLinkedList {
	// 创建单链表时，创建一个哨兵结点且哨兵结点指向自己，统一头尾结点的操作
	headCLinkedListNode := NewCLinkedListNode(0)
	headCLinkedListNode.next = headCLinkedListNode
	return &CLinkedList{headCLinkedListNode, 0}
}

// 在指定结点后追加结点
// 修改指定结点的后继指针为新结点，将新结点的后继指针修改为原结点的后继指针
func (this *CLinkedList) InsertAfter(node *CLinkedListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	// 记录node原来的下一个结点
	oldNext := node.next
	// 修改后继指针
	node.next = NewCLinkedListNode(v)
	node.next.next = oldNext
	// 修改链表长度
	this.len++
	return true
}

// 在指定结点前添加结点
// 获取指点结点的前驱结点，修改此结点的后继指针为新结点，新结点的后继指针在指定结点
func (this *CLinkedList) InsertBefore(node *CLinkedListNode, v interface{}) bool {
	if node == nil {
		return false
	}
	// 查找前驱结点
	cur := this.head.next
	pre := this.head
	// 小于等于是为了将哨兵结点也计算进去
	for i := uint(0); i <= this.len; i++ {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 未找到指定的结点
	if cur != node {
		return false
	}
	// 插入结点
	pre.next = NewCLinkedListNode(v)
	pre.next.next = node
	// 修改链表长度
	this.len++
	return true
}

// 添加头结点
func (this *CLinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 添加尾结点
func (this *CLinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != this.head {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 删除结点
func (this *CLinkedList) DeleteNode(node *CLinkedListNode) bool {
	if node == nil || node == this.head {
		return false
	}
	// 找到前驱结点
	cur := this.head.next
	pre := this.head
	for i := uint(0); i < this.len; i++ {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 未找到指定的结点
	if cur != node {
		return false
	}
	// 删除结点操作
	pre.next = node.next
	node = nil
	this.len--
	return true
}

// 根据索引查找结点
func (this *CLinkedList) FindByIndex(index uint) *CLinkedListNode {
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

// 打印链表信息this.head
func (this *CLinkedList) Print() {
	cur := this.head.next
	formatStr := strings.Builder{}
	formatStr.WriteString(fmt.Sprintf("len: %d，nodes: ", this.len))
	for cur != this.head {
		formatStr.WriteString(fmt.Sprintf("%v", cur.GetValue()))
		cur = cur.next
		if cur != this.head {
			formatStr.WriteString("->")
		}
	}
	fmt.Println(formatStr.String())
}
