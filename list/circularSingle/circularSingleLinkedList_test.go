package circularSingle

import (
	"testing"
)

// 向头部插入
func TestCLinkedList_InsertToHead(t *testing.T) {
	list := NewCLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToHead(i)
	}
	list.Print()
	if list.FindByIndex(list.len-1).next != list.head {
		t.Fatalf("尾结点未指向哨兵结点")
	}
}

// 向尾部插入
func TestCLinkedList_InsertToTail(t *testing.T) {
	list := NewCLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	if list.FindByIndex(list.len-1).next != list.head {
		t.Fatalf("尾结点未指向哨兵结点")
	}
}

// 查找结点
func TestCLinkedList_FindByIndex(t *testing.T) {
	list := NewCLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	t.Log(list.FindByIndex(0))
	t.Log(list.FindByIndex(2))
	t.Log(list.FindByIndex(9))
	t.Log(list.FindByIndex(11))
}

// 删除结点
func TestCLinkedList_DeleteNode(t *testing.T) {
	list := NewCLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	t.Log(&list.head)
	// 删除头结点
	list.DeleteNode(list.head.next)
	list.Print()
	// 删除中间结点
	node := list.FindByIndex(3)
	list.DeleteNode(node)
	list.Print()
	// 删除指定尾结点
	node = list.FindByIndex(list.len - 1)
	list.DeleteNode(node)
	list.Print()
	if list.FindByIndex(list.len-1).next != list.head {
		t.Fatalf("尾结点未指向哨兵结点")
	}
}
