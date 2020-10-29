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
