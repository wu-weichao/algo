package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{5, 4, 9, 10, 6, 1, 7, 2, 8, 3}
	fmt.Printf("old array: %v \r\n", data)
	data = BubbleSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestInsertionSort(t *testing.T) {
	data := []int{5, 4, 9, 10, 6, 1, 7, 2, 8, 3}
	fmt.Printf("old array: %v \r\n", data)
	data = InsertionSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestSelectionSort(t *testing.T) {
	data := []int{5, 4, 9, 10, 6, 1, 7, 2, 8, 3}
	fmt.Printf("old array: %v \r\n", data)
	data = SelectionSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestMergeSort(t *testing.T) {
	data := []int{5, 4, 9, 10, 6, 1, 7, 2, 8, 3}
	fmt.Printf("old array: %v \r\n", data)
	MergeSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestQuickSort(t *testing.T) {
	data := []int{5, 4, 9, 10, 6, 1, 7, 2, 8, 3}
	fmt.Printf("old array: %v \r\n", data)
	QuickSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestBucketSort(t *testing.T) {
	data := []int{2, 5, 9, 10, 9, 3, 4, 2, 5, 3}
	fmt.Printf("old array: %v \r\n", data)
	BucketSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestCountingSort(t *testing.T) {
	data := []int{2, 5, 9, 10, 9, 3, 4, 2, 5, 3}
	fmt.Printf("old array: %v \r\n", data)
	CountingSort(data)
	fmt.Printf("new array: %v \r\n", data)
}

func TestRadixSort(t *testing.T) {
	data := []int{87, 98, 156, 153, 57, 58, 93, 18, 5, 23}
	fmt.Printf("old array: %v \r\n", data)
	RadixSort(data)
	fmt.Printf("new array: %v \r\n", data)
}
