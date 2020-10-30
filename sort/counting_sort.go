package sort

// 计数排序
// 主要思想：排序数据范围不大时，按照数据范围切分成一定数量的桶，桶内数据相同，再根据桶的计数数据调整原数组数据
func CountingSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	// 获取数组最大值，以最大值作为桶的数量
	max := 0
	for i := 0; i < n; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	// 桶的大小 需要包含头尾数据，所以要加1
	bucketLen := max + 1
	// 定义桶
	buckets := make([]int, bucketLen)
	// 分桶
	for i := 0; i < n; i++ {
		buckets[data[i]]++
	}

	// 获取计数数据
	// [元素值 => 大于等于元素值的元素个数]
	counting := make([]int, bucketLen)
	count := 0
	for i := 0; i < bucketLen; i++ {
		counting[i] = buckets[i] + count
		count = counting[i]
	}

	// 调整原数组
	// 实现方法
	// 计数数据 counting 中对应的值为大于某元素大小的个数，可以把这个值理解成对应元素最大下标位置
	tem := make([]int, n)

	// 方式1：从前向后循环，非稳定排序，相同元素位置调换
	//for i := 0; i < n; i++ {
	//	counting[data[i]]--
	//	tem[counting[data[i]]] = data[i]
	//}

	// 方式2：从后向前循环，稳定排序
	for i := n - 1; i >= 0; i-- {
		counting[data[i]]--
		tem[counting[data[i]]] = data[i]
	}

	copy(data, tem)
}
