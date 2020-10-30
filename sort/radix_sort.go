package sort

import (
	"math"
	"strconv"
)

// 基数排序
// 主要思想：将数组中的元素从后向前按位比较大小，指定位比较过后调整依次数据排序，直到最高位结束
func RadixSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	// 获取数组中最大的数
	max := 0
	for i := 0; i < n; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	count := len(strconv.Itoa(max)) // 需要循环的次数

	// 比较的元素为整型，分成10个桶
	buckets := make([][]int, 10)
	radixIndex := 1 // 比较位
	for radixIndex <= count {
		// 分桶
		for i := 0; i < n; i++ {
			// 获取指定位的值
			bucketIndex := data[i] / int(math.Pow10(radixIndex-1)) % 10
			buckets[bucketIndex] = append(buckets[bucketIndex], data[i])
		}
		// 合并数据
		start := 0
		for j := 0; j < 10; j++ {
			curLen := len(buckets[j])
			if curLen > 0 {
				copy(data[start:start+curLen], buckets[j])
				buckets[j] = buckets[j][:0] // 清空切片中的元素
				start += curLen
			}
		}
		// 比较位
		radixIndex++
	}
}
