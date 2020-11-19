package hashmap

import (
	"fmt"
	"strconv"
	"strings"
)

// 散列表-开放寻址法
// 当出现哈希冲突时，依次往后查找，插入第一个空位中
type AddressHashMap struct {
	data       []*AddressHashMapNode
	capacity   int
	length     int
	loadFactor float64
}

type AddressHashMapNode struct {
	k      int
	v      interface{}
	delete bool
}

func NewAddressHashMapNode(k int, v interface{}) *AddressHashMapNode {
	return &AddressHashMapNode{k, v, false}
}

func NewAddressHashMap(capacity int) *AddressHashMap {
	return &AddressHashMap{
		data:       make([]*AddressHashMapNode, capacity),
		capacity:   capacity,
		length:     0,
		loadFactor: 0.7, // 装载因子
	}
}

// 不存在
func (this *AddressHashMap) Put(key int, v interface{}) bool {
	if nil == v {
		return false
	}
	// 判断key是否存在
	existsNode := this.Get(key)
	if nil != existsNode {
		existsNode.v = v
		return true
	}
	// 获取对应的hash值
	hash := this.hash(key)
	// 存在hash冲突
	if this.data[hash] != nil && !this.data[hash].delete {
		// 查找可以插入的位置
		i := hash + 1
		for i != hash {
			if i >= this.capacity {
				i = 0
			}
			// 位置为空 或者 被删除的节点
			if this.data[i] == nil || this.data[i].delete {
				hash = i
				break
			}
			i++
		}
	}
	// 插入
	this.data[hash] = NewAddressHashMapNode(key, v)
	this.length++
	// 扩容
	this.grow()

	return true
}

func (this *AddressHashMap) Get(key int) *AddressHashMapNode {
	if this.length == 0 {
		return nil
	}
	hash := this.hash(key)
	// 不存在
	if this.data[hash] == nil {
		return nil
	}
	// 存在且值等于被查询值
	if this.data[hash].k == key && !this.data[hash].delete {
		return this.data[hash]
	}
	// 存在，但是值不能与被查询值，下后查找
	i := hash + 1
	for i != hash {
		if i >= this.capacity {
			i = 0
		}
		// 存在nil的结点说明不存在值
		if this.data[i] == nil {
			return nil
		}
		if this.data[i].k == key && !this.data[i].delete {
			return this.data[i]
		}
		i++
	}

	return nil
}

func (this *AddressHashMap) Delete(key int) bool {
	if this.length == 0 {
		return false
	}
	node := this.Get(key)
	if nil == node {
		return false
	}
	node.delete = true
	this.length--
	return true
}

func (this *AddressHashMap) Print() {
	strFormat := strings.Builder{}
	strFormat.WriteString(fmt.Sprintf("------ length: %d, capacity: %d ------\r\n", this.length, this.capacity))
	count := 0
	for i := 0; i < this.capacity; i++ {
		if this.data[i] != nil && !this.data[i].delete {
			strFormat.WriteString(fmt.Sprintf("[%d]%v", this.data[i].k, this.data[i].v))
			count++
			if count == this.length {
				break
			}
			strFormat.WriteString("->")
		}
	}
	fmt.Println(strFormat.String())
}

// hash值
func (this *AddressHashMap) hash(key int) int {
	return (key ^ (key >> 32)) & (this.capacity - 1)
}

// 扩容
func (this *AddressHashMap) grow() {
	loadFactor, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(this.length)/float64(this.capacity)), 64)
	if loadFactor < this.loadFactor {
		return
	}

	// 按照2倍扩容
	oldCapacity := this.capacity
	this.capacity = this.capacity << 1
	newData := make([]*AddressHashMapNode, this.capacity)
	// 移动数据
	for _, node := range this.data {
		if node != nil && !node.delete {
			hash := this.hash(node.k)
			//fmt.Println(node.v.(int))
			if newData[hash] != nil {
				// 冲突 向下查找
				i := hash + 1
				for i != hash {
					if i >= oldCapacity {
						i = 0
					}
					// 位置为空 或者 被删除的节点
					if newData[i] == nil {
						hash = i
						break
					}
					i++
				}
			}
			newData[hash] = node
		}
	}
	// 修改
	this.data = newData
}
