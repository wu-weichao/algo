package templete

import (
	"demo/list/single"
)

// 求链表的中间结点
// 使用快慢指针，快指针以慢指针2倍的速度遍历，快指针完成遍历时，此时慢指针位于链表的中点
func GetListMiddleNode(list *single.LinkedList) *single.LinkedListNode {
	// 链表为空  不存在哨兵结点或者不存在头结点
	if list.GetHead() == nil || list.GetHead().GetNext() == nil {
		return nil
	}
	// 存在1个或多个结点
	// 使用快慢指针找出中间结点，快指针以慢指针2倍的速度遍历
	slow, fast := list.GetHead(), list.GetHead()
	// 如果快指针遍历完，此时慢指针为中间结点
	for fast != nil && fast.GetNext() != nil {
		slow = slow.GetNext()
		fast = fast.GetNext().GetNext()
	}
	return slow
}

// 单链表反转
// 遍历环，记录当前结点为后继节点使用，调整当前结点的后继节点，指向遍历的前一个结点
func ListReverse(list *single.LinkedList) *single.LinkedList {
	if list.GetHead() == nil || list.GetHead().GetNext() == nil {
		return list
	}
	// 反转 调整结点指向
	// 1->2->3->4->5
	// 1<-2<-3<-4<-5
	// 反转链的头结点
	var pre *single.LinkedListNode = nil
	cur := list.GetHead().GetNext()
	for cur != nil {
		// 记录当前结点的后继结点
		tem := cur.GetNext()
		// 调整当前结点链表方向并记录
		cur.SetNext(pre)
		pre = cur
		// 恢复原链表数据，用于循环
		cur = tem
	}
	// 调整链表哨兵结点的后继结点
	list.GetHead().SetNext(pre)
	return list
}

// 链表中环的检测
// 使用快慢指针，快指针以慢指针2倍的速度遍历
// 如果链表中存在环，那么快慢指针将会重合，由此判断链表存在换
func ListHasCycle(list *single.LinkedList) bool {
	// 空链表
	if list.GetHead() == nil || list.GetHead().GetNext() == nil {
		return false
	}
	slow, fast := list.GetHead(), list.GetHead()
	for fast != nil && fast.GetNext() != nil {
		slow = slow.GetNext()
		fast = fast.GetNext().GetNext()
		// 快慢指针重合，存在环
		if slow == fast {
			return true
		}
	}
	return false
}

// 回文串判断
// 获取链表中间结点，将链表以中间结点拆分为2条链表
// 反转其中1条链表，比较2条链表是否相同
func ListIsPalindrome(list *single.LinkedList) bool {
	// 空链表
	if list.GetHead() == nil || list.GetHead().GetNext() == nil {
		return false
	}
	// 获取中间结点
	midNode := GetListMiddleNode(list)
	// 拆分链表
	fornt := list
	back := single.NewLinkedList()
	cur := list.GetHead().GetNext()
	var backStartNode *single.LinkedListNode
	for cur != nil {
		if cur == midNode {
			// 截断链表
			backStartNode = cur.GetNext()
			back.GetHead().SetNext(backStartNode)
			cur.SetNext(nil)
			break
		}
		cur = cur.GetNext()
	}
	// 只有一个结点
	if back == nil {
		return true
	}
	// 反转链表
	back = ListReverse(back)
	defer func() {
		// 恢复链表
		back = ListReverse(back)
		midNode.SetNext(backStartNode)
	}()
	forntCur := fornt.GetHead().GetNext()
	backCur := back.GetHead().GetNext()
	for forntCur != nil {
		if forntCur.GetValue() != backCur.GetValue() {
			return false
		}
		forntCur = forntCur.GetNext()
		backCur = backCur.GetNext()
	}
	return true
}

// 两个有序的链表合并
func SortListMerge(listOne, listTwo *single.LinkedList) *single.LinkedList {
	if listOne == nil || listOne.GetHead() == nil {
		return listTwo
	}
	if listTwo == nil || listTwo.GetHead() == nil {
		return listOne
	}
	list := single.NewLinkedList()
	cur := list.GetHead()
	curOne := listOne.GetHead().GetNext()
	curTwo := listTwo.GetHead().GetNext()
	// 升序
	for curOne != nil && curTwo != nil {
		// 需要转换为同类型
		if curOne.GetValue().(int) > curTwo.GetValue().(int) {
			cur.SetNext(curTwo)
			curTwo = curTwo.GetNext()
		} else {
			cur.SetNext(curOne)
			curOne = curOne.GetNext()
		}
		cur = cur.GetNext()
	}
	// 剩下为未合并的数据
	if curOne != nil {
		cur.SetNext(curOne)
	} else if curTwo != nil {
		cur.SetNext(curOne)
	}
	return list
}

// 删除链表倒数第 n 个结点
// 设置2个指针，使其中1个指针移动n个结点
// 再同时遍历2个指针，当第1个指针为nil时，则第2个指针为倒数第n的结点
func ListDeleteBottomN(list *single.LinkedList, n int) *single.LinkedList {
	if list == nil || list.GetHead() == nil {
		return list
	}
	// 设置1个指针，移动n个结点
	cur := list.GetHead()
	for i := 0; i < n && cur != nil; i++ {
		cur = cur.GetNext()
	}
	if cur == nil {
		return list
	}
	// 循环获取待删除的结点，从哨兵结点开始遍历，可以得到待删除结点前的一个结点
	tmp := list.GetHead()
	for cur.GetNext() != nil {
		cur = cur.GetNext()
		tmp = tmp.GetNext()
	}
	// 删除结点
	if tmp.GetNext() != nil {
		tmp.SetNext(tmp.GetNext().GetNext())
	} else {
		tmp.SetNext(nil)
	}
	return list
}
