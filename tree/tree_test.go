package tree

import "testing"

func TestListBinaryTree(t *testing.T) {
	tree := NewListBinaryTree(50)
	tree.Insert(30)
	tree.Insert(60)
	tree.Insert(19)
	tree.Insert(35)
	tree.Insert(56)
	tree.Insert(70)
	tree.Insert(31)
	tree.Insert(45)
	tree.Insert(65)
	tree.Insert(75)
	tree.Insert(47)
	tree.Insert(68)
	tree.Insert(73)
	// 遍历
	tree.PrePrint()
	//tree.InPrint()
	//tree.PostPrint()
	// 查找
	//t.Log(tree.Find(25))
	//t.Log(tree.Find(28))
	// 删除
	//tree.Delete(19)
	//tree.Delete(47)
	//tree.Delete(65)
	//tree.Delete(75)
	//tree.Delete(60)
	tree.Delete(30)
	tree.PrePrint()
}
