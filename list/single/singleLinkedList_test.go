package single

import (
	"testing"
)

// 向头部插入
func TestLinkedList_InsertToHead(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToHead(i)
	}
	list.Print()
}

// 向尾部插入
func TestLinkedList_InsertToTail(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
}

// 查找结点
func TestLinkedList_FindByIndex(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	t.Log(list.FindByIndex(0))
	t.Log(list.FindByIndex(2))
	t.Log(list.FindByIndex(5))
	t.Log(list.FindByIndex(12))
}

// 删除结点
func TestLinkedList_DeleteNode(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	// 删除头结点
	list.DeleteNode(list.head.next)
	list.Print()
	// 删除指定结点
	node := list.FindByIndex(3)
	t.Log(node)
	if node != nil {
		list.DeleteNode(node)
		list.Print()
	}
}
