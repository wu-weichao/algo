package sort

// 插入排序
// 主要思想：将数据分为2个区间（有序区间和无序区间），将无序区间的元素插入到有序区间，并保证有序区间的有序
func InsertionSort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}
	// 需要循环的次数，即需要操作元素的个数，默认第一个元素是有续的，所有i从下表为1的元素开始
	for i := 1; i < n; i++ {
		// 需要比较元素的个数，即有序区间内的个数
		for j := i; j > 0; j-- {
			// 比较，调整元素位置
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			} else {
				// 优化 由于有序区间已经是有序了，如果出现不符合比较条件的情况，说明插入的元素已插入到正确的位置中
				break
			}
		}
	}
	return data
}
