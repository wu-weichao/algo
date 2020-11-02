package binary

// 查找数据集合中第一个等于指定值的元素
func BinarySearchEqFirst(data []int, v int) int {
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
			if mid == 0 || data[mid-1] != v {
				return mid
			}
			// 说明左边更小的值
			high = mid - 1
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

// 查找数据集合中最后一个等于指定值的元素
func BinarySearchEqLast(data []int, v int) int {
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
			if (mid == n-1) || data[mid+1] != v {
				return mid
			}
			// 说明左边更小的值
			low = mid + 1
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

// 查找数据集合中第一个大于等于指定值的元素
func BinarySearchEgtFirst(data []int, v int) int {
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
		if data[mid] >= v {
			if mid == 0 || data[mid-1] < v {
				return mid
			}
			// 查找左边
			high = mid - 1
		} else if data[mid] < v {
			// 查找右边
			low = mid + 1
		}
	}
	return -1
}

// 查找数据集合中最后一个小于等于指定值的元素
func BinarySearchEltLast(data []int, v int) int {
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
		if data[mid] > v {
			// 查找左边
			high = mid - 1
		} else if data[mid] <= v {
			if (mid == n-1) || data[mid+1] > v {
				return mid
			}
			// 查找右边
			low = mid + 1
		}
	}
	return -1
}
