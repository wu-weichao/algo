package stack

import (
	"fmt"
	"strings"
)

type ArrayStack struct {
	data []interface{}
	top  int // 栈顶指针
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(v interface{}) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top++
	}
	// 栈顶指针大于切片长度-1，说明栈顶元素对应的切片下标的元素还未被设置过值
	// 直接设置对应下标元素会报错 “index out of range [x] with length x”
	// 栈顶指针大于切片长度-1，使用append函数
	if this.top > len(this.data)-1 {
		// 使用append函数追加元素，并对数组扩容
		this.data = append(this.data, v)
	} else {
		// 空间足够
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.top < 0 {
		return nil
	}
	v := this.data[this.top]
	this.top--
	return v
}

// 返回栈顶元素，但不删除
func (this *ArrayStack) Top() interface{} {
	if this.top < 0 {
		return nil
	}
	return this.data[this.top]
}

func (this *ArrayStack) Flush() {
	// 将栈顶指针重置
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.top < 0 {
		fmt.Println("stack is empty")
		return
	}
	formatStr := strings.Builder{}
	for i := this.top; i >= 0; i-- {
		formatStr.WriteString(fmt.Sprintf("%v\r\n", this.data[i]))
	}
	fmt.Println(formatStr.String())
}
