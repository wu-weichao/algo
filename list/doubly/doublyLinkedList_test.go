package doubly

import "testing"

func TestDoublyLinkedListNodeList_InsertToHead(t *testing.T) {
	list := NewDoublyLinkedListNodeList()
	for i := 0; i < 10; i++ {
		list.InsertToHead(i)
	}
	list.Print()
}

func TestDoublyLinkedListNodeList_InsertToTail(t *testing.T) {
	list := NewDoublyLinkedListNodeList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
}

func TestDoublyLinkedListNodeList_FindByIndex(t *testing.T) {
	list := NewDoublyLinkedListNodeList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()

	t.Log(list.FindByIndex(7))
	t.Log(list.FindByIndex(8))
	t.Log(list.FindByIndex(9))
	t.Log(list.FindByIndex(11))
}

func TestDoublyLinkedListNodeList_DeleteNode(t *testing.T) {
	list := NewDoublyLinkedListNodeList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	// 删除头结点
	list.DeleteNode(list.head.next)
	list.Print()
	// 删除中间结点
	list.DeleteNode(list.FindByIndex(3))
	list.Print()
	// 删除尾结点
	list.DeleteNode(list.FindByIndex(list.len - 1))
	list.Print()
}
