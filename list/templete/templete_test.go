package templete

import (
	"demo/list/single"
	"testing"
)

func TestGetListMiddleNode(t *testing.T) {
	list := single.NewLinkedList()
	// 1
	list.InsertToTail(1)
	list.Print()
	node := GetListMiddleNode(list)
	t.Log(node)
	// 1 2
	list.InsertToTail(2)
	list.Print()
	node = GetListMiddleNode(list)
	t.Log(node)
	// 1 2 3
	list.InsertToTail(3)
	list.Print()
	node = GetListMiddleNode(list)
	t.Log(node)
	// 1 2 3 4
	list.InsertToTail(4)
	list.Print()
	node = GetListMiddleNode(list)
	t.Log(node)
}

func TestListReverse(t *testing.T) {
	list := single.NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	list = ListReverse(list)
	list.Print()
}

func TestListHasCycle(t *testing.T) {
	list := single.NewLinkedList()
	list.InsertToTail(1)
	list.InsertToTail(2)
	list.InsertToTail(3)
	list.InsertToTail(4)
	if ListHasCycle(list) {
		t.Log("链表存在环")
	} else {
		t.Log("链表不存在环")
	}
	list.FindByIndex(3).SetNext(list.FindByIndex(1))
	if ListHasCycle(list) {
		t.Log("链表存在环")
	} else {
		t.Log("链表不存在环")
	}
}

func TestListIsPalindrome(t *testing.T) {
	list := single.NewLinkedList()
	list.InsertToTail('i')
	list.InsertToTail('m')
	list.InsertToTail('w')
	list.InsertToTail('w')
	list.InsertToTail('c')
	list.InsertToTail('c')
	list.InsertToTail('w')
	list.InsertToTail('w')
	list.InsertToTail('m')
	list.InsertToTail('i')
	list.Print()
	if ListIsPalindrome(list) {
		t.Log("是回文串")
	} else {
		t.Log("不是回文串")
	}
	list.Print()
}

func TestSortListMerge(t *testing.T) {
	listOne := single.NewLinkedList()
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			listOne.InsertToTail(i)
		}
	}
	listTwo := single.NewLinkedList()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			listTwo.InsertToTail(i)
		}
	}
	listOne.Print()
	listTwo.Print()
	list := SortListMerge(listOne, listTwo)
	list.Print()
}

func TestListDeleteBottomN(t *testing.T) {
	list := single.NewLinkedList()
	for i := 0; i < 1; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	ListDeleteBottomN(list, 1)
	list.Print()

	list = single.NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i)
	}
	list.Print()
	ListDeleteBottomN(list, 3)
	list.Print()
}
