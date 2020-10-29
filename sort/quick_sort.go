package sort

// 快速排序
// 主要思想：取数据区间内某个元素作为分区点(下标)，将数组切分为小于、大于分区点的两块数据，递归执行此逻辑
func QuickSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}
	quickSort(data, 0, n-1)
}

func quickSort(data []int, start, end int) {
	// 终止条件
	if start >= end {
		return
	}

	// 获取下标并调整数组元素位置
	pivot := data[end] // 分区点的值
	// 原地排序 将小于分区值的元素交换至左右
	// 左边区间下一个索引位置
	leftIndex := start
	for i := start; i < end; i++ {
		if data[i] < pivot {
			// 将当前元素交换至左区间，并增加索引值
			data[leftIndex], data[i] = data[i], data[leftIndex]
			leftIndex++
		}
	}
	// 将分区值置于索引位置
	data[leftIndex], data[end] = data[end], data[leftIndex]

	quickSort(data, start, leftIndex-1)
	quickSort(data, leftIndex+1, end)
}
