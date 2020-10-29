package sort

// 选择排序
// 主要思想：将数据分为2个区间（有序区间和无序区间），按照排序规则从无序区间获取最大或最小的元素插入到有序区间的末尾
func SelectionSort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}
	// 需要循环的次数，即需要操作元素的个数
	for i := 0; i < n; i++ {
		minIndex := i // 最小的元素的下标
		// 需要比较操作的次数，即无序区间获取指定规则元素的次数
		for j := i + 1; j < n; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		// 交换元素
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}
