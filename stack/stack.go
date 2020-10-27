package stack

// 栈 接口
type Stack interface {
	Push(v interface{}) // 入栈
	Pop() interface{}   // 出栈
	IsEmpty() bool      // 判断栈是否为空
	Top() interface{}   // 获取栈顶元素
	Flush()             // 清空栈
}
