package sort

// 桶排序
// 主要思想：将数据按照一定的规则/范围切分置不同的桶内，再对桶内数据排序，最后依次合并所有的桶内数据
func BucketSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	// 切分数据 按照一定的规则，根据实际数据决定
	// 这里示例，简单的按照数组长度确定分桶数量
	max := 0
	for i := 0; i < n; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	// 定义桶数据
	buckets := make([][]int, n)
	// 分桶 这里按照和最大值比较分桶：0 ~ n-1
	for i := 0; i < n; i++ {
		bucketIndex := data[i] * (n - 1) / max
		buckets[bucketIndex] = append(buckets[bucketIndex], data[i])
	}

	// 对桶内的数据进行排序
	for i := 0; i < n; i++ {
		if len(buckets[i]) > 0 {
			QuickSort(buckets[i])
		}
	}

	// 合并数据
	start := 0
	for i := 0; i < n; i++ {
		curLen := len(buckets[i])
		if curLen > 0 {
			copy(data[start:start+curLen], buckets[i])
			start += curLen
		}
	}
}
