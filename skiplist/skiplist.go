package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// 跳表
type SkipList struct {
	head   *SkipListNode // 头结点，哨兵结点
	level  int           // 跳表层级
	length int           // 跳表长度
}

// 最高层
const MAX_LEVEl = 16

// 跳表结点
type SkipListNode struct {
	v        interface{}     // 结点的值
	score    int             // 结点的排序值
	level    int             // 结点所在的层
	forwards []*SkipListNode // 结点每层前进的指针，即后继结点，只会包含下层的结点
}

// 新建跳表结点
func NewSkipListNode(v interface{}, score, level int) *SkipListNode {
	return &SkipListNode{
		v:        v,
		score:    score,
		level:    level,
		forwards: make([]*SkipListNode, level, level),
	}
}

// 新建跳表
func NewSkipList() *SkipList {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 跳表指针
	return &SkipList{
		head:   NewSkipListNode(nil, math.MinInt32, MAX_LEVEl),
		level:  1,
		length: 0,
	}
}

// 插入结点
// 每一层都是一个有序链表
// 如果元素出现在某一层，则所以比该层小的层都包含该元素
func (this *SkipList) Insert(v interface{}, score int) *SkipListNode {
	if nil == v {
		return nil
	}
	// 查找是否存在相同元素，获取待修改的元素
	cur := this.head
	// update 用于记录插入后需要调整的结点，这些结点为插入结点在各层中的 前驱结点
	update := [MAX_LEVEl]*SkipListNode{}
	for i := MAX_LEVEl - 1; i >= 0; i-- {
		// 当前层存在索引
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				// 存在相同结点，以值为判断
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				// 找到第一个比要插入结点大的结点
				update[i] = cur // 记录当前结点
				break
			}
			// 后继指针，遍历下一个结点的
			cur = cur.forwards[i]
		}
		// 当前层不存在或者遍历到最后未遍历到符合要求的结点
		if nil == cur.forwards[i] {
			update[i] = cur // 记录结点
		}
	}

	// 通过随机函数维护跳表的平衡性，通过随机数向 0-随机数层插入新的结点
	level := rand.Intn(MAX_LEVEl-1) + 1 // 随机层数

	newNode := NewSkipListNode(v, score, level)
	for i := 0; i < level; i++ {
		// 在待修改的列表中，插入结点
		next := update[i].forwards[i]   // 原后继结点
		update[i].forwards[i] = newNode // 新结点
		newNode.forwards[i] = next      // 将原后继结点调整为新结点的后继结点
	}
	// 更新跳表层级，如果随机出来的层级比当前层级高，则更新
	if level > this.level {
		this.level = level
	}
	// 更新跳表长度
	this.length++

	return newNode
}

// 查找结点
func (this *SkipList) Find(v interface{}, score int) *SkipListNode {
	if this.length == 0 || v == nil {
		return nil
	}
	// 查找逻辑
	// 从高层遍历至底层，在当前层出现第一个大于排序值的结点时，则从上一个结点向下一层继续查找，知道找到指定结点或者所有层级遍历完
	// 示例：
	// i层 	 X    X    X               X    X
	// 	     ..........|.............(X>?)
	// i-1层 ..........X   X   X   X   X
	//       ..........|.........(X>?)
	// i-2层 ..................X......
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				// 找到结点
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				// 遍历的结点比要找到的结点大，即当前层遍历的终点
				// 跳出当前循环，遍历下一层链表
				break
			}
			// 小于当前结点，将指针调整至后继指针
			cur = cur.forwards[i]
		}
	}

	return nil
}

func (this *SkipList) Delete(v interface{}, score int) bool {
	if nil == v || this.level == 0 {
		return false
	}
	// 查找要被删除的结点信息
	cur := this.head
	update := [MAX_LEVEl]*SkipListNode{}
	flag := false
	for i := this.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				// 找到待删除结点的前驱结点
				flag = true
				update[i] = cur
				break
			} else if cur.forwards[i].score > score {
				// 找到第一个比要插入结点大的结点
				break
			}
			// 小于当前结点，将指针调整至后继指针
			cur = cur.forwards[i]
		}
	}
	if !flag {
		return false
	}

	// 更新前驱节点信息
	for i := 0; i <= this.level; i++ {
		if update[i] != nil {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	this.length--

	return true
}

func (this *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", this.level, this.length)
}

func (this *SkipList) Print() {
	for i := this.level - 1; i >= 0; i-- {
		cur := this.head.forwards[i]
		fmt.Println("--------", i, "-------------")
		strFormat := strings.Builder{}
		for cur != nil {
			strFormat.WriteString(fmt.Sprintf("%s", cur.v))
			cur = cur.forwards[i]
			if cur != nil {
				strFormat.WriteString("->")
			}
		}
		fmt.Println(strFormat.String())
	}
}
