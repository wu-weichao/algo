package sort

// 归并排序
// 主要思想：将数组对半拆分成最小单元后，再重新排序并合并，递归执行此逻辑
func MergeSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}
	mergeSort(data, 0, n-1)
}

func mergeSort(data []int, start, end int) {
	if start >= end {
		return
	}
	// 获取中间元素的下标位置
	midIndex := (start + end) / 2
	// 拆分数组
	mergeSort(data, start, midIndex)
	mergeSort(data, midIndex+1, end)

	// 合并数组
	leftIndex := start
	rightIndex := midIndex + 1
	tem := make([]int, end-start+1)
	temIndex := 0
	// 比较左右两个区间内数组元素大小
	for ; leftIndex <= midIndex && rightIndex <= end; temIndex++ {
		if data[leftIndex] <= data[rightIndex] {
			tem[temIndex] = data[leftIndex]
			leftIndex++
		} else {
			tem[temIndex] = data[rightIndex]
			rightIndex++
		}
	}
	// 剩余数据处理
	for ; leftIndex <= midIndex; leftIndex++ {
		tem[temIndex] = data[leftIndex]
		temIndex++
	}
	for ; rightIndex <= end; rightIndex++ {
		tem[temIndex] = data[rightIndex]
		temIndex++
	}
	// 切片操作是左闭右开[),end需要+1才能包含end下标的元素的操作
	// 将排好序的数组覆盖到原数组上
	copy(data[start:end+1], tem)
}
