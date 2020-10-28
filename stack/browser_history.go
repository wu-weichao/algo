package stack

import (
	"fmt"
)

type BrowserHistory struct {
	curPage interface{}
	back    *ArrayStack // 后退栈
	goahead *ArrayStack // 前进栈
}

func NewBrowserHistory() *BrowserHistory {
	return &BrowserHistory{nil, NewArrayStack(), NewArrayStack()}
}

// 浏览器跳转
func (this *BrowserHistory) To(v interface{}) {
	if this.curPage != nil {
		// 后退栈记录当前页
		this.back.Push(this.curPage)
	}
	//	前进栈 清空
	this.goahead.Flush()
	// 修改当前页
	this.curPage = v

}

func (this *BrowserHistory) Goahead() {
	if this.goahead.IsEmpty() {
		return
	}
	page := this.goahead.Pop()
	// 后退栈 记录当前页
	this.back.Push(this.curPage)
	// 调整当前页面
	this.curPage = page
}

func (this *BrowserHistory) Back() {
	if this.back.IsEmpty() {
		return
	}
	page := this.back.Pop()
	// 前进栈 记录当前页
	this.goahead.Push(this.curPage)
	// 调整当前页面
	this.curPage = page
}

func (this *BrowserHistory) Print() {
	fmt.Println("current page is:", this.curPage)
	fmt.Println("back list:")
	this.back.Print()
	fmt.Println("goahead list:")
	this.goahead.Print()
	fmt.Println()
}
