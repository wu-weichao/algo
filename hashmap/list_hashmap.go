package hashmap

import (
	"fmt"
	"strings"
)

// 散列表-链表法
type ListHashMap struct {
	data       []*ListHashMapNode
	capacity   int     // 表容量
	length     int     // 表元素个数
	loadFactor float64 // 负载因子
	bucketCnt  int     // 最大桶深 桶内元素大于这个值时，扩容散列表
}

// 散列表结点
type ListHashMapNode struct {
	k     int
	v     interface{}
	hnext *ListHashMapNode // 散列表后继指针
}

func NewListHashMapNode(k int, v interface{}) *ListHashMapNode {
	return &ListHashMapNode{k, v, nil}
}

func NewListHashMap(capacity int) *ListHashMap {
	return &ListHashMap{
		data:       make([]*ListHashMapNode, capacity),
		capacity:   capacity,
		length:     0,
		loadFactor: 0.7,
		bucketCnt:  8,
	}
}

func (this *ListHashMap) Put(key int, v interface{}) bool {
	if nil == v {
		return false
	}
	// 获取对应的hash值
	hash := this.hash(key)
	// 插入操作
	// 追加到散列表桶中的尾部和链表的尾部
	node := this.data[hash]
	if node == nil {
		this.data[hash] = NewListHashMapNode(key, v)
		this.length++
	} else {
		// 获取当前链表 最后一个元素
		for node != nil {
			// 重复插入
			if node.v == v {
				return true
			}
			if node.hnext != nil {
				node = node.hnext
			} else {
				node.hnext = NewListHashMapNode(key, v)
				this.length++
				break
			}
		}
	}
	// 扩容待实现...

	return true
}

func (this *ListHashMap) Get(key int) interface{} {
	// 获取对应的hash值
	hash := this.hash(key)

	node := this.data[hash]
	if node == nil {
		return nil
	}
	for node != nil {
		// 重复插入
		if node.k == key {
			return node.v
		}
		node = node.hnext
	}
	return nil
}

func (this *ListHashMap) Delete(key int) bool {
	// 获取对应的hash值
	hash := this.hash(key)

	node := this.data[hash]
	if node == nil {
		return false
	}
	// 数据是头结点
	if node.k == key {
		this.data[hash] = node.hnext
		this.length--
		return true
	}
	// 数据在链表内
	if node.hnext != nil {
		pre := node
		cur := node.hnext
		for cur != nil {
			if cur.k == key {
				pre.hnext = cur.hnext
				this.length--
				return true
			}
			cur = cur.hnext
		}
		return false
	}

	return false
}

func (this *ListHashMap) Print() {
	strFormat := strings.Builder{}
	strFormat.WriteString(fmt.Sprintf("------ length: %d, capacity: %d ------\r\n", this.length, this.capacity))
	count := 0
	for i := 0; i < this.capacity; i++ {
		strFormat.WriteString(fmt.Sprintf("hash[%d]:", i))
		node := this.data[i]
		if node != nil {
			for node != nil {
				strFormat.WriteString(fmt.Sprintf("[%d]%v", node.k, node.v))
				count++
				node = node.hnext
				if node != nil {
					strFormat.WriteString("->")
				}
			}
			if count == this.length {
				break
			}
		}
		strFormat.WriteString("\n")
	}
	fmt.Println(strFormat.String())
}

// hash值
func (this *ListHashMap) hash(key int) int {
	return (key ^ (key >> 32)) & (this.capacity - 1)
}
