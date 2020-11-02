package binary

// 二分查找
// 主要思想：对一个有序的数据集合，每次通过跟区间的中间元素比较，将待查找的区间缩小为之前的一半，直到找到要查找的元素或者区间被缩小为0
func BinarySearch(data []int, v int) int {
	// 数组大小
	n := len(data)
	if n < 1 {
		return -1
	}

	// 中间元素下标
	low := 0
	high := n - 1
	for low <= high {
		// 高位减低位防止相加的溢出
		// 区中，位运算更快
		mid := low + ((high - low) >> 1)
		// 比较中间元素
		if data[mid] == v {
			return mid
		} else if data[mid] > v {
			// 查找左边
			high = mid - 1
		} else if data[mid] < v {
			// 查找右边
			low = mid + 1
		}
	}
	return -1
}
