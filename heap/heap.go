package heap

type ArrayHeap struct {
	data []int
	cap  int
	len  int
}

func NewArrayHeap(cap int) *ArrayHeap {
	return &ArrayHeap{
		data: make([]int, cap+1), // 0下标不使用
		cap:  cap,
		len:  0,
	}
}

func (this *ArrayHeap) Insert(v int) {
	if this.len == this.cap {
		return
	}
	this.len++
	this.data[this.len] = v

	// 堆化
	i := this.len
	pre := i / 2
	for pre > 0 && this.data[pre] < v {
		this.data[pre], this.data[i] = this.data[i], this.data[pre]
		i = pre
		pre = i / 2
	}
}

// 移除堆顶元素
func (this *ArrayHeap) RemoveTop() {
	if this.len == 0 {
		return
	}
	// 将堆中最后一个元素移动到对顶，再进行堆化
	this.data[1] = this.data[this.len]
	this.data[this.len] = 0
	this.len--
	this.heapify(1)
}

func (this *ArrayHeap) BuildHeap(data []int) {
	k := 1
	for _, v := range data {
		this.data[k] = v
		k++
	}
	this.len = len(data)

	// 堆化
	// 只需要从 1/2之一处开始，相邻部分会比较
	for i := this.len / 2; i >= 1; i-- {
		this.heapify(i)
	}
}

// 大顶堆
// 堆化 自上而下堆化 可以拆分为相邻小堆移动
func (this *ArrayHeap) heapify(i int) {
	for {
		maxPos := i
		// 左
		if i*2 <= this.len && this.data[i*2] > this.data[maxPos] {
			maxPos = i * 2
		}
		// 右
		if i*2+1 <= this.len && this.data[i*2+1] > this.data[maxPos] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		// 交换
		this.data[i], this.data[maxPos] = this.data[maxPos], this.data[i]
		i = maxPos
	}
}
