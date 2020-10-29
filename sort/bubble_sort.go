package sort

// 冒泡排序
// 主要思想：根据某种规则（升序或者降序）依次比较相邻两位元素的大小并调整位置
func BubbleSort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}
	var flag bool // 记录一次排序的过程是否发生交互，如果没有发生，说明数据已经有顺
	// 需要循环的次数，即需要操作元素的次数
	for i := 0; i < n-1; i++ {
		flag = true
		// 单次循环，需要比较元素的次数 从第一位开始
		for j := 0; j < n-i-1; j++ {
			// 升序，交换元素位置
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return data
}
