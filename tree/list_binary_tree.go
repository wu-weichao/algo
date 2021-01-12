package tree

import (
	"demo/stack"
	"fmt"
	"strings"
)

type ListBinaryTree struct {
	root *ListBinaryTreeNode
}

type ListBinaryTreeNode struct {
	v     int
	left  *ListBinaryTreeNode
	right *ListBinaryTreeNode
}

func NewListBinaryTreeNode(v int) *ListBinaryTreeNode {
	return &ListBinaryTreeNode{
		v: v,
	}
}

func NewListBinaryTree(v int) *ListBinaryTree {
	return &ListBinaryTree{root: NewListBinaryTreeNode(v)}
}

// 插入 左小右大
func (this *ListBinaryTree) Insert(v int) bool {
	node := this.root
	for nil != node {
		if node.v == v {
			return false
		} else if v < node.v {
			// 插入值小于当前结点
			if nil == node.left {
				node.left = NewListBinaryTreeNode(v)
				return true
			}
			node = node.left
		} else {
			// 插入值大于当前结点
			if nil == node.right {
				node.right = NewListBinaryTreeNode(v)
				return true
			}
			node = node.right
		}
	}

	return false
}

func (this *ListBinaryTree) Find(v int) interface{} {
	node := this.root
	for nil != node {
		if node.v == v {
			return v
		} else if v < node.v {
			// 插入值小于当前结点
			if nil == node.left {
				return nil
			}
			node = node.left
		} else {
			// 插入值大于当前结点
			if nil == node.right {
				return nil
			}
			node = node.right
		}
	}
	return nil
}

func (this *ListBinaryTree) Delete(v int) bool {
	node := this.root
	var pNode *ListBinaryTreeNode = nil
	deleteLeft := false
	for nil != node {
		if node.v == v {
			break
		} else if v < node.v {
			pNode = node
			node = node.left
			deleteLeft = true
		} else {
			pNode = node
			node = node.right
			deleteLeft = false
		}
	}
	// 删除节点不存在
	if nil == node {
		return false
	}
	// 节点存在，删除节点
	// 1.节点不存在子节点，即页子节点，直接删除
	// 2.节点只存在一个子节点，将原父节点指向删除的节点指针调整为删除节点的子节点
	// 3.节点存在两个子节点，将右子树的最小节点（或左子树最大节点）替换被删除节点

	// 删除的是页子节点
	if nil == node.left && nil == node.right {
		if nil == pNode {
			this.root = nil
		} else {
			if deleteLeft {
				pNode.left = nil
			} else {
				pNode.right = nil
			}
		}
	} else if nil != node.right {
		// 删除节点存在右子节点，将右子树最小节点替换待删除节点
		var minPNode *ListBinaryTreeNode = nil
		minNode := node.right // 最小节点
		for nil != minNode {
			if nil == minNode.left {
				break
			}
			minPNode = minNode
			minNode = minNode.left
		}
		// minPNode 为 nil 说明，取到的是右子节点且该节点为页子节点
		if nil == minPNode {
			node.right = nil
		} else {
			minPNode.left = minNode.right // 只能存在右子节点 或 nil
		}
		// minNode 最小节点
		// 替换节点
		minNode.left = node.left
		minNode.right = node.right
		if nil == pNode {
			this.root = minNode
		} else {
			if deleteLeft {
				pNode.left = minNode
			} else {
				pNode.right = minNode
			}
		}
	} else if nil != node.left {
		// 删除节点存在左子节点
		if nil == pNode {
			this.root = node.left
		} else {
			if deleteLeft {
				pNode.left = node.left
			} else {
				pNode.right = node.left
			}
		}
	}

	return true
}

// 前序 当前节点->左->右
func (this ListBinaryTree) PrePrint() {
	// 当前遍历节点
	node := this.root
	// 使用一个栈记录变量的节点
	s := stack.NewArrayStack()
	strFormat := strings.Builder{}
	// 节点不为空 或 栈不为空
	for nil != node || !s.IsEmpty() {
		if nil != node {
			strFormat.WriteString(fmt.Sprintf("[%d]->", node.v))
			// 用栈记录已遍历的左子节点的节点
			s.Push(node)
			// 向左子节点移动
			node = node.left
		} else {
			// 遍历右子节点
			node = s.Pop().(*ListBinaryTreeNode).right
		}
	}
	fmt.Println(strFormat.String())
}

// 中序 左->当前节点->右
func (this ListBinaryTree) InPrint() {
	// 当前遍历节点
	node := this.root
	// 使用一个栈记录变量的节点
	s := stack.NewArrayStack()
	strFormat := strings.Builder{}
	// 节点不为空 或 栈不为空
	for nil != node || !s.IsEmpty() {
		if nil != node {
			// 用栈记录已遍历的左子节点的节点
			s.Push(node)
			// 向左子节点移动
			node = node.left
		} else {
			node = s.Pop().(*ListBinaryTreeNode)
			strFormat.WriteString(fmt.Sprintf("[%d]->", node.v))
			// 遍历右子节点
			node = node.right
		}
	}
	fmt.Println(strFormat.String())
}

// 后序 左->右->当前节点
func (this ListBinaryTree) PostPrint() {
	s1 := stack.NewArrayStack() // 记录待遍历的节点
	s2 := stack.NewArrayStack() // 记录顺序节点

	s1.Push(this.root)
	for !s1.IsEmpty() {
		// 当前遍历节点
		node := s1.Pop().(*ListBinaryTreeNode)
		// 通过栈的先入后出的特性，将后面的数据后入栈
		s2.Push(node)
		// 将左右子节点入栈 s1
		if nil != node.left {
			s1.Push(node.left)
		}
		if nil != node.right {
			s1.Push(node.right)
		}
	}
	strFormat := strings.Builder{}
	for !s2.IsEmpty() {
		node := s2.Pop().(*ListBinaryTreeNode)
		strFormat.WriteString(fmt.Sprintf("[%d]->", node.v))
	}

	fmt.Println(strFormat.String())
}
